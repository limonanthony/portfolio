package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/limonanthony/portfolio/internal/database"
	"github.com/limonanthony/portfolio/internal/router"
)

func NewRouter(t *testing.T) (*router.Router, humatest.TestAPI) {
	_, api := humatest.New(t)

	return router.NewRouter(api), api
}

type routerTestAPI struct {
	humatest.TestAPI
	router *router.Router
}

func (r *routerTestAPI) Do(method, path string, body ...any) *httptest.ResponseRecorder {
	var bodyReader io.Reader
	if len(body) > 0 {
		jsonBytes, _ := json.Marshal(body[0])
		bodyReader = bytes.NewReader(jsonBytes)
	}

	req := httptest.NewRequest(method, path, bodyReader)
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	r.router.ServeHTTP(recorder, req)

	return recorder
}

func (r *routerTestAPI) DoCtx(ctx context.Context, method, path string, body ...any) *httptest.ResponseRecorder {
	var bodyReader io.Reader
	if len(body) > 0 {
		jsonBytes, _ := json.Marshal(body[0])
		bodyReader = bytes.NewReader(jsonBytes)
	}

	req := httptest.NewRequest(method, path, bodyReader)
	req = req.WithContext(ctx)
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	r.router.ServeHTTP(recorder, req)

	return recorder
}

func (r *routerTestAPI) Get(path string, body ...any) *httptest.ResponseRecorder {
	return r.Do(http.MethodGet, path, body...)
}

func (r *routerTestAPI) Post(path string, body ...any) *httptest.ResponseRecorder {
	return r.Do(http.MethodPost, path, body...)
}

func (r *routerTestAPI) Put(path string, body ...any) *httptest.ResponseRecorder {
	return r.Do(http.MethodPut, path, body...)
}

func (r *routerTestAPI) Patch(path string, body ...any) *httptest.ResponseRecorder {
	return r.Do(http.MethodPatch, path, body...)
}

func (r *routerTestAPI) Delete(path string, body ...any) *httptest.ResponseRecorder {
	return r.Do(http.MethodDelete, path, body...)
}

func NewRouterWithDb(t *testing.T) (*router.Router, humatest.TestAPI) {
	db, err := NewDatabase(t)
	if err != nil {
		t.Fatal(err)
	}

	router_ := router.NewRouter(nil)
	router_.Use(database.TransactionMiddleware(db))

	_, baseAPI := humatest.New(t)

	testAPI := &routerTestAPI{
		TestAPI: baseAPI,
		router:  router_,
	}

	return router_, testAPI
}
