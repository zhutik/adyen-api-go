package adyen

// Cancel structure for Cancel request
type Cancel struct {
	Reference         string `json:"reference"`
	MerchantAccount   string `json:"merchantAccount"`
	OriginalReference string `json:"originalReference"`
}

// CancelResponse is a response structure for Adyen cancellation
type CancelResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}
