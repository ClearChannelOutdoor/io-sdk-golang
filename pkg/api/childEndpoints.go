package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"sync"
)

type ChildEndpoint[T any] struct {
	a    *api
	hdr  *http.Header
	Path string
}

func NewChildEndpoint[T any](svc *Service, parentPath, childPath string, hdr *http.Header) *ChildEndpoint[T] {
	a := api{
		Clnt: &http.Client{},
		Svc:  svc,
		Mu:   &sync.Mutex{},
	}

	return &ChildEndpoint[T]{
		a:    &a,
		hdr:  hdr,
		Path: strings.Join([]string{parentPath, "/%s", childPath}, ""),
	}
}

func (ce *ChildEndpoint[T]) request(ctx context.Context, method string, path string, body io.Reader, opts ...*Options) ([]byte, int, error) {
	r, sts, err := retryRequest(
		ctx,
		ce.hdr,
		ce.a,
		method,
		path,
		body,
		opts...)
	if err != nil {
		return nil, sts, err
	}

	return r, sts, nil
}

func (ce *ChildEndpoint[T]) Create(ctx context.Context, parentID string, mdl *T) (int, error) {
	var created T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return 0, err
	}

	data, sts, err := ce.request(ctx, postMethod, fmt.Sprintf(ce.Path, parentID), bytes.NewBuffer(jd))
	if err != nil {
		return sts, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &created); err != nil {
		return sts, err
	}

	// assign the created model
	*mdl = created

	return sts, nil
}

func (ce *ChildEndpoint[T]) Delete(ctx context.Context, parentID string, id string) (int, error) {
	// make the request to the API
	_, sts, err := ce.request(ctx, deleteMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), nil)
	if err != nil {
		return sts, err
	}

	return sts, nil
}

func (ce *ChildEndpoint[T]) Get(ctx context.Context, parentID string, id string) (*T, int, error) {
	var mdl T

	// make the request to the API
	data, sts, err := ce.request(ctx, getMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), nil)
	if err != nil {
		return nil, sts, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &mdl); err != nil {
		return nil, sts, err
	}

	return &mdl, sts, nil
}

func (ce *ChildEndpoint[T]) Patch(ctx context.Context, parentID string, id string, mdl *T) (int, error) {
	var patched T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return 0, err
	}

	// make the request to the API
	data, sts, err := ce.request(ctx, patchMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), bytes.NewBuffer(jd))
	if err != nil {
		return sts, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &patched); err != nil {
		return sts, err
	}

	// assign the created model
	*mdl = patched

	return sts, nil
}

func (ce *ChildEndpoint[T]) Search(ctx context.Context, parentID string, opts ...*Options) (SearchResult[T], int, error) {
	var sr SearchResult[T]

	// make the request to the API
	data, sts, err := ce.request(ctx, getMethod, fmt.Sprintf(ce.Path, parentID), nil, opts...)
	if err != nil {
		return sr, sts, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &sr); err != nil {
		return sr, sts, err
	}

	return sr, sts, nil
}

func (ce *ChildEndpoint[T]) Update(ctx context.Context, parentID string, id string, mdl *T) (int, error) {
	var updated T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return 0, err
	}

	// make the request to the API
	data, sts, err := ce.request(ctx, putMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), bytes.NewBuffer(jd))
	if err != nil {
		return sts, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &updated); err != nil {
		return sts, err
	}

	// assign the updated model
	*mdl = updated

	return sts, nil
}
