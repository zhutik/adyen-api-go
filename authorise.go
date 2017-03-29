package adyen

type AuthoriseRequest struct {
	AdditionalData  AdditionalData `json:"additionalData"`
	Amount          Amount         `json:"amount"`
	Reference       string         `json:"reference"`
	MerchantAccount string         `json:"merchantAccount"`
}

type AuthoriseResponse struct {
	PspReference  string `json:"pspReference"`
	ResultCode    string `json:"resultCode"`
	AuthCode      string `json:"authCode"`
	RefusalReason string `json:"refusalReason"`
}

type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

type AdditionalData struct {
	Content string `json:"card.encrypted.json"`
}
