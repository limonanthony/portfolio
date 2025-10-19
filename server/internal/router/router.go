package router

import (
	"net/http"
	"slices"

	"github.com/danielgtaylor/huma/v2"
)

type Router struct {
	Mux *http.ServeMux

	path string
	doc  huma.API

	globalChaining []func(http.Handler) http.Handler
	routeChaining  []func(http.Handler) http.Handler

	subRouter bool
}

func NewRouter(api huma.API) *Router {
	mux := http.NewServeMux()
	doc := newDocApi(mux, api)
	router := Router{Mux: mux, subRouter: false, path: "", doc: doc}
	registerScalaDocumentation(router)
	return &router
}

func (r *Router) Use(middleware ...func(http.Handler) http.Handler) {
	if r.subRouter {
		r.routeChaining = append(r.routeChaining, middleware...)
	} else {
		r.globalChaining = append(r.globalChaining, middleware...)
	}
}

func (r *Router) Group(path string, name string, fn func(r *Router)) {
	doc := huma.NewGroup(r.doc, path)
	doc.UseModifier(addTagOperation(name))
	subRouter := &Router{routeChaining: slices.Clone(r.routeChaining), doc: doc, subRouter: true, Mux: r.Mux, path: r.path + path}
	fn(subRouter)
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	for _, middleware := range slices.Backward(r.routeChaining) {
		handler = middleware(handler)
	}

	r.Mux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.Mux

	for _, middleware := range slices.Backward(r.globalChaining) {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, req)
}
