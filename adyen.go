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

const (
	APIVersion         = "25"
	AdyenTestUrl       = "https://pal-test.adyen.com/pal/servlet/Payment"
	AdyenClientTestUrl = "https://test.adyen.com/hpp/cse/js/"
)

func (a *Adyen) ClientURL() string {
	return AdyenClientTestUrl + a.ClientID + ".shtml"
}

func (a *Adyen) AdyenUrl(requestType string) string {
	return AdyenTestUrl + "/" + APIVersion + "/" + requestType
}

func (a *Adyen) Authorise() *AuthoriseGateway {
	return &AuthoriseGateway{a}
}
