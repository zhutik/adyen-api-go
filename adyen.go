package adyen

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// DefaultCurrency default currency for transactions
const DefaultCurrency = "EUR"

// New - creates Adyen instance
func New(username, password, clientID, merchantAccount string) *Adyen {
	return &Adyen{
		Username:        username,
		Password:        password,
		ClientID:        clientID,
		MerchantAccount: merchantAccount,
		Currency:        DefaultCurrency,
	}
}

// Adyen - base structure with configuration options
type Adyen struct {
	Username        string
	Password        string
	ClientID        string
	MerchantAccount string
	Currency        string
}

// Version of a current Adyen API
const (
	APIVersion         = "v25"
	AdyenTestURL       = "https://pal-test.adyen.com/pal/servlet/Payment"
	AdyenClientTestURL = "https://test.adyen.com/hpp/cse/js/"
)

// ClientURL - returns URl, that need to loaded in UI, to encrypt Credit Card information
func (a *Adyen) ClientURL() string {
	return AdyenClientTestURL + a.ClientID + ".shtml"
}

// AdyenURL returns Adyen backend URL
func (a *Adyen) AdyenURL(requestType string) string {
	return AdyenTestURL + "/" + APIVersion + "/" + requestType
}

// SetCurrency set default currency for transactions
func (a *Adyen) SetCurrency(currency string) {
	a.Currency = currency
}

// execute request on Adyen side, transforms "requestEntity" into JSON representation
func (a *Adyen) execute(method string, requestEntity interface{}) (*http.Response, error) {
	body, err := json.Marshal(requestEntity)
	if err != nil {
		return nil, err
	}

	req, err2 := http.NewRequest("POST", a.AdyenURL(method), bytes.NewBuffer(body))
	if err2 != nil {
		return nil, err2
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.Username, a.Password)

	client := &http.Client{}
	resp, err3 := client.Do(req)

	if err3 != nil {
		return nil, err3
	}

	return resp, nil
}

// Payment - returns PaymentGateway
func (a *Adyen) Payment() *PaymentGateway {
	return &PaymentGateway{a}
}

// Modification - returns ModificationGateway
func (a *Adyen) Modification() *ModificationGateway {
	return &ModificationGateway{a}
}
