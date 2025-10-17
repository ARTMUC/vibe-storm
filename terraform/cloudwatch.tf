# CloudWatch Log Group for Application Logs
resource "aws_cloudwatch_log_group" "app_logs" {
  name              = "/${var.project_name}/${var.environment}/app"
  retention_in_days = 30

  tags = {
    Name = "${var.project_name}-${var.environment}-app-logs"
  }
}

# CloudWatch Log Group for EC2 System Logs
resource "aws_cloudwatch_log_group" "ec2_system_logs" {
  name              = "/${var.project_name}/${var.environment}/ec2/system"
  retention_in_days = 30

  tags = {
    Name = "${var.project_name}-${var.environment}-ec2-system-logs"
  }
}

# CloudWatch Log Group for RDS Logs
resource "aws_cloudwatch_log_group" "rds_logs" {
  name              = "/${var.project_name}/${var.environment}/rds"
  retention_in_days = 30

  tags = {
    Name = "${var.project_name}-${var.environment}-rds-logs"
  }
}

# CloudWatch Alarm for EC2 CPU Utilization
resource "aws_cloudwatch_metric_alarm" "ec2_cpu_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-ec2-cpu-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance CPU utilization"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    InstanceId = aws_instance.bastion.id
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-ec2-cpu-alarm"
  }
}

# CloudWatch Alarm for EC2 Memory Utilization
resource "aws_cloudwatch_metric_alarm" "ec2_memory_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-ec2-memory-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "mem_used_percent"
  namespace           = "CWAgent"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance memory utilization"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    InstanceId = aws_instance.bastion.id
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-ec2-memory-alarm"
  }
}

# CloudWatch Alarm for EC2 Disk Utilization
resource "aws_cloudwatch_metric_alarm" "ec2_disk_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-ec2-disk-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "disk_used_percent"
  namespace           = "CWAgent"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors EC2 instance disk utilization"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    InstanceId = aws_instance.bastion.id
    path       = "/"
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-ec2-disk-alarm"
  }
}

# CloudWatch Alarm for RDS Storage Utilization
resource "aws_cloudwatch_metric_alarm" "rds_storage_utilization" {
  alarm_name          = "${var.project_name}-${var.environment}-rds-storage-utilization"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "FreeStorageSpace"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "1073741824" # 1 GB in bytes
  alarm_description   = "This metric monitors RDS storage space"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.main.id
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-rds-storage-alarm"
  }
}

# CloudWatch Alarm for ALB 5XX Errors
resource "aws_cloudwatch_metric_alarm" "alb_5xx_errors" {
  alarm_name          = "${var.project_name}-${var.environment}-alb-5xx-errors"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "HTTPCode_ELB_5XX_Count"
  namespace           = "AWS/ApplicationELB"
  period              = "60"
  statistic           = "Sum"
  threshold           = "10"
  alarm_description   = "This metric monitors ALB 5XX errors"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    LoadBalancer = aws_lb.main.arn_suffix
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-alb-5xx-alarm"
  }
}

# CloudWatch Alarm for ALB 4XX Errors
resource "aws_cloudwatch_metric_alarm" "alb_4xx_errors" {
  alarm_name          = "${var.project_name}-${var.environment}-alb-4xx-errors"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "HTTPCode_ELB_4XX_Count"
  namespace           = "AWS/ApplicationELB"
  period              = "60"
  statistic           = "Sum"
  threshold           = "50"
  alarm_description   = "This metric monitors ALB 4XX errors"
  alarm_actions       = []
  ok_actions          = []

  dimensions = {
    LoadBalancer = aws_lb.main.arn_suffix
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-alb-4xx-alarm"
  }
}

# Outputs for CloudWatch resources
output "cloudwatch_app_logs_group" {
  description = "Name of the CloudWatch log group for application logs"
  value       = aws_cloudwatch_log_group.app_logs.name
}

output "cloudwatch_ec2_system_logs_group" {
  description = "Name of the CloudWatch log group for EC2 system logs"
  value       = aws_cloudwatch_log_group.ec2_system_logs.name
}

output "cloudwatch_rds_logs_group" {
  description = "Name of the CloudWatch log group for RDS logs"
  value       = aws_cloudwatch_log_group.rds_logs.name
}
