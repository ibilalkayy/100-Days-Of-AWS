package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("your-region"))
	if err != nil {
		log.Fatalf("failed to load the SDK %v", err)
	}

	client := s3.NewFromConfig(cfg)

	err = UploadFile(client, "your-bucket-name", "filename-to-be-uploaded")
	if err != nil {
		log.Fatalf("failed to upload the file %v", err)
	}

	fmt.Println("file is successfully uploaded")
}

func UploadFile(client *s3.Client, bucket, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	}

	_, err = client.PutObject(context.TODO(), input)
	return err
}
