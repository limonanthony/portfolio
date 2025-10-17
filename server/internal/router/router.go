package router

import (
	"net/http"
	"slices"
)

type Router struct {
	*http.ServeMux

	globalChaining []func(http.Handler) http.Handler
	routeChaining  []func(http.Handler) http.Handler

	subRouter bool
}

func NewRouter() *Router {
	return &Router{ServeMux: http.NewServeMux(), subRouter: false}
}

func (r *Router) Use(middleware ...func(http.Handler) http.Handler) {
	if r.subRouter {
		r.routeChaining = append(r.routeChaining, middleware...)
	} else {
		r.globalChaining = append(r.globalChaining, middleware...)
	}
}

func (r *Router) Group(fn func(r *Router)) {
	subRouter := &Router{routeChaining: slices.Clone(r.routeChaining), subRouter: true, ServeMux: r.ServeMux}
	fn(subRouter)
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	for _, middleware := range slices.Backward(r.routeChaining) {
		handler = middleware(handler)
	}

	r.ServeMux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.Handle(pattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.ServeMux

	for _, middleware := range slices.Backward(r.globalChaining) {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, req)
}
