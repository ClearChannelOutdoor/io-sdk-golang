package client

import (
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

func (c *client[T]) GetAll(opts ...api.Options) (api.SearchResult[T], error) {
	return c.dr.GetAll(opts...)
}

func (c *client[T]) GetOne(id string) (T, error) {
	return c.dr.GetOne(id)
}

func (c *client[T]) Create(d T) (T, error) {
	return c.dw.Create(d)
}

func (c *client[T]) Delete(id string) error {
	return c.dw.Delete(id)
}

func (c *client[T]) Patch(id string, d T) (T, error) {
	return c.dw.Patch(id, d)
}

func (c *client[T]) Post(id string, d T) (T, error) {
	return c.dw.Post(id, d)
}
