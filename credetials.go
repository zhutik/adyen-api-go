package adyen

// apiCredentials basic API settings
//
// Description:
//
//     - Env - Environment for next API calls
//     - Username - API username for authentication
//     - Password - API password for authentication
//     - MerchantID - Merchant Account settings for payment and modification requests
//     - Hmac - Hash-based Message Authentication Code (HMAC) setting
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type apiCredentials struct {
	Env        environment
	Username   string
	Password   string
	MerchantID string
	Hmac       string
}

// newCredentials create new APICredentials
func newCredentials(env environment, username, password, merchantID string) apiCredentials {
	return apiCredentials{
		Env:        env,
		Username:   username,
		Password:   password,
		MerchantID: merchantID,
	}
}

// newCredentialsWithHMAC create new APICredentials with HMAC singature
func newCredentialsWithHMAC(env environment, username, password, merchantID, hmac string) apiCredentials {
	return apiCredentials{
		Env:        env,
		Username:   username,
		Password:   password,
		MerchantID: merchantID,
		Hmac:       hmac,
	}
}
