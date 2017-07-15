package adyen

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
