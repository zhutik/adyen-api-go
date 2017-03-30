package adyen

// AuthoriseRequest structure for Authorisation request
type AuthoriseRequest struct {
	AdditionalData  AdditionalData `json:"additionalData"`
	Amount          Amount         `json:"amount"`
	Reference       string         `json:"reference"`
	MerchantAccount string         `json:"merchantAccount"`
}

// AuthoriseResponse is a response structure for Adyen
type AuthoriseResponse struct {
	PspReference  string `json:"pspReference"`
	ResultCode    string `json:"resultCode"`
	AuthCode      string `json:"authCode"`
	RefusalReason string `json:"refusalReason"`
}

// AdditionalData stores encrypted information about customer's credit card
type AdditionalData struct {
	Content string `json:"card.encrypted.json"`
}
