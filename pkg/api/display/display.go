package display

import (
	"fmt"

	"cco.dev/io/pkg/api"
)

const (
	displayAPIHostFmt string = "display-api.%s.cco.dev"
)

type display struct {
	e *api.Endpoint
}

func NewDisplay(env api.Environment, hst ...string) *display {
	var h string
	if len(hst) > 0 {
		h = hst[0]
	}

	if h == "" {
		h = fmt.Sprintf(displayAPIHostFmt, env)
	}

	e := api.NewEndpoint(env.Token, env.Host, h, env.Proto)
	return &display{
		e: e,
	}
}
