package dto

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Error codes following AWS-like naming convention
const (
	// Validation Errors (400-499)
	ErrCodeValidationError  = "VALIDATION_ERROR"
	ErrCodeInvalidParameter = "INVALID_PARAMETER"
	ErrCodeMissingParameter = "MISSING_PARAMETER"
	ErrCodeInvalidRequest   = "INVALID_REQUEST"
	ErrCodeUnauthorized     = "UNAUTHORIZED"
	ErrCodeForbidden        = "FORBIDDEN"
	ErrCodeNotFound         = "NOT_FOUND"
	ErrCodeConflict         = "CONFLICT"
	ErrCodeTooManyRequests  = "TOO_MANY_REQUESTS"

	// Server Errors (500-599)
	ErrCodeInternalError        = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable   = "SERVICE_UNAVAILABLE"
	ErrCodeDatabaseError        = "DATABASE_ERROR"
	ErrCodeExternalServiceError = "EXTERNAL_SERVICE_ERROR"

	// Business Logic Errors (400-499)
	ErrCodeUserAlreadyExists  = "USER_ALREADY_EXISTS"
	ErrCodeUserNotFound       = "USER_NOT_FOUND"
	ErrCodeInvalidCredentials = "INVALID_CREDENTIALS"
	ErrCodeAccountDisabled    = "ACCOUNT_DISABLED"
	ErrCodeEmailNotVerified   = "EMAIL_NOT_VERIFIED"
)

// Error definitions with HTTP status codes and messages
var ErrorDefinitions = map[string]ErrorDefinition{
	// Validation Errors
	ErrCodeValidationError: {
		Code:       ErrCodeValidationError,
		StatusCode: http.StatusBadRequest,
		Message:    "Request validation failed",
	},

	ErrCodeInvalidParameter: {
		Code:       ErrCodeInvalidParameter,
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid parameter value",
	},

	ErrCodeMissingParameter: {
		Code:       ErrCodeMissingParameter,
		StatusCode: http.StatusBadRequest,
		Message:    "Missing required parameter",
	},

	ErrCodeInvalidRequest: {
		Code:       ErrCodeInvalidRequest,
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid request format",
	},

	ErrCodeUnauthorized: {
		Code:       ErrCodeUnauthorized,
		StatusCode: http.StatusUnauthorized,
		Message:    "Authentication required",
	},

	ErrCodeForbidden: {
		Code:       ErrCodeForbidden,
		StatusCode: http.StatusForbidden,
		Message:    "Access denied",
	},

	ErrCodeNotFound: {
		Code:       ErrCodeNotFound,
		StatusCode: http.StatusNotFound,
		Message:    "Resource not found",
	},

	ErrCodeConflict: {
		Code:       ErrCodeConflict,
		StatusCode: http.StatusConflict,
		Message:    "Resource conflict",
	},

	ErrCodeTooManyRequests: {
		Code:       ErrCodeTooManyRequests,
		StatusCode: http.StatusTooManyRequests,
		Message:    "Too many requests",
	},

	// Server Errors
	ErrCodeInternalError: {
		Code:       ErrCodeInternalError,
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	},

	ErrCodeServiceUnavailable: {
		Code:       ErrCodeServiceUnavailable,
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service temporarily unavailable",
	},

	ErrCodeDatabaseError: {
		Code:       ErrCodeDatabaseError,
		StatusCode: http.StatusInternalServerError,
		Message:    "Database operation failed",
	},

	ErrCodeExternalServiceError: {
		Code:       ErrCodeExternalServiceError,
		StatusCode: http.StatusBadGateway,
		Message:    "External service error",
	},

	// Business Logic Errors
	ErrCodeUserAlreadyExists: {
		Code:       ErrCodeUserAlreadyExists,
		StatusCode: http.StatusConflict,
		Message:    "User already exists",
	},

	ErrCodeUserNotFound: {
		Code:       ErrCodeUserNotFound,
		StatusCode: http.StatusNotFound,
		Message:    "User not found",
	},

	ErrCodeInvalidCredentials: {
		Code:       ErrCodeInvalidCredentials,
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid credentials",
	},

	ErrCodeAccountDisabled: {
		Code:       ErrCodeAccountDisabled,
		StatusCode: http.StatusForbidden,
		Message:    "Account is disabled",
	},

	ErrCodeEmailNotVerified: {
		Code:       ErrCodeEmailNotVerified,
		StatusCode: http.StatusForbidden,
		Message:    "Email address not verified",
	},
}

// ErrorDefinition represents an error definition
type ErrorDefinition struct {
	Code       string
	StatusCode int
	Message    string
}

// GetErrorDefinition returns the error definition for a given error code
func GetErrorDefinition(code string) (ErrorDefinition, bool) {
	def, exists := ErrorDefinitions[code]
	return def, exists
}

// StructuredError represents a structured error that implements the error interface
type StructuredError struct {
	Response StructuredErrorResponse
}

// Error implements the error interface
func (e *StructuredError) Error() string {
	return e.Response.Message
}

// StructuredErrorResponse represents a structured error response
type StructuredErrorResponse struct {
	Error      bool                   `json:"error" example:"true"`
	Code       string                 `json:"code" example:"VALIDATION_ERROR"`
	Message    string                 `json:"message" example:"Request validation failed"`
	Path       string                 `json:"path" example:"/api/v1/users"`
	RequestID  string                 `json:"request_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174000"`
	Details    map[string]interface{} `json:"details,omitempty"`
	Validation map[string][]string    `json:"validation,omitempty"`
}

// NewStructuredErrorResponse creates a new structured error response
func NewStructuredErrorResponse(code, path string) StructuredErrorResponse {
	def, exists := ErrorDefinitions[code]
	if !exists {
		// Default to internal error if code not found
		def = ErrorDefinitions[ErrCodeInternalError]
	}

	return StructuredErrorResponse{
		Error:   true,
		Code:    code,
		Message: def.Message,
		Path:    path,
	}
}

// NewStructuredErrorResponseWithDetails creates a structured error response with additional details
func NewStructuredErrorResponseWithDetails(code, path string, details map[string]interface{}) StructuredErrorResponse {
	response := NewStructuredErrorResponse(code, path)
	response.Details = details
	return response
}

// NewValidationErrorResponse creates a structured validation error response
func NewValidationErrorResponse(path string, validationErrors validator.ValidationErrors) StructuredErrorResponse {
	details := make(map[string][]string)

	for _, err := range validationErrors {
		field := strings.ToLower(err.Field())
		details[field] = append(details[field], getValidationErrorMessage(err))
	}

	return StructuredErrorResponse{
		Error:      true,
		Code:       ErrCodeValidationError,
		Message:    ErrorDefinitions[ErrCodeValidationError].Message,
		Path:       path,
		Validation: details,
	}
}

// ToHTTPStatus returns the HTTP status code for the error
func (r *StructuredErrorResponse) ToHTTPStatus() int {
	if def, exists := ErrorDefinitions[r.Code]; exists {
		return def.StatusCode
	}
	return http.StatusInternalServerError
}
