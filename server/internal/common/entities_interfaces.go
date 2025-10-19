package common

import "context"

type Creator[Dto any] interface {
	Create(context.Context, Dto) (Id, error)
}

type GetterById[Entity any] interface {
	GetById(context.Context, Id) (Entity, error)
}

type AllGetter[Entity any] interface {
	GetAll(context.Context) ([]Entity, error)
}

type Updater[Dto any] interface {
	Update(context.Context, Id, Dto) error
}

type DeleterById interface {
	DeleteById(context.Context, Id) error
}
