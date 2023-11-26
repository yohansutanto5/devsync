package main

import (
	"app/cmd/config"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func startConsumerListener(config config.Configuration) {
	// Configure Kafka consumer
	consumerClient, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", config.Kafka.Hostname, config.Kafka.Port),
		"sasl.mechanisms":   config.Kafka.AuthMethod,
		"security.protocol": "sasl_ssl",
		"sasl.username":     config.Kafka.User,
		"sasl.password":     config.Kafka.Password,
		"group.id":          config.Kafka.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("Failed to Initiate Kafka Consumer Connection")
	}

	// Subscribe to the Kafka topic
	err = consumerClient.SubscribeTopics([]string{config.Kafka.ApprovalTopic}, nil)
	if err != nil {
		log.Fatal("Error subscribing to topic: ", err)
	}

	// Set up context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Wait group to wait for the consumer and other goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Start a goroutine to handle messages
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return // Exit goroutine on context cancellation
			default:
				ev := consumerClient.Poll(10) // Poll for Kafka events

				switch e := ev.(type) {
				case *kafka.Message:
					fmt.Printf("Received message: %s\n", e.Value)
					// Add your message handling logic here
				case kafka.Error:
					fmt.Printf("Error: %v\n", e)
				}
			}
		}
	}()

	// Handle termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to listen for termination signals
	go func() {
		defer cancel() // Cancel the context on termination signals
		<-signalChan   // Wait for termination signals
	}()

	// Wait for the consumer and other goroutines to finish
	wg.Wait()

	// Close the Kafka consumer
	if err := consumerClient.Close(); err != nil {
		log.Println("Error closing Kafka consumer: ", err)
	}
}
