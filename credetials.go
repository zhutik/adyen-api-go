package adyen

// apiCredentials basic API settings
//
// Description:
//
//     - Env - Environment for next API calls
//     - Username - API username for authentication
//     - Password - API password for authentication
//     - MerchantID - Merchant Account settings for payment and modification requests
//     - ClientID - Used to load external JS files from Adyen, to encrypt client requests
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
type apiCredentials struct {
	Env         environment
	Username    string
	Password    string
	MerchantID  string
	ClientID    string
	HppSettings hppCredentials
}

// hppCredentials used to communicate with Adyen HPP API
//
// Description:
//
//     - Hmac - is generated when new Skin is created in Adyen Customer Area
//     - SkinCode - skin code from Adyen CA
//     - ShopperLocale - merchant local settings (in ISO format)
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type hppCredentials struct {
	Hmac          string
	SkinCode      string
	ShopperLocale string
}

// newCredentials create new APICredentials
func newCredentials(env environment, username, password, merchantID, clientID string) apiCredentials {
	return apiCredentials{
		Env:        env,
		Username:   username,
		Password:   password,
		MerchantID: merchantID,
		ClientID:   clientID,
	}
}

// newCredentialsWithHPPSettings create new APICredentials and specify Adyen Hosted Payment Page settings
func newCredentialsWithHPPSettings(env environment, username, password, merchantID, clientID, hmac, skinCode, shopperLocale string) apiCredentials {
	return apiCredentials{
		Env:        env,
		Username:   username,
		Password:   password,
		MerchantID: merchantID,
		ClientID:   clientID,
		HppSettings: hppCredentials{
			Hmac:          hmac,
			SkinCode:      skinCode,
			ShopperLocale: shopperLocale,
		},
	}
}
