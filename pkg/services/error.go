package services

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("conflict")                  // Indicates a conflict with current state
	ErrValidation          = errors.New("validation error")          // Indicates a validation error
	ErrDatabase            = errors.New("database error")            // Indicates a generic database error
	ErrUnauthorizedAccess  = errors.New("unauthorized access")       // Indicates unauthorized access attempt
	ErrServiceUnavailable  = errors.New("service unavailable")       // Indicates service is temporarily unavailable
	ErrTimeout             = errors.New("request timeout")           // Indicates request timeout
	ErrUnexpected          = errors.New("unexpected error occurred") // Indicates an unexpected internal server error
	ErrNotFound            = errors.New("not found")
)
