package adyen

import (
	"fmt"
	"strings"
)

// Environment allows clients to be configured for Testing
// and Production environments.
type Environment struct {
	Name      string
	apiURL    string
	clientURL string
	hppURL    string
}

const (
	// EnvironmentTesting identifies the Adyen testing environment.
	EnvironmentTesting = "test"

	// EnvironmentProduction identifies the Adyen live environment.
	EnvironmentProduction = "prod"
)

// Testing - instance of testing environment
var Testing = Environment{
	Name:      EnvironmentTesting,
	apiURL:    "https://pal-test.adyen.com/pal/servlet",
	clientURL: "https://test.adyen.com/hpp/cse/js/",
	hppURL:    "https://test.adyen.com/hpp/",
}

// Production - instance of production environment
var Production = Environment{
	Name:      EnvironmentProduction,
	apiURL:    "https://pal-live.adyen.com/pal/servlet",
	clientURL: "https://live.adyen.com/hpp/cse/js/",
	hppURL:    "https://live.adyen.com/hpp/",
}

type errParseEnvironment struct {
	name string
}

func (err errParseEnvironment) Error() string {
	return fmt.Sprintf("%q is not a valid environment name", err.name)
}

// ParseEnvironment returns an Adyen environment from a given name.
func ParseEnvironment(name string) (e Environment, err error) {
	switch strings.ToLower(name) {
	case EnvironmentTesting:
		e = Testing
	case EnvironmentProduction:
		e = Production
	default:
		err = errParseEnvironment{name: name}
	}

	return
}

// BaseURL returns api base url
func (e Environment) BaseURL(service string, version string) string {
	return e.apiURL + "/" + service + "/" + version
}

// ClientURL returns Adyen Client URL to load external scripts
func (e Environment) ClientURL(clientID string) string {
	return e.clientURL + clientID + ".shtml"
}

// HppURL returns Adyen HPP url to execute Hosted Payment Paged API requests
func (e Environment) HppURL(request string) string {
	return e.hppURL + request + ".shtml"
}
