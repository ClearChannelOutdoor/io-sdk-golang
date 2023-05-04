package client

import (
	"errors"

	"cco.dev/io/pkg/api"
)

const (
	DisplayDevelopAPI    string = "https://display-api.dev.cco.dev"
	DisplayProductionAPI string = "https://display-api.cco.dev"
	DisplayStageAPI      string = "https://display-api.stg.cco.dev"
	ScopeDisplayModify   string = "displays-modify"
)

func containsValue(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

type client[T any] struct {
	dr  api.ReadAPI[T]
	dw  api.WriteAPI[T]
	svc *api.Service
}

func (c *client[T]) checkWrite() error {
	if c.dw == nil {
		return errors.New("write operations are not allowed with this client configuration")
	}

	return nil
}

func (c *client[T]) Create(d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.dw.Create(d)
}

func (c *client[T]) Delete(id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.dw.Delete(id)
}

func (c *client[T]) Get(id string) (*T, error) {
	return c.dr.Get(id)
}

func (c *client[T]) Patch(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.dw.Patch(id, d)
}

func (c *client[T]) Search(opts ...api.Options) (api.SearchResult[T], error) {
	return c.dr.Search(opts...)
}

func (c *client[T]) Update(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.dw.Update(id, d)
}
