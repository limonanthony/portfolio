package errordefs

import (
	"net/http"
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

type Conflict string

func (e Conflict) Error() string {
	return string(e)
}
func (e Conflict) StatusCode() int {
	return http.StatusConflict
}

type NotFound string

func (e NotFound) Error() string {
	return string(e)
}
func (e NotFound) StatusCode() int {
	return http.StatusNotFound
}

type BadRequest string

func (e BadRequest) Error() string {
	return string(e)
}
func (e BadRequest) StatusCode() int {
	return http.StatusBadRequest
}

type Unauthorized string

func (e Unauthorized) Error() string {
	return string(e)
}
func (e Unauthorized) StatusCode() int {
	return http.StatusUnauthorized
}

type Forbidden string

func (e Forbidden) Error() string {
	return string(e)
}
func (e Forbidden) StatusCode() int {
	return http.StatusForbidden
}

type InternalServerError string

func (e InternalServerError) Error() string {
	return string(e)
}
func (e InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}
