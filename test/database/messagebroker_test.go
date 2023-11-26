package dbtest

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProduceMessage(t *testing.T) {
	message := fmt.Sprintf("ApprovalGranted:%s", "1")
	approvalTopic := configs.Kafka.ApprovalTopic
	// Produce message to Kafka
	err := ds.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &approvalTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	assert.ErrorIs(t, nil, err)
}

func TestConsumerMessage(t *testing.T) {
	// Subscribe to the approval events topic
	ds.Consumer.SubscribeTopics([]string{configs.Kafka.ApprovalTopic}, nil)

	// Handle termination signals
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Main loop for consuming events
	for {
		select {
		case <-sigchan:
			fmt.Println("Received termination signal. Exiting...")
			return
		default:
			ev := ds.Consumer.Poll(1)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Received message: %s\n", e.Value)
				// Parse the message and perform the corresponding action (e.g., trigger the task)
			case kafka.Error:
				fmt.Printf("Error: %v\n", e)
			}
		}
	}
}
