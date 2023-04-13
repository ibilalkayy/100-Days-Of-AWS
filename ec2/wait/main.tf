provider "aws" {
    region = "ap-south-1"
}

resource "aws_instance" "example" {
    ami = ""
    instance_type = "t2.micro"
}

data "aws_instance" "example_running" {
    filter {
        name = "instance-state-name"
        values = ["running"]
    }

    instance_id = aws_instance.example.id

    depends_on = [aws_instance.example]
}

output "instance_id" {
    value = data.aws_instance.example_running.id
}