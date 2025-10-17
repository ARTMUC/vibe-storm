# Terraform Refactoring Todo List

## Analysis
- [x] Review existing Terraform implementation
- [x] Identify gaps between current implementation and requirements

## Refactoring Tasks
- [x] Update VPC configuration to remove NAT Gateways (cost optimization)
- [x] Refactor security groups to match requirements
- [x] Update EC2 configuration for single instance + ASG option
- [x] Add SSM IAM role for EC2 instances
- [x] Add Cognito integration for authentication
- [x] Add SSM Parameter Store for secrets
- [x] Add CloudWatch monitoring and logging
- [x] Add GitHub Actions CI/CD configuration
- [x] Add OIDC GitHub Actions → AWS integration
- [x] Update ALB configuration for HTTPS with ACM certificate
- [x] Add S3 lifecycle rules
- [x] Organize code into modules (network, compute, database, app, security)
- [x] Update variables for flexibility
- [x] Update outputs as needed
- [x] Update documentation

## Summary
The Terraform implementation has been successfully refactored to meet all the requirements for the MVP of a Miro-like board with AI. The infrastructure is now:
- Secure with proper network isolation
- Cost-optimized with free-tier eligible resources
- Modular for better maintainability
- Integrated with CI/CD pipelines
- Equipped with monitoring and logging
