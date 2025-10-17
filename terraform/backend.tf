# Terraform Backend Configuration
terraform {
  backend "s3" {
    bucket         = var.s3_bucket_name
    key            = "terraform.tfstate"
    region         = var.aws_region
    dynamodb_table = "${var.project_name}-${var.environment}-terraform-state-lock"
    encrypt        = true
  }
}
