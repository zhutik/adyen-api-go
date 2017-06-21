package adyen

// Environment structure to support testing and production
type Environment struct {
	apiURL    string
	clientURL string
}

// NewEnvironment create environment based on api url
func NewEnvironment(apiURL, clientURL string) Environment {
	return Environment{apiURL: apiURL, clientURL: clientURL}
}

// BaseURL returns api base url
func (e Environment) BaseURL(version string) string {
	return e.apiURL + "/" + version
}

// ClientURL returns Adyen Client URL to load external scripts
func (e Environment) ClientURL(clientID string) string {
	return e.clientURL + clientID + ".shtml"
}

// Testing - instance of testing environment
var Testing = NewEnvironment(
	"https://pal-test.adyen.com/pal/servlet/Payment",
	"https://test.adyen.com/hpp/cse/js/",
)

// Production - instance of production environment
var Production = NewEnvironment(
	"https://pal-live.adyen.com/pal/servlet/Payment",
	"https://live.adyen.com/hpp/cse/js/",
)
