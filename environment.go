package adyen

// Environment allows clients to be configured for Testing
// and Production environments.
type Environment struct {
	apiURL    string
	clientURL string
	hppURL    string
}

// Testing - instance of testing environment
var Testing = Environment{
	apiURL:    "https://pal-test.adyen.com/pal/servlet",
	clientURL: "https://test.adyen.com/hpp/cse/js/",
	hppURL:    "https://test.adyen.com/hpp/",
}

// Production - instance of production environment
var Production = Environment{
	apiURL:    "https://pal-live.adyen.com/pal/servlet",
	clientURL: "https://live.adyen.com/hpp/cse/js/",
	hppURL:    "https://live.adyen.com/hpp/",
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
