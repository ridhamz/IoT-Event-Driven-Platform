package infrastructure

import (
	"context"
	"fmt"
	appConfigPackage "go-cqrs-api/config"

	"github.com/aws/aws-sdk-go-v2/config"
	s3client "github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3     *s3client.Client
	bucket string
)

func InitS3Client() {
	bucket = appConfigPackage.AppConfig.S3_BUCKET
	if bucket == "" {
		panic("S3_BUCKET env var is required")
	}

	// Load the default AWS config with region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(appConfigPackage.AppConfig.S3_BUCKET_REGION))
	if err != nil {
		panic(fmt.Sprintf("unable to load AWS SDK config, %v", err))
	}

	// Create S3 client
	s3 = s3client.NewFromConfig(cfg)
}

func GetS3Client() *s3client.Client {
	if s3 == nil {
		InitS3Client()
	}
	return s3
}

func GetS3Bucket() string {
	if bucket == "" {
		panic("S3_BUCKET is not set")
	}
	return bucket
}
