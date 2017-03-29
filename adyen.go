package adyen

func New(username, password, clientId, merchantAccount string) *Adyen {
	return &Adyen{
		Username:        username,
		Password:        password,
		ClientID:        clientId,
		MerchantAccount: merchantAccount,
	}
}

type Adyen struct {
	Username        string
	Password        string
	ClientID        string
	MerchantAccount string
}

func (a *Adyen) ClientURL() string {
	return "https://test.adyen.com/hpp/cse/js/" + a.ClientID + ".shtml"
}

func (a *Adyen) Authorise() *AuthoriseGateway {
	return &AuthoriseGateway{a}
}
