package displays

import (
	"context"
	"fmt"

	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeDisplaysModify string = "displays-modify"
	serverFmt           string = "https://direct%s.cco.io"
)

func NewClient(ctx context.Context, env api.Environment, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Display], error) {
	svr := fmt.Sprintf(serverFmt, "")
	if env != api.ProductionEnvironment {
		svr = fmt.Sprintf(serverFmt, fmt.Sprintf(".%s", env.String()))
	}

	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Display](ctx, env, svr, "/v1/displays", oauth2, scopeDisplaysModify)
}
