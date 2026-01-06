package errors

import (
	"fmt"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrBadRequest          = NewAppError(400, "Bad Request", nil)
	ErrUnauthorized        = NewAppError(401, "Unauthorized", nil)
	ErrForbidden           = NewAppError(403, "Forbidden", nil)
	ErrNotFound            = NewAppError(404, "Resource Not Found", nil)
	ErrConflict            = NewAppError(409, "Conflict", nil)
	ErrInternalServerError = NewAppError(500, "Internal Server Error", nil)
)

func NewValidationError(message string) *AppError {
	return NewAppError(400, message, nil)
}

func NewNotFoundError(resource string) *AppError {
	return NewAppError(404, fmt.Sprintf("%s not found", resource), nil)
}

func NewConflictError(message string) *AppError {
	return NewAppError(409, message, nil)
}
