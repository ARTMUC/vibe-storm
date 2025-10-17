# Key Pair for Application Servers
resource "aws_key_pair" "app" {
  key_name   = "${var.project_name}-${var.environment}-app-key"
  public_key = file("${path.module}/app.pub")
}

# Launch Template for Application Servers
resource "aws_launch_template" "app" {
  name_prefix   = "${var.project_name}-${var.environment}-app-"
  image_id      = "ami-0c02fb55956c7d316" # Amazon Linux 2 AMI (change as needed for your region)
  instance_type = var.instance_type
  key_name      = aws_key_pair.app.key_name
  iam_instance_profile {
    name = aws_iam_instance_profile.ec2_profile.name
  }

  vpc_security_group_ids = [aws_security_group.app.id]

  # User data script to install and start the application
  user_data = base64encode(<<-EOF
#!/bin/bash
yum update -y
yum install -y docker
systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user
# Add your application deployment commands here
# For example:
# docker run -d -p 80:8080 your-app-image
EOF
  )

  tag_specifications {
    resource_type = "instance"
    tags = {
      Name = "${var.project_name}-${var.environment}-app"
    }
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-app-launch-template"
  }
}
