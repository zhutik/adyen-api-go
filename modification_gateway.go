package adyen

import "encoding/json"

// CaptureType - capture type request, @TODO: move to enums
const CaptureType = "capture"

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
