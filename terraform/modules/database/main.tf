# RDS Subnet Group
resource "aws_db_subnet_group" "main" {
  name       = "${var.project_name}-${var.environment}-db-subnet-group"
  subnet_ids = var.private_subnet_ids

  tags = {
    Name = "${var.project_name}-${var.environment}-db-subnet-group"
  }
}

# RDS MySQL Instance
resource "aws_db_instance" "main" {
  identifier             = "${var.project_name}-${var.environment}-mysql"
  db_name                = var.db_name
  username               = var.db_username
  password               = var.db_password
  engine                 = "mysql"
  engine_version         = "8.0"
  instance_class         = var.db_instance_class
  allocated_storage      = 20
  storage_type           = "gp2"
  storage_encrypted      = true
  vpc_security_group_ids = [var.rds_security_group_id]
  db_subnet_group_name   = aws_db_subnet_group.main.name
  
  # Backup configuration
  backup_retention_period = 7
  backup_window           = "03:00-04:00"
  maintenance_window      = "sun:04:00-sun:05:00"
  
  # Multi-AZ for production (set to true for production)
  multi_az = false
  
  # Skip final snapshot for development (set to true for production)
  skip_final_snapshot = true
  
  # Enable deletion protection for production
  deletion_protection = false

  tags = {
    Name = "${var.project_name}-${var.environment}-mysql"
  }
}
