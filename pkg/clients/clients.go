package clients

import (
	"errors"

	"cco.dev/io/internal"
	"cco.dev/io/pkg/api"
	"golang.org/x/oauth2/clientcredentials"
)

type Client[T any] struct {
	rep      api.ReadAPI[T]
	wep      api.WriteAPI[T]
	svc      *api.Service
	writable bool
}

func (c *Client[T]) checkWrite() error {
	if !c.writable {
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

func (c *Client[T]) Search(opts ...*api.Options) (api.SearchResult[T], error) {
	return c.rep.Search(opts...)
}

func (c *Client[T]) Update(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Update(id, d)
}

func NewClient[T any](env api.Environment, svr string, resource string, oauth2 *clientcredentials.Config, writeScopes ...string) (*Client[T], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	// define the host and protocol details for the API
	svc := api.NewService(env.String(), oauth2).SetServer(svr)

	// ensure there is a valid server to connect to
	if svc == nil || svc.Proto == "" || svc.Host == "" {
		return nil, errors.New("the API server URL is invalid")
	}

	// create new endpoint
	ep := api.NewEndpoint[T](svc, resource)

	// determine if oauth supports write operations
	for _, scope := range writeScopes {
		if internal.ContainsValue(oauth2.Scopes, scope) {
			return &Client[T]{
				rep:      ep,
				svc:      svc,
				wep:      ep,
				writable: true,
			}, nil
		}
	}

	return &Client[T]{
		rep: ep,
		svc: svc,
	}, nil
}
