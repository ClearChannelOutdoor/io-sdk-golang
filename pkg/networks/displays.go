package networks

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

func NewDisplayClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.ChildClient[NetworkDisplay], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewChildClient[NetworkDisplay](ctx, svr, "/v1/networks", "/displays", oauth2, scopeNetworksModify)
}
