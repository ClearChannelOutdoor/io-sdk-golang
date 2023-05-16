package api

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	defaultEnvironmentProtocol string = "https"
)

type api struct {
	Clnt       *http.Client
	Svc        *Service
	OAuthToken *oauth2.Token
}

func (a *api) GetContext() context.Context {
	return context.Background()
}
