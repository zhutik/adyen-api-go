package adyen

// APICredentials basic API settings
//
// Description:
//
//     - env - Environment for next API calls
//     - Username - API username for authentication
//     - Password - API password for authentication
//     - merchantID - Merchant Account settings for payment and modification requests
//     - clientID - Used to load external JS files from Adyen, to encrypt client requests
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
type APICredentials struct {
	env         Environment
	Username    string
	Password    string
	merchantID  string
	clientID    string
	hppSettings HPPCredentials
}

// HPPCredentials used to communicate with Adyen HPP API
//
// Description:
//
//     - hmac - is generated when new Skin is created in Adyen Customer Area
//     - skinCode - skin code from Adyen CA
//     - shopperLocale - merchant local settings (in ISO format)
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type HPPCredentials struct {
	hmac          string
	skinCode      string
	shopperLocale string
}

// NewCredentials create new APICredentials
func NewCredentials(env Environment, username, password, merchantID, clientID string) APICredentials {
	return APICredentials{
		env:        env,
		Username:   username,
		Password:   password,
		merchantID: merchantID,
		clientID:   clientID,
	}
}

// NewCredentialsWithHPPSettings create new APICredentials and specify Adyen Hosted Payment Page settings
func NewCredentialsWithHPPSettings(env Environment, username, password, merchantID, clientID, hmac, skinCode, shopperLocale string) APICredentials {
	return APICredentials{
		env:        env,
		Username:   username,
		Password:   password,
		merchantID: merchantID,
		clientID:   clientID,
		hppSettings: HPPCredentials{
			hmac:          hmac,
			skinCode:      skinCode,
			shopperLocale: shopperLocale,
		},
	}
}
