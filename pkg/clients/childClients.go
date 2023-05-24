package clients

import (
	"context"
	"errors"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"golang.org/x/oauth2/clientcredentials"
)

type ChildClient[T any] struct {
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

func (c *ChildClient[T]) Create(ctx context.Context, parentID string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Create(ctx, parentID, d)
}

func (c *ChildClient[T]) Delete(ctx context.Context, parentID string, id string) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Delete(ctx, parentID, id)
}

func (c *ChildClient[T]) Get(ctx context.Context, parentID string, id string) (*T, error) {
	return c.rep.Get(ctx, parentID, id)
}

func (c *ChildClient[T]) Patch(ctx context.Context, parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Patch(ctx, parentID, id, d)
}

func (c *ChildClient[T]) Search(ctx context.Context, parentID string, opts ...*api.Options) (api.SearchResult[T], error) {
	return c.rep.Search(ctx, parentID, opts...)
}

func (c *ChildClient[T]) Update(ctx context.Context, parentID string, id string, d *T) error {
	if err := c.checkWrite(); err != nil {
		return err
	}

	return c.wep.Update(ctx, parentID, id, d)
}

func NewChildClient[T any](env api.Environment, svr string, parentResource, childResouce string, oauth2 *clientcredentials.Config, writeScopes ...string) (*ChildClient[T], error) {
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
				rep:      ep,
				wep:      ep,
				writable: true,
			}, nil
		}
	}

	return &ChildClient[T]{
		rep: ep,
	}, nil
}
