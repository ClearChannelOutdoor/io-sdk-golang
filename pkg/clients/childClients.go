package clients

import (
	"context"
	"errors"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"golang.org/x/oauth2/clientcredentials"
)

type ChildClient[T any] struct {
	ctx      context.Context
	rep      api.ReadChildResource[T]
	wep      api.WriteChildResouce[T]
	writable bool
}

func (c *ChildClient[T]) checkWrite() error {
	if !c.writable {
		return errors.New("write operations are not allowed with this client configuration")
	}

	return nil
}

func (c *ChildClient[T]) Create(parentID string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Create(c.ctx, parentID, d)
}

func (c *ChildClient[T]) Delete(parentID string, id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Delete(c.ctx, parentID, id)
}

func (c *ChildClient[T]) Get(parentID string, id string) (*T, error) {
	return c.rep.Get(c.ctx, parentID, id)
}

func (c *ChildClient[T]) Patch(parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Patch(c.ctx, parentID, id, d)
}

func (c *ChildClient[T]) Search(parentID string, opts ...*api.Options) (api.SearchResult[T], error) {
	return c.rep.Search(c.ctx, parentID, opts...)
}

func (c *ChildClient[T]) Update(parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Update(c.ctx, parentID, id, d)
}

func NewChildClient[T any](ctx context.Context, env api.Environment, svr string, parentResource, childResouce string, oauth2 *clientcredentials.Config, writeScopes ...string) (*ChildClient[T], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	// define the host and protocol details for the API
	svc := api.NewService(env.String(), oauth2).SetServer(svr)

	// ensure there is a valid server to connect to
	if svc == nil || svc.Proto == "" || svc.Host == "" {
		return nil, errors.New("the API server URL is invalid")
	}

	ep := api.NewChildEndpoint[T](svc, parentResource, childResouce)

	// determine if oauth supports write operations
	for _, scope := range writeScopes {
		if internal.ContainsValue(oauth2.Scopes, scope) {
			return &ChildClient[T]{
				ctx:      ctx,
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
