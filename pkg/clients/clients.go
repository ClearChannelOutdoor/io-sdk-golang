package clients

import (
	"context"
	"errors"
	"net/http"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"golang.org/x/oauth2/clientcredentials"
)

type Client[T any] struct {
	ctx      context.Context
	Headers  *http.Header
	rep      api.ReadResource[T]
	scb      func(int)
	wep      api.WriteResource[T]
	writable bool
}

func (c *Client[T]) checkWrite() error {
	if !c.writable {
		return errors.New("write operations are not allowed with this client configuration")
	}

	return nil
}

func (c *Client[T]) ClearStatusCallback() {
	c.scb = nil
}

func (c *Client[T]) Create(d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Create(c.ctx, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *Client[T]) Delete(id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Delete(c.ctx, id)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *Client[T]) Get(id string) (*T, error) {
	v, sts, err := c.rep.Get(c.ctx, id)

	if c.scb != nil {
		c.scb(sts)
	}

	return v, err
}

func (c *Client[T]) Patch(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Patch(c.ctx, id, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *Client[T]) Search(opts ...*api.Options) (api.SearchResult[T], error) {
	sr, sts, err := c.rep.Search(c.ctx, opts...)

	if c.scb != nil {
		c.scb(sts)
	}

	return sr, err
}

func (c *Client[T]) SetStatusCallback(cb func(int)) {
	c.scb = cb
}

func (c *Client[T]) Update(id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Update(c.ctx, id, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func NewClient[T any](ctx context.Context, svr string, resource string, oauth2 *clientcredentials.Config, writeScopes ...string) (*Client[T], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	// define the host and protocol details for the API
	svc := api.NewService(svr, oauth2)

	// ensure there is a valid server to connect to
	if !svc.IsValid() {
		return nil, errors.New("the API server URL is invalid")
	}

	// create a header collection for the client
	hdr := make(http.Header)

	// create new endpoint
	ep := api.NewEndpoint[T](svc, resource, &hdr)

	// determine if oauth supports write operations
	for _, scope := range writeScopes {
		if internal.ContainsValue(oauth2.Scopes, scope) {
			return &Client[T]{
				ctx:      ctx,
				Headers:  &hdr,
				rep:      ep,
				wep:      ep,
				writable: true,
			}, nil
		}
	}

	return &Client[T]{
		ctx: ctx,
		rep: ep,
	}, nil
}
