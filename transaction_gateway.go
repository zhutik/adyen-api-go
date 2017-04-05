package adyen

import "encoding/json"

// TransactionGateway - Adyen payment transaction logic
type TransactionGateway struct {
	*Adyen
}

// AuthoriseType - authorise type request, @TODO: move to enums
const AuthoriseType = "authorise"

// CaptureType - capture type request, @TODO: move to enums
const CaptureType = "capture"

// Authorise - Perform authorise payment in Adyen
func (a *TransactionGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	resp, err := a.execute(AuthoriseType, req)

	if err != nil {
		return nil, err
	}

	var val AuthoriseResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}

// Capture - Perform capture payment in Adyen
func (a *TransactionGateway) Capture(req *Capture) (*CaptureResponse, error) {
	resp, err := a.execute(CaptureType, req)

	if err != nil {
		return nil, err
	}

	var val CaptureResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}
