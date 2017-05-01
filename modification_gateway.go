package adyen

import "encoding/json"

/*
  Adyen Modification actions
*/
const (
	CaptureType        = "capture"
	CancelType         = "cancel"
	CancelOrRefundType = "cancelOrRefund"
	RefundType         = "refund"
)

// ModificationGateway - Adyen modification transaction logic, capture, cancel, refunds and e.t.c
type ModificationGateway struct {
	*Adyen
}

// Capture - Perform capture payment in Adyen
func (a *ModificationGateway) Capture(req *Capture) (*CaptureResponse, error) {
	resp, err := a.execute(CaptureType, req)

	if err != nil {
		return nil, err
	}

	var val CaptureResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}

// Cancel - Perform cancellation of the authorised transaction
func (a *ModificationGateway) Cancel(req *Cancel) (*CancelResponse, error) {
	resp, err := a.execute(CancelType, req)

	if err != nil {
		return nil, err
	}

	var val CancelResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}

// CancelOrRefund - Perform cancellation for not captured transaction
// otherwise perform refund action
func (a *ModificationGateway) CancelOrRefund(req *Cancel) (*CancelOrRefundResponse, error) {
	resp, err := a.execute(CancelOrRefundType, req)

	if err != nil {
		return nil, err
	}

	var val CancelOrRefundResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}

// Refund - perform refund for already captured request
func (a *ModificationGateway) Refund(req *Refund) (*RefundResponse, error) {
	resp, err := a.execute(RefundType, req)

	if err != nil {
		return nil, err
	}

	var val RefundResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}
