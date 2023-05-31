package api

import (
	query "go.jtlabs.io/query"
)

// SearchResult is a generic struct for returning search results
// from any API. It contains the list of results, the query options
// used to generate the results, and the total number of results.
type SearchResult[T any] struct {
	Data    []*T          `json:"data"`
	Options query.Options `json:"options"`
	Total   uint          `json:"total"`
}
