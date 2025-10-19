package errordefs

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	code int
	msg  string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.msg)
}

func (e AppError) StatusCode() int {
	return e.code
}

func (e AppError) Message() string {
	return e.msg
}

func NewError(msg string, code int) AppError {
	return AppError{code: code, msg: msg}
}

func Error(code int, msg string) AppError {
	return AppError{code: code, msg: msg}
}

func Errorf(code int, format string, args ...interface{}) AppError {
	return AppError{code: code, msg: fmt.Sprintf(format, args...)}
}

func Forbidden(msg string) AppError {
	return Error(http.StatusForbidden, msg)
}

func Forbiddenf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusForbidden, format, args...)
}

func NotFound(msg string) AppError {
	return Error(http.StatusNotFound, msg)
}

func NotFoundf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusNotFound, format, args...)
}

func Unauthorized(msg string) AppError {
	return Error(http.StatusUnauthorized, msg)
}

func Unauthorizedf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusUnauthorized, format, args...)
}

func BadRequest(msg string) AppError {
	return Error(http.StatusBadRequest, msg)
}

func BadRequestf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusBadRequest, format, args...)
}

func Conflict(msg string) AppError {
	return Error(http.StatusConflict, msg)
}

func Conflictf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusConflict, format, args...)
}

func Internal(msg string) AppError {
	return Error(http.StatusInternalServerError, msg)
}

func Internalf(format string, args ...interface{}) AppError {
	return Errorf(http.StatusInternalServerError, format, args...)
}

func IsAppError(err error) (AppError, bool) {
	var appErr AppError
	ok := errors.As(err, &appErr)
	return appErr, ok
}
