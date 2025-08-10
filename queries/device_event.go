package queries

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-cqrs-api/domain"
	"go-cqrs-api/infrastructure"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3client "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func ProcessDeviceEvent(event domain.DeviceEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	key := fmt.Sprintf("events/%s/%d.json", event.DeviceID, time.Now().UnixNano())
	s3 := infrastructure.GetS3Client()
	bucket := infrastructure.GetS3Bucket()
	// PutObject with context
	_, err = s3.PutObject(context.TODO(), &s3client.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		ACL:    types.ObjectCannedACLPrivate,
	})
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	fmt.Println("Uploaded event to S3:", key)
	return nil
}
