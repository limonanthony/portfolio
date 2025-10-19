package router

import "github.com/danielgtaylor/huma/v2"

func addTagOperation(tag string) func(o *huma.Operation, next func(operation *huma.Operation)) {
	return func(o *huma.Operation, next func(operation *huma.Operation)) {
		o.Tags = append(o.Tags, tag)
		next(o)
	}
}
