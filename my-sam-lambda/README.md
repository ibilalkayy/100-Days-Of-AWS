# **Simple SAM Lambda Function**

This repository contains a simple AWS Serverless Application Model (SAM) application with a Lambda function and an API Gateway. The Lambda function is written in Node.js and is triggered by an HTTP GET request to the /hello path of the API Gateway.

## **Prerequisites**

To deploy and run this serverless application, you need the following:

- An AWS account
- AWS CLI installed and configured with your AWS account credentials
- AWS SAM CLI installed
- NodeJS installed

## **Project Structure**

The project has the following structure:

    .
    ├── README.md
    ├── hello-world
    │   └── app.js
    └── template.yaml

- **README.md:** This file, which contains the instructions to set up, deploy, and test the serverless application.
- **hello-world/:** Directory containing the Lambda function's source code (**app.js**).
- **template.yaml:** The AWS SAM template defining the serverless application's resources and configurations.

## **Deployment**

Follow these steps to deploy the serverless application:

1. Clone the repository:

       git clone https://github.com/ibilalkayy/aws.git
       cd my-sam-lambda

2. Build the serverless application using the SAM CLI:

       sam build

3. Deploy the serverless application using the SAM CLI with guided prompts:

       sam deploy --guided

Provide the necessary information when prompted, such as the stack name, AWS region, and whether to allow or disallow rollback on failure.

4. After the deployment is complete, the SAM CLI will output the API Gateway endpoint URL, which you can find under the Outputs section.

## **Testing**

To test the serverless application, send an HTTP GET request to the API Gateway endpoint URL:

    curl https://your-api-gateway-url/Prod/hello/

Replace **your-api-gateway-url** with the actual API Gateway endpoint URL from the deployment output. The response should be:

    {
    "Hello from Lambda!"
    }

## **Cleanup**

To remove the serverless application and its resources, delete the CloudFormation stack using the AWS CLI:

    aws cloudformation delete-stack --stack-name your-stack-name

Replace **your-stack-name** with the name you provided during deployment.

## **License**

This project is licensed under the terms of the MIT License.