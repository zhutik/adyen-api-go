package adyen

// CaptureRequest structure for Capture request
type CaptureRequest struct {
	ModificationAmount Amount `json:"modificationAmount"`
	Reference          string `json:"reference"`
	MerchantAccount    string `json:"merchantAccount"`
	OriginalReference  string `json:"originalReference"`
}

// CaptureResponse is a response structure for Adyen capture
type CaptureResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}
