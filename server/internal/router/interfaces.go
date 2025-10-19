package router

import "context"

type Handler[I, O any] interface {
	func(context.Context, *I) (*O, error)
}
