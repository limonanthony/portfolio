package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func newDocApi(mux *http.ServeMux, api huma.API) huma.API {
	config := huma.DefaultConfig("Portfolio API", "1.0.0")
	config.DocsPath = ""

	if api != nil {
		return api
	}

	return humago.New(mux, config)
}

func registerScalaDocumentation(router Router) {
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(`<!doctype html>
	<html>
	 <head>
	   <title>API Reference</title>
	   <meta charset="utf-8" />
	   <meta
	     name="viewport"
	     content="width=device-width, initial-scale=1" />
	 </head>
	 <body>
	   <script
	     id="api-reference"
	     data-url="/openapi.json"></script>
	   <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	 </body>
	</html>`))
		if err != nil {
			return
		}
	})
}
