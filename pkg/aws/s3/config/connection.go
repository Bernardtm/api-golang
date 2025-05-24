package config

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ConnectionS3(s3Region string, s3AccessKeyID string, s3SecretAccessKey string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s3Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s3AccessKeyID,
			s3SecretAccessKey,
			"", // Optional, leave as "" for long-term credentials
		)))

	if err != nil {
		return nil, errors.New("error loading AWS config")
	}

	return s3.NewFromConfig(cfg), nil
}
