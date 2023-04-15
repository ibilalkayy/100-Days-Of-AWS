resource "null_resource" "delete_instances" {
    triggers = {
        instance_id = "enter the instance id"
    }

    provisioner "local-exec" {
        command = "aws ec2 terminate-instances --instance-ids enter the instance id"
    }
}