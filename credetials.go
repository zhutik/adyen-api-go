package adyen

// apiCredentials basic API settings
//
// Description:
//
//     - Env - Environment for next API calls
//     - Username - API username for authentication
//     - Password - API password for authentication
//     - Hmac - Hash-based Message Authentication Code (HMAC) setting
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
type apiCredentials struct {
	Env      environment
	Username string
	Password string
	Hmac     string
}

// makeCredentials create new APICredentials
func makeCredentials(env environment, username, password string) apiCredentials {
	return apiCredentials{
		Env:      env,
		Username: username,
		Password: password,
	}
}

// makeCredentialsWithHMAC create new APICredentials with HMAC singature
func makeCredentialsWithHMAC(env environment, username, password, hmac string) apiCredentials {
	return apiCredentials{
		Env:      env,
		Username: username,
		Password: password,
		Hmac:     hmac,
	}
}
