package clients

import (
	"errors"

	"cco.dev/io/pkg/api"
)

type Client[T any] struct {
	rep api.ReadAPI[T]
	wep api.WriteAPI[T]
	svc *api.Service
}

func (c *Client[T]) checkWrite() error {
	if c.wep == nil {
		return errors.New("write operations are not allowed with this client configuration")
	}

	return nil
}

func (c *Client[T]) Create(d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Create(d)
}

func (c *Client[T]) Delete(id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Delete(id)
}

func (c *Client[T]) Get(id string) (*T, error) {
	return c.rep.Get(id)
}

func (c *Client[T]) Patch(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Patch(id, d)
}

func (c *Client[T]) Search(opts ...api.Options) (api.SearchResult[T], error) {
	return c.rep.Search(opts...)
}

func (c *Client[T]) Update(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Update(id, d)
}

func NewClient[T any](svc *api.Service, readEP *api.Endpoint[T], writeEP *api.Endpoint[T]) *Client[T] {
	return &Client[T]{
		rep: readEP,
		svc: svc,
		wep: writeEP,
	}
}
