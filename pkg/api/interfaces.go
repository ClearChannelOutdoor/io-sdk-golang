package api

import "context"

type ReadChildResource[T any] interface {
	Search(context.Context, string, ...*Options) (SearchResult[T], int, error)
	Get(context.Context, string, string) (*T, int, error)
}

type ReadResource[T any] interface {
	Search(context.Context, ...*Options) (SearchResult[T], int, error)
	Get(context.Context, string) (*T, int, error)
}

type WriteChildResouce[T any] interface {
	Create(context.Context, string, *T) (int, error)
	Delete(context.Context, string, string) (int, error)
	Patch(context.Context, string, string, *T) (int, error)
	Update(context.Context, string, string, *T) (int, error)
}

type WriteResource[T any] interface {
	Create(context.Context, *T) (int, error)
	Delete(context.Context, string) (int, error)
	Patch(context.Context, string, *T) (int, error)
	Update(context.Context, string, *T) (int, error)
}
