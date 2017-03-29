package adyen

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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

// execute request on Adyen side, transforms "requestEntity" into JSON representation
func (a *Adyen) execute(method string, requestEntity interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestEntity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", a.AdyenUrl(method), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.Username, a.Password)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Adyen) Authorise() *AuthoriseGateway {
	return &AuthoriseGateway{a}
}
