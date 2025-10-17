# Terraform configuration for Vibe Storm application
# This file contains general configuration and data sources

# Get the current AWS caller identity
data "aws_caller_identity" "current" {}

# Get the current AWS region
data "aws_region" "current" {}

# Get the current AWS availability zones
data "aws_availability_zones" "available" {
  state = "available"
}

# Get the AMI ID for Amazon Linux 2
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# Network Module
module "network" {
  source = "./modules/network"

  project_name         = var.project_name
  environment          = var.environment
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
  aws_region           = var.aws_region
}

# Security Module
module "security" {
  source = "./modules/security"

  project_name = var.project_name
  environment  = var.environment
  vpc_id       = module.network.vpc_id
}

# Database Module
module "database" {
  source = "./modules/database"

  project_name         = var.project_name
  environment          = var.environment
  vpc_id               = module.network.vpc_id
  private_subnet_ids   = module.network.private_subnet_ids
  rds_security_group_id = module.security.rds_security_group_id
  db_instance_class    = var.db_instance_class
  db_name              = var.db_name
  db_username          = var.db_username
  db_password          = var.db_password
}

# App Module
module "app" {
  source = "./modules/app"

  project_name           = var.project_name
  environment            = var.environment
  domain_name            = var.domain_name
  hosted_zone_domain     = var.hosted_zone_domain
  public_subnet_ids      = module.network.public_subnet_ids
  alb_security_group_id  = module.security.alb_security_group_id
  app_security_group_id  = module.security.app_security_group_id
  rds_endpoint           = module.database.rds_endpoint
  rds_database_name      = module.database.rds_database_name
  s3_bucket_name         = var.s3_bucket_name
}

# Compute Module
module "compute" {
  source = "./modules/compute"

  project_name           = var.project_name
  environment            = var.environment
  instance_type          = var.instance_type
  key_pair_name          = var.key_pair_name
  vpc_id                 = module.network.vpc_id
  public_subnet_ids      = module.network.public_subnet_ids
  private_subnet_ids     = module.network.private_subnet_ids
  app_security_group_id  = module.security.app_security_group_id
  alb_target_group_arn   = module.app.target_group_arn
  bastion_security_group_id = module.security.bastion_security_group_id
}
