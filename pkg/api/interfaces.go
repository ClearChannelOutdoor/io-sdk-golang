package api

type ReadAPI[T any] interface {
	GetAll(...Options) (SearchResult[T], error)
	GetOne(string) (T, error)
}

type WriteAPI[T any] interface {
	Create(T) (T, error)
	Delete(string) error
	Patch(string, T) (T, error)
	Post(string, T) (T, error)
}
