package networks

import (
	"fmt"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeNetworksModify string = "networks-modify"
	serverFmt           string = "https://direct%s.cco.io"
)

func NewClient(env api.Environment, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Network], error) {
	svr := fmt.Sprintf(serverFmt, "")
	if env != api.ProductionEnvironment {
		svr = fmt.Sprintf(serverFmt, fmt.Sprintf(".%s", env.String()))
	}

	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Network](env, svr, "/v1/networks", oauth2, scopeNetworksModify)
}
