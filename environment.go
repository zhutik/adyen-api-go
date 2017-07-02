package adyen

// Environment structure to support testing and production
type Environment struct {
	apiURL    string
	clientURL string
	hppURL    string
}

// NewEnvironment create environment based on api url
func NewEnvironment(apiURL, clientURL, hppURL string) Environment {
	return Environment{apiURL: apiURL, clientURL: clientURL, hppURL: hppURL}
}

// BaseURL returns api base url
func (e Environment) BaseURL(version string) string {
	return e.apiURL + "/" + version
}

// ClientURL returns Adyen Client URL to load external scripts
func (e Environment) ClientURL(clientID string) string {
	return e.clientURL + clientID + ".shtml"
}

// HppURL returns Adyen HPP url to execute Hosted Payment Paged API requests
func (e Environment) HppURL(request string) string {
	return e.hppURL + request + ".shtml"
}

// Testing - instance of testing environment
var Testing = NewEnvironment(
	"https://pal-test.adyen.com/pal/servlet/Payment",
	"https://test.adyen.com/hpp/cse/js/",
	"https://test.adyen.com/hpp/",
)

// Production - instance of production environment
var Production = NewEnvironment(
	"https://pal-live.adyen.com/pal/servlet/Payment",
	"https://live.adyen.com/hpp/cse/js/",
	"https://live.adyen.com/hpp/",
)
