package interfaces

import "cco.dev/io/pkg/api"

type ReadAPI[T any] interface {
	GetAll(api.Options) (T, error)
	GetOne(string) (T, error)
}

type WriteAPI[T any] interface {
	Create(T) (T, error)
	Delete(string) error
	Patch(T) (T, error)
	Update(T) (T, error)
}
