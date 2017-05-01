package adyen

// Refund structure for refund request
type Refund struct {
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	MerchantAccount    string  `json:"merchantAccount"`
	OriginalReference  string  `json:"originalReference"`
}

// RefundResponse is a response structure for Adyen refund request
type RefundResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}
