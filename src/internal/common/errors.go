package common

import (
	"fmt"
)

// Error types
type ErrorType string

const (
	ValidationError ErrorType = "validation_error"
	NotFoundError   ErrorType = "not_found_error"
	InternalError   ErrorType = "internal_error"
)

// AppError represents an application error
type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewValidationError creates a validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: message,
	}
}

// NewNotFoundError creates a not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFoundError,
		Message: message,
	}
}

// NewInternalError creates an internal error
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: message,
		Err:     err,
	}
}

// IsValidationError checks if error is a validation error
func IsValidationError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == ValidationError
	}
	return false
}

// IsNotFoundError checks if error is a not found error
func IsNotFoundError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == NotFoundError
	}
	return false
}

// IsInternalError checks if error is an internal error
func IsInternalError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == InternalError
	}
	return false
}
