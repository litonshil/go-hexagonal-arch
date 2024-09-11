// Package errors holds custom error definition
package errors

import (
	"fmt"
	"net/http"
)

// Error implements the error interface
type Error interface {
	error
	Status() int
}

// CustomError is a custom error type
type CustomError struct {
	// StatusCode is the http status code
	StatusCode int `json:"-"`
	// Message is the error message to display
	Message string `json:"error,omitempty"`
	// Details is the error details for debugging
	Details interface{} `json:"details,omitempty"`
}

// ForbiddenError is mainly used for unauthorized access
// StatusCode is 403
func ForbiddenError(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusForbidden,
		Message:    msg,
	}
}

// InvalidRequestParsingError is used for invalid parsing of request body
// StatusCode is 400
func InvalidRequestParsingError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid request",
		Details:    err.Error(),
	}
}

// BadRequest is used to indicate that the request is invalid
// StatusCode is 400
func BadRequest(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

// PreconditionFailed is used to indicate that the request is invalid anymore
// StatusCode is 412
func PreconditionFailed(msg string) CustomError {
	return CustomError{
		StatusCode: http.StatusPreconditionFailed,
		Message:    msg,
	}
}

// InternalDBError is used to indicate that an internal database error occurred
// StatusCode is 500
func InternalDBError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal database error",
		Details:    err.Error(),
	}
}

// NoEntityError indicates that the entity is not found
// StatusCode is 404
func NoEntityError(entity string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    fmt.Sprintf("no '%s' found", entity),
	}
}

// ValidationError indicates that the request data is failed to validate
// StatusCode is 400
func ValidationError(err string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "validation failed",
		Details:    err,
	}
}

// DBMigrationError is used to indicate that an internal database migration error occurred
func DBMigrationError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "database migration error",
		Details:    err.Error(),
	}
}

func InternalError(err error) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal error",
		Details:    err.Error(),
	}
}

// Status returns the status code
func (e CustomError) Status() int {
	return e.StatusCode
}

// Error returns the error message
func (e CustomError) Error() string {
	return e.Message
}
