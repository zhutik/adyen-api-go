// Package adyen is Adyen API Library for GO
package adyen

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	// DefaultCurrency is the default currency for transactions
	DefaultCurrency = "EUR"

	// DefaultClientTimeout is the default timeout used when making
	// HTTP requests to Adyen.
	DefaultClientTimeout = time.Second * 10

	// APIVersion of current Adyen API
	APIVersion = "v25"

	// PaymentService is used to identify the standard payment workflow.
	PaymentService = "Payment"

	// RecurringService is used to identify the recurring payment workflow.
	RecurringService = "Recurring"
)

// Adyen - base structure with configuration options
//
//       - Credentials instance of API creditials to connect to Adyen API
//       - Currency is a default request currency. Request data overrides this setting
//       - MerchantAccount is default merchant account to be used. Request data overrides this setting
//       - Logger optionally logs to a configured io.Writer.
//
// Currency and MerchantAccount should be used only to store the data and be able to use it later.
// Requests won't be automatically populated with given values
type Adyen struct {
	Credentials     apiCredentials
	Currency        string
	MerchantAccount string
	ClientTimeout   time.Duration
	Logger          *log.Logger
}

// New - creates Adyen instance
//
// Description:
//
//     - env - Environment for next API calls
//     - username - API username for authentication
//     - password - API password for authentication
//     - logger optionally logs to a configured io.Writer.
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// You can create new API user there: https://ca-test.adyen.com/ca/ca/config/users.shtml
func New(env environment, username, password string, logger *log.Logger, opts ...Option) *Adyen {
	creds := makeCredentials(env, username, password)
	return NewWithCredentials(env, creds, logger, opts...)
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
//     - logger optionally logs to a configured io.Writer.
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
func NewWithHMAC(env environment, username, password, hmac string, logger *log.Logger, opts ...Option) *Adyen {
	creds := makeCredentialsWithHMAC(env, username, password, hmac)
	return NewWithCredentials(env, creds, logger, opts...)
}

// NewWithCredentials - create new Adyen instance with pre-configured credentials.
//
// Description:
//
//     - env - Environment for next API calls
//     - credentials - configured apiCredentials to use when interacting with Adyen.
//     - logger - optionally logs to a configured io.Writer.
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
// New skin can be created there https://ca-test.adyen.com/ca/ca/skin/skins.shtml
func NewWithCredentials(env environment, creds apiCredentials, logger *log.Logger, opts ...Option) *Adyen {
	a := Adyen{
		Credentials:   creds,
		Currency:      DefaultCurrency,
		ClientTimeout: DefaultClientTimeout,
		Logger:        logger,
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
		a.ClientTimeout = d
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
func (a *Adyen) adyenURL(service string, requestType string) string {
	return a.Credentials.Env.BaseURL(service, APIVersion) + "/" + requestType + "/"
}

// createHPPUrl returns Adyen HPP url
func (a *Adyen) createHPPUrl(requestType string) string {
	return a.Credentials.Env.HppURL(requestType)
}

// execute request on Adyen side, transforms "requestEntity" into JSON representation
//
// internal method to do a request to Adyen API endpoint
// request Type: POST, request body format - JSON
func (a *Adyen) execute(service string, method string, requestEntity interface{}) (*Response, error) {
	body, err := json.Marshal(requestEntity)
	if err != nil {
		return nil, err
	}

	url := a.adyenURL(service, method)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	a.Logger.Printf("[Request]: %s %s\n%s", method, url, body)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.Credentials.Username, a.Credentials.Password)

	client := &http.Client{
		Timeout: a.ClientTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	a.Logger.Printf("[Response]: %s %s\n%s", method, url, buf.String())

	providerResponse := &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	if err = providerResponse.handleHTTPError(); err != nil {
		return nil, err
	}

	return providerResponse, nil
}

// executeHpp - execute request without authorization to Adyen Hosted Payment API
//
// internal method to request Adyen HPP API via GET
func (a *Adyen) executeHpp(method string, requestEntity interface{}) (*Response, error) {
	url := a.createHPPUrl(method)

	v, _ := query.Values(requestEntity)
	url = url + "?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	a.Logger.Printf("[Request]: %s %s", method, url)

	client := &http.Client{
		Timeout: a.ClientTimeout,
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	a.Logger.Printf("[Response]: %s %s\n%s", method, url, buf.String())

	providerResponse := &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	return providerResponse, nil
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
