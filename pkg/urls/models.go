package urls

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeLocationsModify string = "locations-modify"
	serverUrl            string = "https://direct.cco.io"
)

func NewClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Location], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Location](ctx, svr, "/v1/locations", oauth2, scopeLocationsModify)
}
