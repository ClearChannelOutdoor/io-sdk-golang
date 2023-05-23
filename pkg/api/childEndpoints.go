package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
)

type ChildEndpoint[T any] struct {
	a    *api
	Path string
}

func NewChildEndpoint[T any](svc *Service, parentPath, childPath string) *ChildEndpoint[T] {
	a := api{
		Clnt: &http.Client{},
		Svc:  svc,
	}

	return &ChildEndpoint[T]{
		a:    &a,
		Path: strings.Join([]string{parentPath, "/%s", childPath}, ""),
	}
}

func (ce *ChildEndpoint[T]) request(method string, path string, body io.Reader, opts ...*Options) ([]byte, error) {
	r, err := retryRequest(
		ce.a,
		method,
		path,
		body,
		opts...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (ce *ChildEndpoint[T]) Create(parentID string, mdl *T) error {
	var created T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	data, err := ce.request(postMethod, fmt.Sprintf(ce.Path, parentID), bytes.NewBuffer(jd))
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

func (ce *ChildEndpoint[T]) Delete(parentID string, id string) error {
	// make the request to the API
	_, err := ce.request(deleteMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), nil)
	if err != nil {
		return err
	}

	return nil
}

func (ce *ChildEndpoint[T]) Get(parentID string, id string) (*T, error) {
	var mdl T

	// make the request to the API
	data, err := ce.request(getMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), nil)
	if err != nil {
		return nil, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &mdl); err != nil {
		return nil, err
	}

	return &mdl, nil
}

func (ce *ChildEndpoint[T]) Patch(parentID string, id string, mdl *T) error {
	var patched T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := ce.request(patchMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &patched); err != nil {
		return err
	}

	// assign the created model
	*mdl = patched

	return nil
}

func (ce *ChildEndpoint[T]) Search(parentID string, opts ...*Options) (SearchResult[T], error) {
	var sr SearchResult[T]

	// make the request to the API
	data, err := ce.request(getMethod, fmt.Sprintf(ce.Path, parentID), nil, opts...)
	if err != nil {
		return sr, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &sr); err != nil {
		return sr, err
	}

	return sr, nil
}

func (ce *ChildEndpoint[T]) Update(parentID string, id string, mdl *T) error {
	var updated T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := ce.request(putMethod, path.Join(fmt.Sprintf(ce.Path, parentID), id), bytes.NewBuffer(jd))
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
