package events

import (
	"context"
	"encoding/json"
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func PublishDeviceDataEvent(data domain.DeviceEvent) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = infrastructure.SqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    aws.String(infrastructure.QueueURL),
		MessageBody: aws.String(string(body)),
	})
	return err
}
