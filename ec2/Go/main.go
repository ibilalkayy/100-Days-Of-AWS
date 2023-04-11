package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	instanceID, err := createEC2Instance(serviceClient)
	if err != nil {
		log.Fatalf("failed to create an instance %v", err)
	}

	err = waitUntilInstanceRunning(serviceClient, instanceID)
	if err != nil {
		log.Fatalf("failed to run an EC2 instance %v", err)
	}

	time.Sleep(10 * time.Second)

	err = terminateEC2Instance(serviceClient, instanceID)
	if err != nil {
		log.Fatalf("failed to terminate an EC2 instance %v", err)
	}
}

func createEC2Instance(serviceClient *ec2.EC2) (string, error) {
	parameters := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-05f96ebf267205daa"),
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

	return instanceID, nil
}

func waitUntilInstanceRunning(svc *ec2.EC2, instanceID string) error {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	err := svc.WaitUntilInstanceRunning(input)
	if err != nil {
		return err
	}

	fmt.Printf("EC2 instance with ID %s is running\n", instanceID)

	return nil
}

func terminateEC2Instance(svc *ec2.EC2, instanceID string) error {
	input := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	_, err := svc.TerminateInstances(input)
	if err != nil {
		return err
	}

	fmt.Printf("EC2 instance with an ID %s is terminated\n", instanceID)

	return nil
}
