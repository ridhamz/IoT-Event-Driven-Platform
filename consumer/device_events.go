package sqsconsumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type HandlerFunc func(message types.Message) error

type Consumer struct {
	client   *sqs.Client
	queueURL string
	handler  HandlerFunc
}

// NewConsumerFromClient creates a consumer from an existing client and queue URL
func NewConsumerFromClient(client *sqs.Client, queueURL string, handler HandlerFunc) (*Consumer, error) {
	if client == nil {
		return nil, fmt.Errorf("SQS client cannot be nil")
	}
	if queueURL == "" {
		return nil, fmt.Errorf("Queue URL cannot be empty")
	}

	return &Consumer{
		client:   client,
		queueURL: queueURL,
		handler:  handler,
	}, nil
}

func (c *Consumer) StartPolling(ctx context.Context) {
	for {
		// Check for context cancellation for graceful shutdown
		select {
		case <-ctx.Done():
			log.Println("SQS consumer stopped")
			return
		default:
		}

		err := c.pollAndHandleMessages(ctx)
		if err != nil {
			log.Printf("Error polling messages: %v", err)
			time.Sleep(5 * time.Second) // wait before retrying on error
		}
	}
}

func (c *Consumer) pollAndHandleMessages(ctx context.Context) error {
	resp, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(c.queueURL),
		MaxNumberOfMessages: 5,  // max 10 allowed
		WaitTimeSeconds:     20, // long polling
		VisibilityTimeout:   30, // hide message during processing
	})
	if err != nil {
		return err
	}

	if len(resp.Messages) == 0 {
		// no messages this cycle, just return
		return nil
	}

	for _, msg := range resp.Messages {
		log.Printf("Received message: %s", aws.ToString(msg.Body))

		err := c.handler(msg)
		if err != nil {
			log.Printf("Handler error: %v", err)
			// don't delete message on error, it will reappear later
			continue
		}

		_, err = c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(c.queueURL),
			ReceiptHandle: msg.ReceiptHandle,
		})
		if err != nil {
			log.Printf("Failed to delete message: %v", err)
		}
	}

	return nil
}
