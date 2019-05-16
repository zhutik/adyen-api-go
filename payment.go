package adyen

/**********
* Payment *
**********/

// One-click functionality gives the shopper the option to store their payment details with the merchant, within the Adyen environment.
//
// In this type of transaction, the shopper needs to enter the CVC code for the transaction to get through.
//
// Link: https://docs.adyen.com/developers/api-reference/payments-api#recurring
const (
	RecurringPaymentOneClick              = "ONECLICK"
	RecurringPaymentRecurring             = "RECURRING"
	ShopperInteractionContAuth            = "ContAuth"
	SelectRecurringDetailReferenceLatests = "LATEST"
)

// AuthoriseEncrypted structure for Authorisation request (with encrypted card information)
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest
type AuthoriseEncrypted struct {
	AdditionalData                   *AdditionalData `json:"additionalData,omitempty"`
	Amount                           *Amount         `json:"amount"`
	BillingAddress                   *Address        `json:"billingAddress,omitempty"`
	DeliveryAddress                  *Address        `json:"deliveryAddress,omitempty"`
	Reference                        string          `json:"reference"`
	MerchantAccount                  string          `json:"merchantAccount"`
	ShopperReference                 string          `json:"shopperReference,omitempty"` // Mandatory for recurring payment
	Recurring                        *Recurring      `json:"recurring,omitempty"`
	ShopperEmail                     string          `json:"shopperEmail,omitempty"`
	ShopperInteraction               string          `json:"shopperInteraction,omitempty"`
	ShopperIP                        string          `json:"shopperIP,omitempty"`
	ShopperLocale                    string          `json:"shopperLocale,omitempty"`
	ShopperName                      *Name           `json:"shopperName,omitempty"`
	SelectedRecurringDetailReference string          `json:"selectedRecurringDetailReference,omitempty"`
	BrowserInfo                      *BrowserInfo    `json:"browserInfo,omitempty"` // Required for a 3DS process
	CaptureDelayHours                *int            `json:"captureDelayHours,omitempty"`
	FundingSource                    string          `json:"fundingSource,omitempty"`
}

// Authorise structure for Authorisation request (card is not encrypted)
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest
type Authorise struct {
	Card                             *Card        `json:"card,omitempty"`
	Amount                           *Amount      `json:"amount"`
	BillingAddress                   *Address     `json:"billingAddress,omitempty"`
	DeliveryAddress                  *Address     `json:"deliveryAddress,omitempty"`
	Reference                        string       `json:"reference"`
	MerchantAccount                  string       `json:"merchantAccount"`
	ShopperReference                 string       `json:"shopperReference,omitempty"` // Mandatory for recurring payment
	Recurring                        *Recurring   `json:"recurring,omitempty"`
	ShopperEmail                     string       `json:"shopperEmail,omitempty"`
	ShopperInteraction               string       `json:"shopperInteraction,omitempty"`
	ShopperIP                        string       `json:"shopperIP,omitempty"`
	ShopperLocale                    string       `json:"shopperLocale,omitempty"`
	ShopperName                      *Name        `json:"shopperName,omitempty"`
	SelectedRecurringDetailReference string       `json:"selectedRecurringDetailReference,omitempty"`
	BrowserInfo                      *BrowserInfo `json:"browserInfo,omitempty"` // Required for a 3DS process
	CaptureDelayHours                *int         `json:"captureDelayHours,omitempty"`
}

// AuthoriseResponse is a response structure for Adyen
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentresult
type AuthoriseResponse struct {
	PspReference   string          `json:"pspReference"`
	ResultCode     string          `json:"resultCode"`
	AuthCode       string          `json:"authCode"`
	RefusalReason  string          `json:"refusalReason"`
	IssuerURL      string          `json:"issuerUrl"`
	MD             string          `json:"md"`
	PaRequest      string          `json:"paRequest"`
	AdditionalData *AdditionalData `json:"additionalData,omitempty"`
}

