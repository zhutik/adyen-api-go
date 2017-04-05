package adyen

import "encoding/json"

// CaptureType - capture type request, @TODO: move to enums
const CaptureType = "capture"

// CancelType - cancel type request, @TODO: move to enums
const CancelType = "cancel"

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
