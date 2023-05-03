package api

// Environment represents the environment in which the API is being used.
// It contains the name of the environment, the protocol, and a valid
// token for making API requests within the environment.
type Environment struct {
	Host  string
	Name  string
	Proto string
	Token string
}

// IsValid returns true if the environment name, protocol, and token are
// not blank.
func (e *Environment) IsValid() bool {
	return e.Name != "" && e.Token != "" && e.Proto != ""
}

// NewEnvironment returns a new Environment with the given name and token.
// The protocol is set to the default protocol (https).
func NewEnvironment(name, token string) *Environment {
	return &Environment{
		Name:  name,
		Proto: defaultEnvironmentProtocol,
		Token: token,
	}
}
