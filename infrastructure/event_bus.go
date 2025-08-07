package infrastructure

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var SqsClient *sqs.Client
var QueueURL string

const queueName = "devices-data-events"

func InitSQS() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Failed to load AWS config:", err)
	}

	SqsClient = sqs.NewFromConfig(cfg)

	// Get or create the queue
	out, err := SqsClient.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		log.Fatal("Failed to get SQS queue URL:", err)
	}
	QueueURL = *out.QueueUrl
}
