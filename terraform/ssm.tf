# SSM Parameter for Database Host
resource "aws_ssm_parameter" "db_host" {
  name  = "/${var.project_name}/${var.environment}/db/host"
  type  = "String"
  value = aws_db_instance.main.endpoint

  tags = {
    Name = "${var.project_name}-${var.environment}-db-host"
  }
}

# SSM Parameter for Database Name
resource "aws_ssm_parameter" "db_name" {
  name  = "/${var.project_name}/${var.environment}/db/name"
  type  = "String"
  value = aws_db_instance.main.db_name

  tags = {
    Name = "${var.project_name}-${var.environment}-db-name"
  }
}

# SSM Parameter for Database Username
resource "aws_ssm_parameter" "db_username" {
  name  = "/${var.project_name}/${var.environment}/db/username"
  type  = "String"
  value = aws_db_instance.main.username

  tags = {
    Name = "${var.project_name}-${var.environment}-db-username"
  }
}

# SSM Parameter for Database Password (SecureString)
resource "aws_ssm_parameter" "db_password" {
  name  = "/${var.project_name}/${var.environment}/db/password"
  type  = "SecureString"
  value = var.db_password

  tags = {
    Name = "${var.project_name}-${var.environment}-db-password"
  }
}

# SSM Parameter for LLM API Key (SecureString)
resource "aws_ssm_parameter" "llm_api_key" {
  name  = "/${var.project_name}/${var.environment}/llm/api-key"
  type  = "SecureString"
  value = "your-llm-api-key" # This should be replaced with actual API key

  tags = {
    Name = "${var.project_name}-${var.environment}-llm-api-key"
  }
}

# SSM Parameter for JWT Secret (SecureString)
resource "aws_ssm_parameter" "jwt_secret" {
  name  = "/${var.project_name}/${var.environment}/auth/jwt-secret"
  type  = "SecureString"
  value = "your-jwt-secret" # This should be replaced with actual secret

  tags = {
    Name = "${var.project_name}-${var.environment}-jwt-secret"
  }
}

# Outputs for SSM parameters
output "ssm_db_host_name" {
  description = "Name of the SSM parameter for database host"
  value       = aws_ssm_parameter.db_host.name
}

output "ssm_db_name_name" {
  description = "Name of the SSM parameter for database name"
  value       = aws_ssm_parameter.db_name.name
}

output "ssm_db_username_name" {
  description = "Name of the SSM parameter for database username"
  value       = aws_ssm_parameter.db_username.name
}

output "ssm_db_password_name" {
  description = "Name of the SSM parameter for database password"
  value       = aws_ssm_parameter.db_password.name
}

output "ssm_llm_api_key_name" {
  description = "Name of the SSM parameter for LLM API key"
  value       = aws_ssm_parameter.llm_api_key.name
}

output "ssm_jwt_secret_name" {
  description = "Name of the SSM parameter for JWT secret"
  value       = aws_ssm_parameter.jwt_secret.name
}