// AdditionalData stores encrypted information about customer's credit card
type AdditionalData struct {
	Content                  string      `json:"card.encrypted.json,omitempty"`
	AliasType                string      `json:"aliasType,omitempty"`
	Alias                    string      `json:"alias,omitempty"`
	ExpiryDate               string      `json:"expiryDate,omitempty"`
	CardBin                  string      `json:"cardBin,omitempty"`
	CardSummary              string      `json:"cardSummary,omitempty"`
	PaymentMethod            string      `json:"paymentMethod,omitempty"`
	CardPaymentMethod        string      `json:"cardPaymentMethod,omitempty"`
	CardIssuingCountry       string      `json:"cardIssuingCountry,omitempty"`
	RecurringDetailReference string      `json:"recurring.recurringDetailReference,omitempty"`
	ExecuteThreeD            *StringBool `json:"executeThreeD,omitempty"`
}

// BrowserInfo hold information on the user browser
type BrowserInfo struct {
	AcceptHeader string `json:"acceptHeader"`
	UserAgent    string `json:"userAgent"`
}

// Recurring hold the behavior for a future payment : could be ONECLICK or RECURRING
type Recurring struct {
	Contract string `json:"contract"`
}

/*************
* Payment 3D *
*************/

// Authorise3D structure for Authorisation request (card is not encrypted)
//
// https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest3d
type Authorise3D struct {
	BillingAddress  *Address     `json:"billingAddress,omitempty"`
	DeliveryAddress *Address     `json:"deliveryAddress,omitempty"`
	MD              string       `json:"md"`
	MerchantAccount string       `json:"merchantAccount"`
	BrowserInfo     *BrowserInfo `json:"browserInfo"`
	PaResponse      string       `json:"paResponse"`
	ShopperEmail    string       `json:"shopperEmail,omitempty"`
	ShopperIP       string       `json:"shopperIP,omitempty"`
	ShopperLocale   string       `json:"shopperLocale,omitempty"`
	ShopperName     *Name        `json:"shopperName,omitempty"`
}

/*******************
* Directory lookup *
*******************/

// DirectoryLookupRequest - get list of available payment methods based on skin, country and order details
//
// Description - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#directoryrequest
// CountryCode could be used to test local payment methods, if client's IP is from different country
type DirectoryLookupRequest struct {
	CurrencyCode      string `url:"currencyCode"`
	MerchantAccount   string `url:"merchantAccount"`
	PaymentAmount     int    `url:"paymentAmount"`
	SkinCode          string `url:"skinCode"`
	MerchantReference string `url:"merchantReference"`
	SessionsValidity  string `url:"sessionValidity"`
	MerchantSig       string `url:"merchantSig"`
	CountryCode       string `url:"countryCode"`
	ShipBeforeDate    string `url:"shipBeforeDate"`
}

// DirectoryLookupResponse - api response for DirectoryLookupRequest
//
// Description - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#directoryresponse
type DirectoryLookupResponse struct {
	PaymentMethods []PaymentMethod `json:"paymentMethods"`
}

// PaymentMethod - structure for single payment method in directory look up response
//
// Part of DirectoryLookupResponse
type PaymentMethod struct {
	BrandCode string   `json:"brandCode"`
	Name      string   `json:"name"`
	Logos     logos    `json:"logos"`
	Issuers   []issuer `json:"issuers"`
}

// logos - payment method logos
//
// Part of DirectoryLookupResponse
type logos struct {
	Normal string `json:"normal"`
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
}

// issuer - bank issuer type
//
// Part of DirectoryLookupResponse
type issuer struct {
	IssuerID string `json:"issuerId"`
	Name     string `json:"name"`
}

/***********
* Skip HPP *
***********/

// SkipHppRequest contains data that would be used to create Adyen HPP redirect URL
//
// Link: https://docs.adyen.com/developers/ecommerce-integration/local-payment-methods
//
// Request description: https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#skipdetailsrequest
type SkipHppRequest struct {
	MerchantReference string `url:"merchantReference"`
	PaymentAmount     int    `url:"paymentAmount"`
	CurrencyCode      string `url:"currencyCode"`
	ShipBeforeDate    string `url:"shipBeforeDate"`
	SkinCode          string `url:"skinCode"`
	MerchantAccount   string `url:"merchantAccount"`
	ShopperLocale     string `url:"shopperLocale"`
	SessionsValidity  string `url:"sessionValidity"`
	MerchantSig       string `url:"merchantSig"`
	CountryCode       string `url:"countryCode"`
	BrandCode         string `url:"brandCode"`
	IssuerID          string `url:"issuerId"`
}
