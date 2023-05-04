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

	var svc *api.Service

	switch env {
	case DevelopEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayDevelopAPI)
	case ProductionEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayProductionAPI)
	case StagingEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayStageAPI)
	default:
		if len(svr) == 0 || len(svr[0]) == 0 {
			return nil, errors.New("custom environment requires server URL to be specified")
		}

		svc = api.NewService(env.String(), oauth2).SetServer(svr[0])
		if svc.Proto == "" || svc.Host == "" {
			return nil, errors.New("custom environment server URL is invalid")
		}
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
