# Cognito User Pool
resource "aws_cognito_user_pool" "main" {
  name = "${var.project_name}-${var.environment}-user-pool"

  # Set up basic authentication with email
  username_attributes = ["email"]
  
  # Set up password policy
  password_policy {
    minimum_length                   = 8
    require_lowercase                = true
    require_numbers                  = true
    require_symbols                  = true
    require_uppercase                = true
    temporary_password_validity_days = 7
  }

  # Set up email configuration
  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  # Set up account recovery
  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  # Set up MFA
  mfa_configuration = "OFF"

  # Set up verification message
  verification_message_template {
    default_email_option = "CONFIRM_WITH_CODE"
  }

  # Set up schema
  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true

    string_attribute_constraints {
      min_length = 0
      max_length = 2048
    }
  }

  tags = {
    Name = "${var.project_name}-${var.environment}-user-pool"
  }
}

# Cognito User Pool Client
resource "aws_cognito_user_pool_client" "main" {
  name         = "${var.project_name}-${var.environment}-client"
  user_pool_id = aws_cognito_user_pool.main.id

  # Set up allowed OAuth flows
  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_SRP_AUTH"
  ]

  # Set up token validity
  access_token_validity  = 60
  id_token_validity      = 60
  refresh_token_validity = 30

  token_validity_units {
    access_token  = "minutes"
    id_token      = "minutes"
    refresh_token = "days"
  }

  # Prevent user existence errors
  prevent_user_existence_errors = "ENABLED"

  # Set up supported identity providers
  supported_identity_providers = ["COGNITO"]

  # Callback URLs for web application
  callback_urls = [
    "https://${var.domain_name}/callback",
    "http://localhost:3000/callback"
  ]

  # Logout URLs
  logout_urls = [
    "https://${var.domain_name}/logout",
    "http://localhost:3000/logout"
  ]

  # Allowed OAuth scopes
  allowed_oauth_scopes = [
    "phone",
    "email",
    "openid",
    "profile"
  ]

  # Allowed OAuth flows
  allowed_oauth_flows = [
    "code"
  ]

  allowed_oauth_flows_user_pool_client = true

  # Generate secret for client
  generate_secret = false

  tags = {
    Name = "${var.project_name}-${var.environment}-client"
  }
}

# Cognito User Pool Domain
resource "aws_cognito_user_pool_domain" "main" {
  domain       = "${var.project_name}-${var.environment}"
  user_pool_id = aws_cognito_user_pool.main.id
}

# Outputs for Cognito resources
output "cognito_user_pool_id" {
  description = "ID of the Cognito User Pool"
  value       = aws_cognito_user_pool.main.id
}

output "cognito_user_pool_client_id" {
  description = "ID of the Cognito User Pool Client"
  value       = aws_cognito_user_pool_client.main.id
}

output "cognito_user_pool_domain" {
  description = "Domain of the Cognito User Pool"
  value       = aws_cognito_user_pool_domain.main.domain
}
