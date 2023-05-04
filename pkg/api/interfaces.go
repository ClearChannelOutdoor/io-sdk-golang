package api

type ReadAPI[T any] interface {
	Search(...Options) (SearchResult[T], error)
	Get(string) (*T, error)
}

type WriteAPI[T any] interface {
	Create(*T) error
	Delete(string) error
	Patch(string, *T) error
	Update(string, *T) error
}
