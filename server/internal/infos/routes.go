package infos

import (
	"context"
	"net/http"

	"github.com/limonanthony/portfolio/internal/router"
)

func RegisterRoutes(r *router.Router) {
	router.Register(r, router.Operation{
		OperationID:   "ping",
		Method:        http.MethodGet,
		Path:          "/ping",
		Summary:       "Ping",
		Description:   "Get ponged",
		Tags:          []string{"Infos"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *struct{}) (*pingResponse, error) {
		resp := &pingResponse{
			Body: "pong",
		}
		return resp, nil
	})
}
