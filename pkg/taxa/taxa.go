package taxa

import (
	"context"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	serverUrl string = "https://direct.cco.io"
)

func NewCCOCodeClient(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[CCOCode], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[CCOCode](ctx, svr, "/v1/codes", oauth2)
}

func NewIABV1Client(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[V1Taxonomy], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[V1Taxonomy](ctx, svr, "/v1/taxa", oauth2)
}

func NewIABV2Client(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[IABTaxonomy], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[IABTaxonomy](ctx, svr, "/v2/taxa", oauth2)
}

func NewIABV3Client(ctx context.Context, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[IABTaxonomy], error) {
	svr := serverUrl
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[IABTaxonomy](ctx, svr, "/v3/taxa", oauth2)
}
