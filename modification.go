package adyen

/*********
* Cancel *
*********/

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

// CancelOrRefundResponse is a response structure for Adyen cancelOrRefund
type CancelOrRefundResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}

/**********
* Capture *
**********/

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

/*********
* Refund *
*********/

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
