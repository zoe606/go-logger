package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func NewUnauthorized(message string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func NewValidationError(message string) *AppError {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: message}
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}
