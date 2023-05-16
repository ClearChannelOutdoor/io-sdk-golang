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
