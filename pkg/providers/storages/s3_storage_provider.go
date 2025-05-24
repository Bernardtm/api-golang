package storages

import (
	"bernardtm/backend/configs"
	"bernardtm/backend/pkg/aws/s3/config"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type UploadFileOutput struct {
	Name   string `json:"file_name"`
	Link   string `json:"file_link"`
	Folder string `json:"file_folder"`
	Type   string `json:"file_type"`
}

type s3StorageProvider struct {
	S3_BUCKET_NAME       string
	S3_REGION            string
	S3_ACCESS_KEY_ID     string
	S3_SECRET_ACCESS_KEY string
}

func NewS3StorageProvider(config *configs.AppConfig) *s3StorageProvider {
	return &s3StorageProvider{
		S3_BUCKET_NAME:       config.S3_BUCKET_NAME,
		S3_REGION:            config.S3_REGION,
		S3_ACCESS_KEY_ID:     config.S3_ACCESS_KEY_ID,
		S3_SECRET_ACCESS_KEY: config.S3_SECRET_ACCESS_KEY,
	}
}

func (p *s3StorageProvider) Upload(dto UploadDto) (interface{}, error) {
	s3Client, err := config.ConnectionS3(p.S3_REGION, p.S3_ACCESS_KEY_ID, p.S3_SECRET_ACCESS_KEY)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client: %v", err)
	}

	fileType := strings.ToLower(filepath.Ext(dto.OriginalFileName))

	if fileType == "" {
		fileType = "others"
	}

	fileType = fileType[1:]

	folder := fileType
	// Get the current time
	currentTime := time.Now()

	// Extract year, month, and day
	year, month, day := currentTime.Date()

	key := fmt.Sprintf("%s/%d/%d/%d/%s", folder, year, month, day, dto.FileName)
	fmt.Println(key)

	// Perform the upload
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(p.S3_BUCKET_NAME),
		Key:    aws.String(key),
		Body:   dto.FileStream,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	result := UploadFileOutput{
		Name:   dto.FileName,
		Link:   fmt.Sprintf("https://%s.s3.amazonaws.com/%s", p.S3_BUCKET_NAME, key),
		Folder: folder,
		Type:   fileType,
	}
	fmt.Println(result)

	return result, nil
}
