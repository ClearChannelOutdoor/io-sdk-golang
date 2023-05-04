package client

import (
	"errors"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/api/display"
	"golang.org/x/oauth2/clientcredentials"
)

func NewDisplayClient(env Environment, oauth2 *clientcredentials.Config, svr ...string) (*client[display.Display], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	override := len(svr) > 0 && svr[0] != ""

	var svc *api.Service

	switch env {
	case DevelopEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayDevelopAPI)
	case ProductionEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayProductionAPI)
	case StagingEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayStageAPI)
	default:
		if !override {
			return nil, errors.New("custom environment requires server URL to be specified")
		}
	}

	// apply any server overrides, even if these are intended to override
	// a preset environment (e.g. DevelopEnvironment)
	if override {
		svc = api.NewService(env.String(), oauth2).SetServer(svr[0])
	}

	// ensure there is a valid server to connect to
	if svc == nil || svc.Proto == "" || svc.Host == "" {
		return nil, errors.New("the API server URL is invalid")
	}

	// create new endpoint
	ep := api.NewEndpoint[display.Display](svc, "/v1/displays")
	clnt := &client[display.Display]{
		dr:  ep,
		svc: svc,
	}

	// determine if oauth supports write operations
	if containsValue(oauth2.Scopes, ScopeDisplayModify) {
		clnt.dw = ep
	}

	return clnt, nil
}
