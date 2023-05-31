package api

import "context"

type ReadChildResource[T any] interface {
	Search(context.Context, string, ...*Options) (SearchResult[T], error)
	Get(context.Context, string, string) (*T, error)
}

type ReadResource[T any] interface {
	Search(context.Context, ...*Options) (SearchResult[T], error)
	Get(context.Context, string) (*T, error)
}

type WriteChildResouce[T any] interface {
	Create(context.Context, string, *T) error
	Delete(context.Context, string, string) error
	Patch(context.Context, string, string, *T) error
	Update(context.Context, string, string, *T) error
}

type WriteResource[T any] interface {
	Create(context.Context, *T) error
	Delete(context.Context, string) error
	Patch(context.Context, string, *T) error
	Update(context.Context, string, *T) error
}
