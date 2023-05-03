package api

import (
	query "go.jtlabs.io/query"
)

type SearchResult[T any] struct {
	Data    []*T          `json:"data"`
	Options query.Options `json:"options"`
	Total   uint          `json:"total"`
}
