package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})

	if err != nil {
		log.Fatalf("failed to create a session %v\n", err)
		os.Exit(1)
	}

	serviceClient := ec2.New(session)

	parameters := &ec2.RunInstancesInput{
		ImageId:      aws.String("image-id"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	createInstance, err := serviceClient.RunInstances(parameters)
	if err != nil {
		log.Fatalf("failed to create an instance %v\n", err)
		os.Exit(1)
	}

	instanceID := *createInstance.Instances[0].InstanceId
	fmt.Printf("Created an EC2 instance and here is it's ID: %v\n", instanceID)
}
