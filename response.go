package adyen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	ErrorType string `json:"errorType"`
	ErrorCode int32  `json:"errorCode"`
	Message   string `json:"message"`
	Status    int32  `json:"status"`
}

type Response struct {
	*http.Response
	Body []byte
}

func (r *Response) checkError() error {
	var a ApiError

	json.Unmarshal(r.Body, &a)
	if a.Status > 299 {
		return httpError(a.Status)
	}

	return nil
}

type httpError int

func (e httpError) StatusCode() int {
	return int(e)
}

func (e httpError) Error() string {
	return fmt.Sprintf("%s (%d)", http.StatusText(e.StatusCode()), e.StatusCode())
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
