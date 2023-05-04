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
	Name   string
	Proto  string
}

// IsValid returns true if the environment name, protocol, and token are
// not blank.
func (e *Service) IsValid() bool {
	return e.Name != "" && e.oauth2 != nil && e.Proto != ""
}

// SetServer sets the host and protocol of the environment based on the
// given server string. If the server string is invalid, the host and
// protocol are not changed.
func (e *Service) SetServer(s string) *Service {
	if u, err := url.Parse(s); err == nil {
		e.Host = u.Host
		e.Proto = u.Scheme
	}

	return e
}

// NewService returns a new Environment with the given name and token.
// The protocol is set to the default protocol (https).
func NewService(name string, oauth2 *clientcredentials.Config) *Service {
	return &Service{
		oauth2: oauth2,
		Name:   name,
		Proto:  defaultEnvironmentProtocol,
	}
}
