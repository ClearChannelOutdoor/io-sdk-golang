package networks

import (
	"errors"

	"cco.dev/io/internal"
	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	NetworkDevelopAPI    string = "https://network-api.dev.cco.dev"
	NetworkProductionAPI string = "https://network-api.cco.dev"
	NetworkStageAPI      string = "https://network-api.stg.cco.dev"
	ScopeNetworkModify   string = "networks-modify"
)

func NewClient(env api.Environment, oauth2 *clientcredentials.Config, svr ...string) (*clients.Client[Network], error) {
	if oauth2 == nil {
		return nil, errors.New("oauth2 configuration is required")
	}

	override := len(svr) > 0 && svr[0] != ""

	var svc *api.Service

	switch env {
	case api.DevelopEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(NetworkDevelopAPI)
	case api.ProductionEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(NetworkProductionAPI)
	case api.StagingEnvironment:
		svc = api.NewService(env.String(), oauth2).SetServer(NetworkStageAPI)
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
	ep := api.NewEndpoint[Network](svc, "/v1/networks")

	// determine if oauth supports write operations
	if internal.ContainsValue(oauth2.Scopes, ScopeNetworkModify) {
		// return a read-write client
		return clients.NewClient(svc, ep, ep), nil
	}

	// return a read-only client
	return clients.NewClient(svc, ep, nil), nil
}
