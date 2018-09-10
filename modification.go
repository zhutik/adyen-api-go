package adyen

/*********
* Cancel *
*********/

// Adjust authorisation reasons
//
// Link https://docs.adyen.com/developers/api-reference/payments-api/modificationrequest/adjustauthorisationmodificationrequest
const (
	DelayedCharge = "DelayedCharge"
	NoShow        = "NoShow"
)

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

/***********************
* Adjust Authorisation *
***********************/

// AdjustAuthorisation structure for adjusting previously authorised amount
type AdjustAuthorisation struct {
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	MerchantAccount    string  `json:"merchantAccount"`
	OriginalReference  string  `json:"originalReference"`
	AdditionalData     struct {
		IndustryUsage string `json:"industryUsage"`
	} `json:"additionalData,omitempty"`
}

// AdjustAuthorisationResponse is a response for AdjustAuthorisation request
type AdjustAuthorisationResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}

/******************
* Techical Cancel *
******************/

// TechnicalCancel structure for performing technical cancellation
//
// Link - https://docs.adyen.com/developers/payment-modifications#technicalcancel
type TechnicalCancel struct {
	MerchantAccount           string `json:"merchantAccount"`
	OriginalMerchantReference string `json:"originalMerchantReference"`
	Reference                 string `json:"reference,omitempty"`
}

// TechnicalCancelResponse is a response for TechnicalCancel request
type TechnicalCancelResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}
