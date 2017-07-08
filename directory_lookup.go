package adyen

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
	Logos     Logos    `json:"logos"`
	Issuers   []Issuer `json:"issuers"`
}

// Logos - payment method logos
//
// Part of DirectoryLookupResponse
type Logos struct {
	Normal string `json:"normal"`
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
}

// Issuer - bank issuer type
//
// Part of DirectoryLookupResponse
type Issuer struct {
	IssuerID string `json:"issuerId"`
	Name     string `json:"name"`
}
