package api

type ReadChildResource[T any] interface {
	Search(string, ...*Options) (SearchResult[T], error)
	Get(string, string) (*T, error)
}

type ReadResource[T any] interface {
	Search(...*Options) (SearchResult[T], error)
	Get(string) (*T, error)
}

type WriteChildResouce[T any] interface {
	Create(string, *T) error
	Delete(string, string) error
	Patch(string, string, *T) error
	Update(string, string, *T) error
}

type WriteResource[T any] interface {
	Create(*T) error
	Delete(string) error
	Patch(string, *T) error
	Update(string, *T) error
}
