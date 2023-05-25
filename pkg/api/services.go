package api

import (
	"net/url"

	"golang.org/x/oauth2/clientcredentials"
)

// Service represents the environment in which the API is being used.
// It contains the name of the environment, the protocol, and a valid
// token for making API requests within the environment.
type Service struct {
	oauth2 *clientcredentials.Config
	Host   string
	Proto  string
}

// IsValid returns true if the environment name, protocol, and token are
// not blank.
func (e *Service) IsValid() bool {
	return e.Host != "" && e.oauth2 != nil && e.Proto != ""
}

// NewService returns a new Environment with the given name and token.
// The protocol is set to the default protocol (https).
func NewService(svr string, oauth2 *clientcredentials.Config) *Service {
	host := ""
	proto := defaultEnvironmentProtocol
	if u, err := url.Parse(svr); err == nil {
		host = u.Host
		proto = u.Scheme
	}

	return &Service{
		oauth2: oauth2,
		Host:   host,
		Proto:  proto,
	}
}
