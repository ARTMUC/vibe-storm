# Key Pair for Application Servers
resource "aws_key_pair" "app" {
  key_name   = var.key_pair_name
  public_key = file("~/.ssh/id_rsa.pub") # Update with your public key path
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

  vpc_security_group_ids = [var.app_security_group_id]

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

# Auto Scaling Group
resource "aws_autoscaling_group" "main" {
  name_prefix = "${var.project_name}-${var.environment}-asg-"

  vpc_zone_identifier = var.private_subnet_ids

  launch_template {
    id      = aws_launch_template.app.id
    version = "$Latest"
  }

  target_group_arns = [var.alb_target_group_arn]

  min_size = 1
  max_size = 2
  desired_capacity = 1

  health_check_type = "ELB"
  health_check_grace_period = 300

  tag {
    key                 = "Name"
    value               = "${var.project_name}-${var.environment}-app"
    propagate_at_launch = true
  }

  lifecycle {
    ignore_changes = [desired_capacity]
  }

  tags = [
    {
      key                 = "Environment"
      value               = var.environment
      propagate_at_launch = true
    },
    {
      key                 = "Project"
      value               = var.project_name
      propagate_at_launch = true
    }
  ]
}

# Bastion Host
resource "aws_instance" "bastion" {
  ami           = "ami-0c02fb55956c7d316" # Amazon Linux 2 AMI (change as needed for your region)
  instance_type = var.instance_type
  key_name      = aws_key_pair.app.key_name
  subnet_id     = var.public_subnet_ids[0]

  vpc_security_group_ids = [var.bastion_security_group_id]

  tags = {
    Name = "${var.project_name}-${var.environment}-bastion"
  }
}
