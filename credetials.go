package adyen

// apiCredentials basic API settings
//
// Description:
//
//     env - Environment for next API calls
//     Username - API username for authentication
//     Password - API password for authentication
//     merchantID - Merchant Account settings for payment and modification requests
//     clientID - Used to load external JS files from Adyen, to encrypt client requests
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
type apiCredentials struct {
	env         Environment
	Username    string
	Password    string
	merchantID  string
	clientID    string
	hppSettings hppCredentials
}

// hppCredentials used to communicate with Adyen HPP API
//
// Description:
//
//     hmac - is generated when new Skin is created in Adyen Customer Area
//     skinCode - skin code from Adyen CA
//     shopperLocale - merchant local settings (in ISO format)
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type hppCredentials struct {
	hmac          string
	skinCode      string
	shopperLocale string
}

// newCredentials create new APICredentials
func newCredentials(env Environment, username, password, merchantID, clientID string) apiCredentials {
	return apiCredentials{
		env:        env,
		Username:   username,
		Password:   password,
		merchantID: merchantID,
		clientID:   clientID,
	}
}

// newCredentialsWithHPPSettings create new APICredentials and specify Adyen Hosted Payment Page settings
func newCredentialsWithHPPSettings(env Environment, username, password, merchantID, clientID, hmac, skinCode, shopperLocale string) apiCredentials {
	return apiCredentials{
		env:        env,
		Username:   username,
		Password:   password,
		merchantID: merchantID,
		clientID:   clientID,
		hppSettings: hppCredentials{
			hmac:          hmac,
			skinCode:      skinCode,
			shopperLocale: shopperLocale,
		},
	}
}
