package main

import (
	"context"
	"encoding/json"
	"go-cqrs-api/api"
	"go-cqrs-api/config"
	sqsconsumer "go-cqrs-api/consumer"
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"
	"go-cqrs-api/logger"
	"go-cqrs-api/queries"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func main() {
	config.Load()
	logger.Init()
	infrastructure.InitDB()
	infrastructure.InitSQS()

	handler := func(msg types.Message) error {
		var event domain.DeviceEvent
		jsonBody := *msg.Body
		err := json.Unmarshal([]byte(jsonBody), &event)
		if err != nil {
			log.Println("Failed to unmarshal message body:", err)
			return err
		}
		queries.ProcessDeviceEvent(event)

		return nil
	}

	consumer, err := sqsconsumer.NewConsumerFromClient(infrastructure.SqsClient, infrastructure.QueueURL, handler)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	go consumer.StartPolling(context.Background())

	logger.Log.Info("App started")

	r := api.InitRouter()
	http.ListenAndServe(":8080", r)
}
