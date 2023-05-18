package networks

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

func NewDisplayClient(env api.Environment, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.ChildClient[NetworkDisplay], error) {
	svr := fmt.Sprintf(serverFmt, "")
	if env != api.ProductionEnvironment {
		svr = fmt.Sprintf(serverFmt, fmt.Sprintf(".%s", env.String()))
	}

	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewChildClient[NetworkDisplay](env, svr, "/v1/networks", "/displays", oauth2, scopeNetworksModify)
}
