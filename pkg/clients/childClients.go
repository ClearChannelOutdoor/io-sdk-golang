package clients

import (
	"context"
	"errors"
	"net/http"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"golang.org/x/oauth2/clientcredentials"
)

type ChildClient[T any] struct {
	ctx      context.Context
	Headers  *http.Header
	rep      api.ReadChildResource[T]
	scb      func(int)
	wep      api.WriteChildResouce[T]
	writable bool
}

func (c *ChildClient[T]) checkWrite() error {
	if !c.writable {
		return errors.New("write operations are not allowed with this client configuration")
	}

	return nil
}

func (c *ChildClient[T]) ClearStatusCallback() {
	c.scb = nil
}

func (c *ChildClient[T]) Create(parentID string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Create(c.ctx, parentID, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *ChildClient[T]) Delete(parentID string, id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Delete(c.ctx, parentID, id)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *ChildClient[T]) Get(parentID string, id string) (*T, error) {
	v, sts, err := c.rep.Get(c.ctx, parentID, id)

	if c.scb != nil {
		c.scb(sts)
	}

	return v, err
}

func (c *ChildClient[T]) Patch(parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Patch(c.ctx, parentID, id, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func (c *ChildClient[T]) Search(parentID string, opts ...*api.Options) (api.SearchResult[T], error) {
	sr, sts, err := c.rep.Search(c.ctx, parentID, opts...)

	if c.scb != nil {
		c.scb(sts)
	}

	return sr, err
}

func (c *ChildClient[T]) SetStatusCallback(cb func(int)) {
	c.scb = cb
}

func (c *ChildClient[T]) Update(parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	sts, err := c.wep.Update(c.ctx, parentID, id, d)

	if c.scb != nil {
		c.scb(sts)
	}

	return err
}

func NewChildClient[T any](ctx context.Context, svr string, parentResource, childResouce string, oauth2 *clientcredentials.Config, writeScopes ...string) (*ChildClient[T], error) {
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

	ep := api.NewChildEndpoint[T](svc, parentResource, childResouce, &hdr)

	// determine if oauth supports write operations
	for _, scope := range writeScopes {
		if internal.ContainsValue(oauth2.Scopes, scope) {
			return &ChildClient[T]{
				ctx:      ctx,
				Headers:  &hdr,
				rep:      ep,
				wep:      ep,
				writable: true,
			}, nil
		}
	}

	return &ChildClient[T]{
		ctx: ctx,
		rep: ep,
	}, nil
}
