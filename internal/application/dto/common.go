package dto

import (
	"github.com/go-playground/validator/v10"
)

// Base response structures
type HealthResponse struct {
	Status  string `json:"status" example:"healthy"`
	Service string `json:"service" example:"VibeStorm"`
	Version string `json:"version" example:"1.0.0"`
	Env     string `json:"env" example:"development"`
}

type HomeResponse struct {
	Message string `json:"message" example:"Welcome to VibeStorm"`
	Version string `json:"version" example:"1.0.0"`
	Status  string `json:"status" example:"running"`
}

// Error response structures
type ErrorResponse struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Invalid request data"`
	Path    string `json:"path" example:"/api/v1/users"`
	Code    string `json:"code,omitempty" example:"VALIDATION_ERROR"`
}

type ValidationErrorResponse struct {
	Error   bool                `json:"error" example:"true"`
	Message string              `json:"message" example:"Validation failed"`
	Path    string              `json:"path" example:"/api/v1/users"`
	Code    string              `json:"code" example:"VALIDATION_ERROR"`
	Details map[string][]string `json:"details" example:"email:invalid email format;password:password too weak"`
}

// Pagination structures
type PaginationMeta struct {
	Page       int   `json:"page" example:"1"`
	PerPage    int   `json:"per_page" example:"10"`
	TotalCount int64 `json:"total_count" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
}

// Global validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register custom validations
	validate.RegisterValidation("strong_password", validateStrongPassword)
	validate.RegisterValidation("username_format", validateUsernameFormat)
}

// ValidateStruct validates a struct using the global validator
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message, path string) ErrorResponse {
	return ErrorResponse{
		Error:   true,
		Message: message,
		Path:    path,
	}
}

// getValidationErrorMessage returns a human-readable validation error message
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short (minimum " + err.Param() + " characters)"
	case "max":
		return "Value is too long (maximum " + err.Param() + " characters)"
	case "alphanum":
		return "Only alphanumeric characters are allowed"
	case "strong_password":
		return "Password must contain at least 8 characters with uppercase, lowercase, number, and special character"
	case "username_format":
		return "Username must be 3-20 characters long and contain only letters, numbers, and underscores"
	default:
		return "Invalid value"
	}
}
