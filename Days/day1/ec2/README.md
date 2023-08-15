# **EC2 Repository**

This repository contains a collection of files for managing and working with Amazon EC2 instances using both the Go programming language and Terraform. The purpose of this repository is to demonstrate two different approaches to creating, terminating, and waiting for the instances to run on AWS.

## **Repository Structure**

    ec2-repository/
    │
    ├── main.go
    ├── terraform/
    │   ├── create_instance.tf
    │   ├── terminate_instance.tf
    │   └── wait_instance_running.tf
    └── README.md

### **main.go**

The **main.go** file is a Go program that interacts with the AWS SDK for Go to perform the following tasks:

- Create an Amazon EC2 instance
- Terminate the created instance
- Wait until the instance is in a running state
- To run the main.go file, you'll need the Go programming language installed on your system. You can follow the official [Go installation guide](https://go.dev/doc/install) to get started.

### **Terraform**

Inside the ec2 directory there are three directories that contain three separate Terraform configuration files for performing the same tasks as the **main.go** file:

- **create/main.tf:** Creates an Amazon EC2 instance
- **terminate/main.tf:** Terminates the specified instance
- **wait/main.tf:** Waits until the specified instance is in a running state

To use these Terraform files, you'll need to have Terraform installed on your system. You can follow the official [Terraform installation guide](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) to get started.

## **Usage**

### **Go**

To execute the main.go file, navigate to the root directory of this repository and run the following command:

    go run main.go

## **Terraform**

To apply the Terraform configuration files, follow these steps:

Navigate to one of the directories in which the configuration file is present.

- Initialize the Terraform working directory:

      terraform init

- Apply the configuration files:

      terraform apply

- To destroy the created resources when you are done, run:

      terraform destroy

## **License**

This project is licensed under the MIT License.

## **Contributing**

Please feel free to submit issues and pull requests for improvements and bug fixes.