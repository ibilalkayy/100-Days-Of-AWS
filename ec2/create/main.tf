provider "aws" {
    region = "ap-south-1"
}

resource "aws_instance" "example" {
    ami = ""
    instance_type = "t2.micro"

    tags = {
        Name = "testing-instance"
    }
}