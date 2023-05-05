package displays

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/clients"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scopeDisplaysModify string = "displays-modify"
	serverFmt           string = "https://display-api%s.cco.dev"
)

func NewClient(env api.Environment, oauth2 *clientcredentials.Config, overrideSvr ...string) (*clients.Client[Display], error) {
	svr := fmt.Sprintf(serverFmt, fmt.Sprintf(".%s", env.String()))
	if len(overrideSvr) > 0 && overrideSvr[0] != "" {
		svr = overrideSvr[0]
	}

	return clients.NewClient[Display](env, svr, "/v1/displays", oauth2, scopeDisplaysModify)
}
