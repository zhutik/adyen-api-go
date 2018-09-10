package adyen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIError - handle error (non 200 status) response from Adyen
type APIError struct {
	ErrorType string `json:"errorType"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Status    int32  `json:"status"`
}

// Response - Adyen API response structure
type Response struct {
	*http.Response
	Body []byte
}

// handleHTTPError - handle non 200 response from Adyen and create Error response instance
func (r *Response) handleHTTPError() error {
	var aerr APIError
	if err := json.Unmarshal(r.Body, &aerr); err != nil {
		return err
	}

	if aerr.Status >= http.StatusBadRequest {
		return aerr
	}

	return nil
}

// Error - error interface for ApiError
func (e APIError) Error() string {
	return fmt.Sprintf("[%s][%d]: (%s) %s", e.ErrorType, e.Status, e.ErrorCode, e.Message)
}

// authorize - generate Adyen Authorize API Response
func (r *Response) authorize() (*AuthoriseResponse, error) {
	var a AuthoriseResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// capture - generate Adyen Capture API Response
func (r *Response) capture() (*CaptureResponse, error) {
	var a CaptureResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// cancel - generate Adyen Cancel API Response
func (r *Response) cancel() (*CancelResponse, error) {
	var a CancelResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// cancelOrRefund - generate Adyen CancelOrRefund API Response
func (r *Response) cancelOrRefund() (*CancelOrRefundResponse, error) {
	var a CancelOrRefundResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// refund - generate Adyen Refund API Response
func (r *Response) refund() (*RefundResponse, error) {
	var a RefundResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// adjustAuthorisation - generate Adyen Refund API Response
func (r *Response) adjustAuthorisation() (*AdjustAuthorisationResponse, error) {
	var a AdjustAuthorisationResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// technicalCancel - generate Adyen Technical Cancel API Response
func (r *Response) technicalCancel() (*TechnicalCancelResponse, error) {
	var a TechnicalCancelResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// directoryLookup - generate Adyen Directory Lookup response
func (r *Response) directoryLookup() (*DirectoryLookupResponse, error) {
	var a DirectoryLookupResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// listRecurringDetails  - generate Adyen List Recurring Details response
func (r *Response) listRecurringDetails() (*RecurringDetailsResult, error) {
	var a RecurringDetailsResult
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// disableRecurring  - generate Adyen disable recurring
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#disableresult
func (r *Response) disableRecurring() (*RecurringDisableResponse, error) {
	var a RecurringDisableResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// paymentMethods - generate Adyen CheckoutAPI paymentMethods response.
func (r *Response) paymentMethods() (*PaymentMethodsResponse, error) {
	var a PaymentMethodsResponse
	if err := json.Unmarshal(r.Body, &a); err != nil {
		return nil, err
	}

	return &a, nil
}
