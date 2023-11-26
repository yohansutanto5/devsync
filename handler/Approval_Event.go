package handler

import (
	"app/constanta"
	"app/model"
	"app/pkg/log"
	"app/service"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ApprovalEventAction(message *kafka.Message, ReleaseOPS service.ReleaseOPSService) {
	// Perform processing based on the content of the message
	var inputMessage model.ApprovalTopicMessage

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(message.Value, &inputMessage)
	if err != nil {
		log.Error(0, "Error decoding JSON:", err.Error(), nil)
		return
	}

	if inputMessage.Service == "releaseops" {
		errSVC := ReleaseOPS.WorkflowSignal(inputMessage.ID, inputMessage.Action)
		if err != nil {
			log.Error(inputMessage.ID, "Failed Change Ticket Status", constanta.CodeErrorService, errSVC.Error())
		}
	} else if inputMessage.Service == "application"{
		// invalid service
	}

}
