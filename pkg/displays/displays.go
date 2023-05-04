package displays

import (
	"errors"

	"cco.dev/io/internal"
	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/client"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	DisplayDevelopAPI    string = "https://display-api.dev.cco.dev"
	DisplayProductionAPI string = "https://display-api.cco.dev"
	DisplayStageAPI      string = "https://display-api.stg.cco.dev"
	ScopeDisplayModify   string = "displays-modify"
)

func NewClient(env api.Environment, oauth2 *clientcredentials.Config, svr ...string) (*client.Client[Display], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	override := len(svr) > 0 && svr[0] != ""

	var svc *api.Service

	switch env {
	case api.DevelopEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayDevelopAPI)
	case api.ProductionEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(DisplayProductionAPI)
	case api.StagingEnvironment:
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
	ep := api.NewEndpoint[Display](svc, "/v1/displays")

	// determine if oauth supports write operations
	if internal.ContainsValue(oauth2.Scopes, ScopeDisplayModify) {
		// return a read-write client
		return client.NewClient(svc, ep, ep), nil
	}

	// return a read-only client
	return client.NewClient(svc, ep, nil), nil
}
