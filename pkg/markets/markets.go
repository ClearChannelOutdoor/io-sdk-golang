package markets

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeMarketsModify string = "markets-modify"
	serverFmt          string = "https://direct%s.cco.io"
)

func NewClient(env api.Environment, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Market], error) {
	svr := fmt.Sprintf(serverFmt, "")
	if env != api.ProductionEnvironment {
		svr = fmt.Sprintf(serverFmt, fmt.Sprintf(".%s", env.String()))
	}

	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Market](env, svr, "/v1/markets", oauth2, scopeMarketsModify)
}
