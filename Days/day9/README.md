# Dockerization Odyssey Readme

## Introduction
Welcome to the Dockerization Odyssey! In this journey, we will explore the process of containerizing an application, unleashing its potential, and hosting it in the vast expanse of Amazon EC2. Follow along as Bilal Khan takes you through each step, from unveiling the application to setting sail into the clouds.

Here is the [video explanation](https://www.youtube.com/watch?v=Ha1s9Hhjc_Q) of it.

### Application Overview
The application under consideration is a simple yet powerful landing page and subscription form. You can explore the code on Bilal Khan's GitHub repository. To witness the magic locally, navigate to `localhost:8000` in your browser.

## Dockerization Decoded

### The Dockerfile
Our starship for this mission is the mighty Dockerfile. It serves as the blueprint guiding the construction of our containerized application. Let's dive into the key steps:

```docker
# Dockerfile

# Choosing the base image from Docker Hub
FROM golang:1.21.0.4-alpine3.8

# Creating a directory for our site
RUN mkdir /site

# Copying application files to the site directory
COPY . /site

# Switching to the site directory
WORKDIR /site

# Building the Go application
RUN go build -o main .

# Making the application executable and running it
CMD ["./main"]
```

In this Dockerfile, the Alpine version of the official Golang image from Docker Hub is used. The process involves creating a site directory, copying application files into it, building the Go application, and finally, running it. Think of it as assembling a spaceship for your application to journey through the container galaxy.

### Launching the Dockerization Process
With the Dockerfile ready, it's time to set your application on the path to containerized greatness. Run the following commands in your terminal:

```bash
# Building the Docker image
docker build -t my-awesome-app .

# Running the Docker container
docker run -p 8000:8000 my-awesome-app
```

Just like that, your application is encapsulated in a Docker container, ready to defy gravity and operate seamlessly across different environments.

## Sailing into the EC2 Waters

### Hosting on Amazon's Cloud
The journey doesn't end with Dockerization; it extends to the clouds, specifically Amazon EC2. Follow this brief guide to set sail to the cloud:

1. **Create an EC2 Instance:**
   Navigate to the AWS Management Console, launch an EC2 instance, and choose your preferred Amazon Machine Image (AMI).

2. **Configure Security Groups:**
   Ensure your security groups allow traffic on port 8000, where your application is sailing.

3. **Connect to Your EC2 Instance:**
   Use SSH to connect to your EC2 instance and prepare it for the application's arrival.

4. **Deploy Your Dockerized App:**
   Upload your Docker image to your EC2 instance, run the container, and expose the necessary ports.

5. **Access Your App in the Cloud:**
   Open your browser and navigate to your EC2 instance's public IP or DNS, appending port 8000. Witness your application shining in the cloud!

## Final Words

### Unleash Your Applications into the Cosmos
Dockerization is a pivotal step toward achieving application portability and scalability. With Docker, you can navigate the complex galaxies of software deployment with ease. So, set sail, containerize your creations, and let them soar into the celestial heights of the cloud!

Thank you for joining Bilal Khan on this Dockerization odyssey. Don't forget to like, share, and subscribe for more tech adventures. Until next time, happy coding! 🌌🚀

## Connect with Bilal Khan
- [YouTube](https://www.youtube.com/@coderoamer)
- [Twitter](https://www.x.com/ibilalkayy)
- [LinkedIn](https://www.linkedin.com/in/ibilalkayy)

If you have any questions or would like to share your own experiences, feel free to leave a comment below. Bilal is here to support and engage with you.

Happy Clouding! 🎉