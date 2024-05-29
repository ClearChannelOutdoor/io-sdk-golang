package creatives

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeCreativesModify string = "creatives-modify"
	serverUrl            string = "https://direct.cco.io"
)

func NewCreativeV1Client(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Creative], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Creative](ctx, svr, "/v1/creatives", oauth2, scopeCreativesModify)
}

func NewCreativeV2Client(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[AdCreative], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[AdCreative](ctx, svr, "/v2/creatives", oauth2, scopeCreativesModify)
}
