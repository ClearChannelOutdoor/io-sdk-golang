package api

import (
	"net/http"
	"sync"

	"golang.org/x/oauth2"
)

const (
	defaultEnvironmentProtocol string = "https"
)

var (
	Version string
)

type api struct {
	Mu         *sync.Mutex
	Clnt       *http.Client
	Svc        *Service
	OAuthToken *oauth2.Token
}
