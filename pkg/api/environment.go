package api

import "golang.org/x/oauth2/clientcredentials"

// Environment represents the environment in which the API is being used.
// It contains the name of the environment, the protocol, and a valid
// token for making API requests within the environment.
type Environment struct {
	oauth2 *clientcredentials.Config
	Host   string
	Name   string
	Proto  string
}

// IsValid returns true if the environment name, protocol, and token are
// not blank.
func (e *Environment) IsValid() bool {
	return e.Name != "" && e.oauth2 != nil && e.Proto != ""
}

// NewEnvironment returns a new Environment with the given name and token.
// The protocol is set to the default protocol (https).
func NewEnvironment(name string, oauth2 *clientcredentials.Config) *Environment {
	return &Environment{
		oauth2: oauth2,
		Name:   name,
		Proto:  defaultEnvironmentProtocol,
	}
}
