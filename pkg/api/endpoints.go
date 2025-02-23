package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

const (
	authorizationHeader     string = "Authorization"
	bearerFmt               string = "Bearer %s"
	defaultThrottleRetrySec uint   = 5
	deleteMethod            string = "DELETE"
	getMethod               string = "GET"
	patchMethod             string = "PATCH"
	postMethod              string = "POST"
	putMethod               string = "PUT"
)

var retryAfterRE *regexp.Regexp = regexp.MustCompile(`retry after (\d+)s\: `)

type Endpoint[T any] struct {
	a    *api
	hdr  *http.Header
	Path string
}

func NewEndpoint[T any](svc *Service, path string, hdr *http.Header) *Endpoint[T] {
	a := api{
		Clnt: &http.Client{},
		Svc:  svc,
		Mu:   &sync.Mutex{},
	}

	return &Endpoint[T]{
		a:    &a,
		hdr:  hdr,
		Path: path,
	}
}

func (e *Endpoint[T]) request(ctx context.Context, method string, path string, body io.Reader, opts ...*Options) ([]byte, error) {
	r, err := retryRequest(
		ctx,
		e.hdr,
		e.a,
		method,
		path,
		body,
		opts...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Create creates a new model of type T within the API
func (e *Endpoint[T]) Create(ctx context.Context, mdl *T) error {
	var created T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(ctx, postMethod, e.Path, bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &created); err != nil {
		return err
	}

	// assign the created model
	*mdl = created

	return nil
}

// Delete deletes a model of type T from the API
func (e *Endpoint[T]) Delete(ctx context.Context, id string) error {
	// make the request to the API
	_, err := e.request(ctx, deleteMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return err
	}

	return nil
}

// Get gets a single model of type T from the API
func (e *Endpoint[T]) Get(ctx context.Context, id string) (*T, error) {
	var mdl T

	// make the request to the API
	data, err := e.request(ctx, getMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return nil, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &mdl); err != nil {
		return nil, err
	}

	return &mdl, nil
}

// Patch updates a model of type T within the API while preserving
// any fields that are omitted and any additional externalIDs
// that already exist
func (e *Endpoint[T]) Patch(ctx context.Context, id string, mdl *T) error {
	var patched T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(ctx, patchMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &patched); err != nil {
		return err
	}

	// assign the patched model
	*mdl = patched

	return nil
}

// Search searches for models of type T within the API
func (e *Endpoint[T]) Search(ctx context.Context, opts ...*Options) (SearchResult[T], error) {
	var sr SearchResult[T]

	// make the request to the API
	data, err := e.request(ctx, getMethod, e.Path, nil, opts...)
	if err != nil {
		return sr, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &sr); err != nil {
		return sr, err
	}

	return sr, nil
}

// Update updates a model of type T within the API
func (e *Endpoint[T]) Update(ctx context.Context, id string, mdl *T) error {
	var updated T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(ctx, putMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &updated); err != nil {
		return err
	}

	// assign the updated model
	*mdl = updated

	return nil
}
