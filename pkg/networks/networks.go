package networks

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeNetworksModify string = "networks-modify"
	serverUrl           string = "https://direct.cco.io"
)

func NewClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Network], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Network](ctx, svr, "/v1/networks", oauth2, scopeNetworksModify)
}

func NewDisplayClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.ChildClient[NetworkDisplay], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewChildClient[NetworkDisplay](ctx, svr, "/v1/networks", "/displays", oauth2, scopeNetworksModify)
}
