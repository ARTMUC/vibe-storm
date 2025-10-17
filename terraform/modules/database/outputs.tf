output "rds_endpoint" {
  description = "Endpoint of the RDS instance"
  value       = aws_db_instance.main.endpoint
}

output "rds_database_name" {
  description = "Database name of the RDS instance"
  value       = aws_db_instance.main.db_name
}
