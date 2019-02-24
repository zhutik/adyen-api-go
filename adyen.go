// Package adyen is Adyen API Library for GO
package adyen

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	// DefaultCurrency is the default currency for transactions
	DefaultCurrency = "EUR"

	// DefaultClientTimeout is the default timeout used when making
	// HTTP requests to Adyen.
	DefaultClientTimeout = time.Second * 10

	// PaymentAPIVersion - API version of current payment API
	PaymentAPIVersion = "v40"

	// RecurringAPIVersion - API version of current recurring API
	RecurringAPIVersion = "v25"

	// PaymentService is used to identify the standard payment workflow.
	PaymentService = "Payment"

	// RecurringService is used to identify the recurring payment workflow.
	RecurringService = "Recurring"

	// CheckoutAPIVersion - API version of current checkout API
	CheckoutAPIVersion = "v32"
)

// Adyen - base structure with configuration options
//
//       - Credentials instance of API creditials to connect to Adyen API
//       - Currency is a default request currency. Request data overrides this setting
//       - MerchantAccount is default merchant account to be used. Request data overrides this setting
//       - client is http client instance
//
// Currency and MerchantAccount should be used only to store the data and be able to use it later.
// Requests won't be automatically populated with given values
type Adyen struct {
	Credentials     apiCredentials
	Currency        string
	MerchantAccount string

	client *http.Client
}

// New - creates Adyen instance
//
// Description:
//
//     - env - Environment for next API calls
//     - username - API username for authentication
//     - password - API password for authentication
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
func New(env Environment, username, password string, opts ...Option) *Adyen {
	creds := makeCredentials(env, username, password)
	return NewWithCredentials(env, creds, opts...)
}

// NewWithHMAC - create new Adyen instance with HPP credentials
//
// Use this constructor when you need to use Adyen HPP API.
//
// Description:
//
//     - env - Environment for next API calls
//     - username - API username for authentication
//     - password - API password for authentication
//     - hmac - is generated when new Skin is created in Adyen Customer Area
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
func NewWithHMAC(env Environment, username, password, hmac string, opts ...Option) *Adyen {
	creds := makeCredentialsWithHMAC(env, username, password, hmac)
	return NewWithCredentials(env, creds, opts...)
}

// NewWithCredentials - create new Adyen instance with pre-configured credentials.
//
// Description:
//
//     - env - Environment for next API calls
//     - credentials - configured apiCredentials to use when interacting with Adyen.
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
func NewWithCredentials(env Environment, creds apiCredentials, opts ...Option) *Adyen {
	a := Adyen{
		Credentials: creds,
		Currency:    DefaultCurrency,
		client:      &http.Client{},
	}

	if opts != nil {
		for _, opt := range opts {
			opt(&a)
		}
	}

	return &a
}

// Option allows for custom configuration overrides.
type Option func(*Adyen)

// WithTimeout allows for a custom timeout to be provided to the underlying
// HTTP client that's used to communicate with Adyen.
func WithTimeout(d time.Duration) func(*Adyen) {
	return func(a *Adyen) {
		a.client.Timeout = d
	}
}

// WithTransport allows customer HTTP transports to be provider to the Adyen
func WithTransport(transport http.RoundTripper) func(*Adyen) {
	return func(a *Adyen) {
		a.client.Transport = transport
	}
}

// WithCurrency allows for custom currencies to be provided to the Adyen.
func WithCurrency(c string) func(*Adyen) {
	return func(a *Adyen) {
		a.Currency = c
	}
}

// ClientURL - returns URl, that need to loaded in UI, to encrypt Credit Card information
//
//           - clientID - Used to load external JS files from Adyen, to encrypt client requests
func (a *Adyen) ClientURL(clientID string) string {
	return a.Credentials.Env.ClientURL(clientID)
}

// adyenURL returns Adyen backend URL
func (a *Adyen) adyenURL(service, requestType, apiVersion string) string {
	return a.Credentials.Env.BaseURL(service, apiVersion) + "/" + requestType + "/"
}

// createHPPUrl returns Adyen HPP url
func (a *Adyen) createHPPUrl(requestType string) string {
	return a.Credentials.Env.HppURL(requestType)
}

// checkoutURL returns the Adyen checkout URL.
func (a *Adyen) checkoutURL(requestType, apiVersion string) string {
	return a.Credentials.Env.CheckoutURL(requestType, apiVersion)
}

// execute request on Adyen side, transforms "requestEntity" into JSON representation
//
// internal method to do a request to Adyen API endpoint
// request Type: POST, request body format - JSON
func (a *Adyen) execute(url string, requestEntity interface{}) (r *Response, err error) {
	body, err := json.Marshal(requestEntity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.Credentials.Username, a.Credentials.Password)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = cerr
		}
	}()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	r = &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	if err = r.handleHTTPError(); err != nil {
		return nil, err
	}

	return
}

// executeHpp - execute request without authorization to Adyen Hosted Payment API
//
// internal method to request Adyen HPP API via GET
func (a *Adyen) executeHpp(url string, requestEntity interface{}) (r *Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = cerr
		}
	}()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	r = &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	return
}

// Payment - returns PaymentGateway
func (a *Adyen) Payment() *PaymentGateway {
	return &PaymentGateway{a}
}

// Modification - returns ModificationGateway
func (a *Adyen) Modification() *ModificationGateway {
	return &ModificationGateway{a}
}

// Recurring - returns RecurringGateway
func (a *Adyen) Recurring() *RecurringGateway {
	return &RecurringGateway{a}
}

// Checkout - returns CheckoutGateway
func (a *Adyen) Checkout() *CheckoutGateway {
	return &CheckoutGateway{a}
}
