package router

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type Operation huma.Operation

func Post[I, O struct{}](r *Router, path string, handler func(context.Context, *I) (*O, error)) {
	huma.Post(r.doc, path, handler)
}

func Get[I, O struct{}](r *Router, path string, handler func(context.Context, *I) (*O, error)) {
	huma.Get(r.doc, path, handler)
}

func Put[I, O struct{}](r *Router, path string, handler func(context.Context, *I) (*O, error)) {
	huma.Put(r.doc, path, handler)
}

func Patch[I, O struct{}](r *Router, path string, handler func(context.Context, *I) (*O, error)) {
	huma.Patch(r.doc, path, handler)
}

func Delete[I, O struct{}](r *Router, path string, handler func(context.Context, *I) (*O, error)) {
	huma.Delete(r.doc, path, handler)
}

func Register[I any, O any](r *Router, op Operation, handler func(context.Context, *I) (*O, error)) {
	huma.Register(r.doc, huma.Operation(op), handler)
}
