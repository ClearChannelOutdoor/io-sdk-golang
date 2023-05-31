package api

import (
	"net/http"

	"golang.org/x/oauth2"
)

const (
	defaultEnvironmentProtocol string = "https"
)

var (
	Version string
)

type api struct {
	Clnt       *http.Client
	Svc        *Service
	OAuthToken *oauth2.Token
}
