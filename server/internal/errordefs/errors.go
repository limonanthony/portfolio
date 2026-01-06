package errordefs

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrApp          = errors.New("application error")
	ErrConflict     = errors.New("conflict")
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal server error")
)

type StatusError interface {
	error
	StatusCode() int
}

type AppError struct {
	code int
	msg  string
	err  error
}

func (e *AppError) Error() string {
	if e.msg != "" {
		return e.msg
	}
	if e.err != nil {
		return e.err.Error()
	}
	return "an unknown error occurred"
}

func (e *AppError) StatusCode() int {
	return e.code
}

func (e *AppError) Unwrap() error {
	return e.err
}

func NewConflict(msg string) *AppError {
	return &AppError{
		code: http.StatusConflict,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrConflict, ErrApp),
	}
}

func NewNotFound(msg string) *AppError {
	return &AppError{
		code: http.StatusNotFound,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrNotFound, ErrApp),
	}
}

func NewBadRequest(msg string) *AppError {
	return &AppError{
		code: http.StatusBadRequest,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrBadRequest, ErrApp),
	}
}

func NewUnauthorized(msg string) *AppError {
	return &AppError{
		code: http.StatusUnauthorized,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrUnauthorized, ErrApp),
	}
}

func NewForbidden(msg string) *AppError {
	return &AppError{
		code: http.StatusForbidden,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrForbidden, ErrApp),
	}
}

func NewInternal(msg string) *AppError {
	return &AppError{
		code: http.StatusInternalServerError,
		msg:  msg,
		err:  fmt.Errorf("%w: %w", ErrInternal, ErrApp),
	}
}
