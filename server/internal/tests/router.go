package tests

import (
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/limonanthony/portfolio/internal/router"
)

func NewRouter(t *testing.T) (*router.Router, humatest.TestAPI) {
	_, api := humatest.New(t)

	return router.NewRouter(api), api
}
