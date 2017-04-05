package adyen

// Authorise structure for Authorisation request
type Authorise struct {
	AdditionalData  *AdditionalData `json:"additionalData"`
	Amount          *Amount         `json:"amount"`
	Reference       string          `json:"reference"`
	MerchantAccount string          `json:"merchantAccount"`
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

// Capture structure for Capture request
type Capture struct {
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	MerchantAccount    string  `json:"merchantAccount"`
	OriginalReference  string  `json:"originalReference"`
}

// CaptureResponse is a response structure for Adyen capture
type CaptureResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}
