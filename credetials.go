package adyen

// apiCredentials contains
type apiCredentials struct {
	env        Environment
	Username   string
	Password   string
	merchantID string
	clientID   string
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
