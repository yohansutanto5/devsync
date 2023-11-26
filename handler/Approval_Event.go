package handler

import (
	"app/model"
	"app/pkg/log"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ApprovalEventAction(message *kafka.Message) {
	// Perform processing based on the content of the message
	var inputMessage model.ApprovalTopicMessage

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(message.Value, &inputMessage)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	log.Warning(0, "print incoming message", inputMessage)
}
