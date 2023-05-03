package api

type ReadAPI[T any] interface {
	GetAll(Options) (T, error)
	GetOne(string) (T, error)
}

type WriteAPI[T any] interface {
	Create(T) (T, error)
	Delete(string) error
	Patch(T) (T, error)
	Post(T) (T, error)
}
