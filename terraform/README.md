# Terraform Infrastructure for Vibe Storm

This directory contains the Terraform configuration files to deploy the Vibe Storm application infrastructure on AWS.

## Architecture Overview

The infrastructure is organized into the following modules:

1. **Network Module** - VPC, subnets, internet gateway, and S3 endpoint
2. **Security Module** - Security groups for all components
3. **Database Module** - RDS MySQL instance
4. **App Module** - ALB, S3 bucket, Route53, and ACM certificate
5. **Compute Module** - EC2 instances, Auto Scaling Group, and bastion host

## Infrastructure Components

The Terraform configuration creates the following AWS resources:

1. **VPC** with public and private subnets across multiple availability zones
2. **Internet Gateway** for internet access
3. **Security Groups** for all components
4. **MySQL RDS instance** in private subnet
5. **EC2 Bastion host** in public subnet for secure access
6. **EC2 Application instances** in private subnet with Auto Scaling Group
7. **Application Load Balancer** with HTTPS support
8. **S3 Bucket** with VPC endpoint and lifecycle rules
9. **Route53** hosted zone and DNS records
10. **ACM Certificate** for HTTPS
11. **Cognito User Pool** for authentication
12. **SSM Parameter Store** for secrets
13. **CloudWatch** for monitoring and logging
14. **IAM Roles** for EC2 and GitHub Actions
15. **DynamoDB** for Terraform state locking

## Prerequisites

1. [Terraform](https://www.terraform.io/downloads.html) installed
2. AWS CLI configured with appropriate credentials
3. SSH key pairs for EC2 instances

## Setup Instructions

1. **Create SSH Key Pairs**
   ```bash
   # Create key pair for bastion host and application servers
   ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa
   ```

2. **Configure Variables**
   Copy the example variables file and modify it with your settings:
   ```bash
   cp terraform.tfvars.example terraform.tfvars
   ```
   
   Edit `terraform.tfvars` to set your specific values, especially:
   - AWS region
   - Domain names
   - Database password
   - S3 bucket name
   - GitHub repository

3. **Initialize Terraform**
   ```bash
   terraform init
   ```

4. **Plan the Infrastructure**
   ```bash
   terraform plan
   ```

5. **Deploy the Infrastructure**
   ```bash
   terraform apply
   ```

6. **Destroy the Infrastructure** (when no longer needed)
   ```bash
   terraform destroy
   ```

## Security Features

- Application servers are in private subnets with no direct internet access
- Bastion host provides secure SSH access to private instances
- Security groups control traffic between components
- Database is accessible only from application servers and bastion
- HTTPS enforced through ALB redirect from HTTP
- IAM roles and policies for least privilege access
- SSM Parameter Store for secrets management
- Cognito for user authentication

## Cost Optimization

- Uses t3.micro instances (free tier eligible)
- No NAT Gateways (uses public subnet for outbound + S3 VPC Endpoint)
- S3 lifecycle rules for cost-effective storage
- DynamoDB for Terraform state locking (pay-per-request)

## CI/CD Integration

- GitHub Actions workflows for Terraform and application deployment
- OIDC integration for secure AWS credentials
- SSM Run Command for application deployment

## Accessing the Infrastructure

1. **Application**: Access via the domain name configured in Route53
2. **Database**: Access through the bastion host using SSH tunneling
3. **Application Servers**: Access through the bastion host only

## Customization

You can customize the infrastructure by modifying the variables in `terraform.tfvars`:
- Change instance types
- Adjust Auto Scaling Group parameters
- Modify network CIDR blocks
- Update security group rules
