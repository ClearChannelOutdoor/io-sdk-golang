package interfaces

import "cco.dev/io/pkg/api"

type ReadAPI[T any] interface {
	GetAll(o api.Options) T
	GetOne(id string) T
}

type WriteAPI[T any] interface {
	Create(t T) T
	Delete(t T) T
	Overwrite(t T) T
	Patch(t T) T
}
