package errordefs

import (
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func ToHttpError(err error) error {
	if err == nil {
		return nil
	}

	var statusErr StatusError
	if errors.As(err, &statusErr) {
		return huma.NewError(statusErr.StatusCode(), err.Error())
	}

	return huma.NewError(http.StatusInternalServerError, err.Error())
}
