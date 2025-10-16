package http

import (
	"net/http"
	"strings"
	"vibe-storm/internal/application/dto"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// HandlerFunc represents a handler function that returns DTO and error
type HandlerFunc func(c echo.Context) (interface{}, error)

// WrapHandler wraps a HandlerFunc to work with Echo
func WrapHandler(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := handlerFunc(c)
		if err != nil {
			// If it's a validation error, handle it appropriately
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				validationError := dto.NewValidationErrorResponse(
					c.Request().URL.Path,
					validationErrors,
				)
				return c.JSON(validationError.ToHTTPStatus(), validationError)
			}

			// If it's already a structured error, use it
			if structuredErr, ok := err.(*dto.StructuredError); ok {
				return c.JSON(structuredErr.Response.ToHTTPStatus(), structuredErr.Response)
			}

			// Otherwise, create a generic internal error
			internalError := &dto.StructuredError{
				Response: dto.NewStructuredErrorResponse(
					dto.ErrCodeInternalError,
					c.Request().URL.Path,
				),
			}
			return c.JSON(internalError.Response.ToHTTPStatus(), internalError.Response)
		}

		return c.JSON(http.StatusOK, result)
	}
}

// Helper functions for common error responses

// BadRequestError creates a bad request error response
func BadRequestError(c echo.Context, code string) *dto.StructuredErrorResponse {
	return &dto.StructuredErrorResponse{
		Error:   true,
		Code:    code,
		Message: dto.ErrorDefinitions[code].Message,
		Path:    c.Request().URL.Path,
	}
}

// NotFoundError creates a not found error response
func NotFoundError(c echo.Context) *dto.StructuredErrorResponse {
	return &dto.StructuredErrorResponse{
		Error:   true,
		Code:    dto.ErrCodeNotFound,
		Message: dto.ErrorDefinitions[dto.ErrCodeNotFound].Message,
		Path:    c.Request().URL.Path,
	}
}

// InternalError creates an internal server error response
func InternalError(c echo.Context) *dto.StructuredErrorResponse {
	return &dto.StructuredErrorResponse{
		Error:   true,
		Code:    dto.ErrCodeInternalError,
		Message: dto.ErrorDefinitions[dto.ErrCodeInternalError].Message,
		Path:    c.Request().URL.Path,
	}
}

// ValidationError creates a validation error response
func ValidationError(c echo.Context, validationErrors validator.ValidationErrors) *dto.StructuredErrorResponse {
	details := make(map[string][]string)

	for _, err := range validationErrors {
		field := strings.ToLower(err.Field())
		details[field] = append(details[field], getValidationErrorMessage(err))
	}

	return &dto.StructuredErrorResponse{
		Error:      true,
		Code:       dto.ErrCodeValidationError,
		Message:    dto.ErrorDefinitions[dto.ErrCodeValidationError].Message,
		Path:       c.Request().URL.Path,
		Validation: details,
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
