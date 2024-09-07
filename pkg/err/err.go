package err

import (
	"fmt"
)

// Define custom error types
type Error struct {
	Err            error
	HttpStatusCode int
}

// Error codes
const (
	ErrBadRequest = 400
	ErrNotFound   = 404
	ErrInternal   = 500
)

// NewError creates a new Error instance
func NewError(err error, httpStatusCode int) Error {
	return Error{
		Err:            err,
		HttpStatusCode: httpStatusCode,
	}
}

// GetError extracts the Error from a given error
func GetError(err error) Error {
	if e, ok := err.(Error); ok {
		return e
	}
	return NewError(err, ErrInternal) // Default to internal error if not recognized
}

// Implement the Error method for the Error type
func (e Error) Error() string {
	return e.Err.Error()
}

// Example usage of creating a new error
func NewBadRequest(msg string) Error {
	return NewError(fmt.Errorf(msg), ErrBadRequest)
}

func NewNotFound(msg string) Error {
	return NewError(fmt.Errorf(msg), ErrNotFound)
}

func NewInternalError(msg string) Error {
	return NewError(fmt.Errorf(msg), ErrInternal)
}
