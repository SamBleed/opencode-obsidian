package errors

import (
	"fmt"
)

// AppError extiende el error estándar con contexto para auditoría
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func Wrap(code, msg string, err error) error {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

// Códigos estándar del Bunker
const (
	ErrInternal = "INTERNAL_ERROR"
	ErrInvalid  = "INVALID_INPUT"
	ErrNotFound = "NOT_FOUND"
	ErrSecurity = "SECURITY_VIOLATION"
)
