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

// BaseUrl returns api base url
func (e Environment) BaseUrl(version string) string {
	return e.apiURL + "/" + version
}

// ClientUrl returns Adyen Client URL to load external scripts
func (e Environment) ClientUrl(clientID string) string {
	return e.clientURL + clientID + ".shtml"
}

var (
	Testing = NewEnvironment(
		"https://pal-test.adyen.com/pal/servlet/Payment",
		"https://test.adyen.com/hpp/cse/js/",
	)
	Production = NewEnvironment(
		"https://pal-live.adyen.com/pal/servlet/Payment",
		"https://live.adyen.com/hpp/cse/js/",
	)
)
