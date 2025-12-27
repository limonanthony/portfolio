package errordefs

import (
	"errors"
	"net/http"
)

type StatusError interface {
	error
	StatusCode() int
}

type AppError struct {
	code    int
	message string
	err     error
}

func (e *AppError) Error() string {
	if e.err != nil {
		return e.message + ": " + e.err.Error()
	}
	return e.message
}

func (e *AppError) StatusCode() int {
	return e.code
}

func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) Is(target error) bool {
	var t *AppError
	ok := errors.As(target, &t)
	if !ok {
		return false
	}

	return e.StatusCode() == t.StatusCode()
}

func NewAppError(msg string, code int, err error) *AppError {
	return &AppError{code: code, message: msg, err: err}
}

func NewConflict(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusConflict, err)
}

func NewNotFound(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusNotFound, err)
}

func NewBadRequest(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusBadRequest, err)
}

func NewUnauthorized(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusUnauthorized, err)
}

func NewForbidden(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusForbidden, err)
}

func NewInternalServerError(msg string, err error) *AppError {
	return NewAppError(msg, http.StatusInternalServerError, err)
}

var (
	ErrConflict            = NewConflict("conflict", nil)
	ErrNotFound            = NewNotFound("not found", nil)
	ErrBadRequest          = NewBadRequest("bad request", nil)
	ErrUnauthorized        = NewUnauthorized("unauthorized", nil)
	ErrForbidden           = NewForbidden("forbidden", nil)
	ErrInternalServerError = NewInternalServerError("internal server error", nil)
)
