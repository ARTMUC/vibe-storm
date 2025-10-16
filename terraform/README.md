# Terraform Infrastructure for Vibe Storm

This directory contains the Terraform configuration files to deploy the Vibe Storm application infrastructure on AWS.

## Infrastructure Components

The Terraform configuration creates the following AWS resources:

1. **VPC** with public and private subnets across multiple availability zones
2. **Internet Gateway** and **NAT Gateways** for internet access
3. **Security Groups** for all components
4. **MySQL RDS instance** in private subnet
5. **EC2 Bastion host** in public subnet for secure access
6. **EC2 Application instances** in private subnet with Auto Scaling Group
7. **Application Load Balancer** with HTTPS support
8. **S3 Bucket** with VPC endpoint
9. **Route53** hosted zone and DNS records

## Prerequisites

1. [Terraform](https://www.terraform.io/downloads.html) installed
2. AWS CLI configured with appropriate credentials
3. SSH key pairs for EC2 instances

## Setup Instructions

1. **Create SSH Key Pairs**
   ```bash
   # Create key pair for bastion host
   ssh-keygen -t rsa -b 4096 -f bastion
   
   # Create key pair for application servers
   ssh-keygen -t rsa -b 4096 -f app
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

## Architecture Overview

```
Internet
    ↓
Route53 DNS → ALB → EC2 Instances (Auto Scaling Group)
    ↓
VPC with Public and Private Subnets
    ↓
RDS MySQL (Private) ← Bastion Host (Public)
    ↓
S3 Bucket with VPC Endpoint
```

## Security Features

- Application servers are in private subnets with no direct internet access
- Bastion host provides secure SSH access to private instances
- Security groups control traffic between components
- Database is accessible only from application servers and bastion
- HTTPS enforced through ALB redirect from HTTP

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
