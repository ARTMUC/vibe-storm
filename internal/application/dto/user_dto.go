package dto

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// CreateUserRequest represents the request DTO for creating a user
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email" example:"user@example.com"`
	Username  string `json:"username" validate:"required,min=3,max=20,alphanum" example:"johndoe"`
	Password  string `json:"password" validate:"required,min=8,containsany=!@#$%^&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz" example:"password123!"`
	FirstName string `json:"first_name" validate:"required,min=1,max=100" example:"John"`
	LastName  string `json:"last_name" validate:"required,min=1,max=100" example:"Doe"`
}

// UpdateUserRequest represents the request DTO for updating a user
type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty" validate:"omitempty,email" example:"user@example.com"`
	Username  *string `json:"username,omitempty" validate:"omitempty,min=3,max=20,alphanum" example:"johndoe"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=1,max=100" example:"John"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=1,max=100" example:"Doe"`
}

// UserResponse represents the response DTO for user data
type UserResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Email     string    `json:"email" example:"user@example.com"`
	Username  string    `json:"username" example:"johndoe"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	IsActive  bool      `json:"is_active" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// UsersResponse represents the response DTO for multiple users
type UsersResponse struct {
	Users      []UserResponse `json:"users"`
	Page       int            `json:"page" example:"1"`
	PerPage    int            `json:"per_page" example:"10"`
	TotalCount int64          `json:"total_count" example:"100"`
	TotalPages int            `json:"total_pages" example:"10"`
}

// validateStrongPassword validates password strength requirements
func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check minimum length
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if !hasUpper {
		return false
	}

	// Check for at least one lowercase letter
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	if !hasLower {
		return false
	}

	// Check for at least one digit
	hasDigit := strings.ContainsAny(password, "0123456789")
	if !hasDigit {
		return false
	}

	// Check for at least one special character
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*")
	if !hasSpecial {
		return false
	}

	return true
}

// validateUsernameFormat validates username format
func validateUsernameFormat(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// Check length
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	// Check for only alphanumeric characters and underscores
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '_') {
			return false
		}
	}

	return true
}

// ToUserResponse converts domain user to response DTO
func ToUserResponse(user interface{}) UserResponse {
	// This would be implemented when we have the actual domain model
	// For now, return a placeholder
	return UserResponse{}
}

// ToUsersResponse converts multiple users to response DTO
func ToUsersResponse(users interface{}, page, perPage int, totalCount int64) UsersResponse {
	// This would be implemented when we have the actual domain model
	totalPages := int((totalCount + int64(perPage) - 1) / int64(perPage))
	if totalPages == 0 {
		totalPages = 1
	}

	return UsersResponse{
		Users:      []UserResponse{},
		Page:       page,
		PerPage:    perPage,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}
}

// SigninRequest represents the request DTO for user signin
type SigninRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123!"`
}

// SignupRequest represents the request DTO for user signup
type SignupRequest struct {
	Email     string `json:"email" validate:"required,email" example:"user@example.com"`
	Username  string `json:"username" validate:"required,min=3,max=20,alphanum" example:"johndoe"`
	Password  string `json:"password" validate:"required,min=8,containsany=!@#$%^&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz" example:"password123!"`
	FirstName string `json:"first_name" validate:"required,min=1,max=100" example:"John"`
	LastName  string `json:"last_name" validate:"required,min=1,max=100" example:"Doe"`
}

// RefreshTokenRequest represents the request DTO for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// AuthResponse represents the response DTO for authentication
type AuthResponse struct {
	AccessToken  string       `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string       `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenType    string       `json:"token_type" example:"Bearer"`
	ExpiresAt    time.Time    `json:"expires_at" example:"2023-01-01T00:00:00Z"`
	User         UserResponse `json:"user"`
}

// TokenResponse represents the response DTO for token operations
type TokenResponse struct {
	AccessToken string    `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenType   string    `json:"token_type" example:"Bearer"`
	ExpiresAt   time.Time `json:"expires_at" example:"2023-01-01T00:00:00Z"`
}
