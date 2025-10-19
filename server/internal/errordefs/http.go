package errordefs

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func ToHttpError(err error) error {
	if err == nil {
		return nil
	}

	if appErr, ok := IsAppError(err); ok {
		return huma.NewError(appErr.StatusCode(), appErr.Message())
	}

	return huma.NewError(http.StatusInternalServerError, err.Error())
}
