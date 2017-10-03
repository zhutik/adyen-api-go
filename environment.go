package adyen

// environment structure to support testing and production
type environment struct {
	apiURL    string
	clientURL string
	hppURL    string
}

// newEnvironment create environment based on api url
func newEnvironment(apiURL, clientURL, hppURL string) environment {
	return environment{apiURL: apiURL, clientURL: clientURL, hppURL: hppURL}
}

// BaseURL returns api base url
func (e environment) BaseURL(service string, version string) string {
	return e.apiURL + "/" + service + "/" + version
}

// ClientURL returns Adyen Client URL to load external scripts
func (e environment) ClientURL(clientID string) string {
	return e.clientURL + clientID + ".shtml"
}

// HppURL returns Adyen HPP url to execute Hosted Payment Paged API requests
func (e environment) HppURL(request string) string {
	return e.hppURL + request + ".shtml"
}

// Testing - instance of testing environment
var Testing = newEnvironment(
	"https://pal-test.adyen.com/pal/servlet",
	"https://test.adyen.com/hpp/cse/js/",
	"https://test.adyen.com/hpp/",
)

// Production - instance of production environment
var Production = newEnvironment(
	"https://pal-live.adyen.com/pal/servlet",
	"https://live.adyen.com/hpp/cse/js/",
	"https://live.adyen.com/hpp/",
)
