# **To-Do List App**

This repository contains the source code for a simple To-Do List application hosted on an AWS S3 Bucket. The application allows users to add new tasks, mark tasks as completed, and delete tasks from the list.

## **Table of Contents**

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Technologies](#technologies)
- [Support](#support)

## **Features**

- Add new tasks to the list
- Mark tasks as completed
- Delete tasks from the list

## **Prerequisites**

- AWS account with S3 bucket permissions
- Basic understanding of HTML, CSS, and JavaScript

## **Installation**

- Clone the repository or download the source code.

      git clone https://github.com/ibilalkayy/aws.git

- Navigate to the project directory.

      cd s3bucket

- Upload the contents of the project directory to your S3 bucket.

- Configure your S3 bucket to act as a static website, and make sure to set the **index.html** file as the index document.

- Make your S3 bucket public (or grant appropriate permissions) so that users can access your To-Do List application.

## **Usage**

- Visit the URL of your S3 bucket's static website (e.g., http://your-bucket-name.s3-website.region.amazonaws.com).

- Enter a task description in the input field and click the "Add" button to add a task to the list.

- Click on a task to mark it as completed or uncompleted.

- To delete a task, click on the task while holding the Ctrl (or Cmd on Mac) key.

## **Technologies**

- HTML
- CSS
- JavaScript
- AWS S3

## **Support**

If you encounter any issues or need help with the To-Do List app, please open an issue in the repository or contact the maintainers directly.