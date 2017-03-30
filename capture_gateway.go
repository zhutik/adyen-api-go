package adyen

import "encoding/json"

// CaptureGateway - Adyen payment capture logic
type CaptureGateway struct {
	*Adyen
}

// CaptureType - authorise type request, @TODO: move to enums
const CaptureType = "capture"

// Payment - Perform authorise payment in Adyen
func (a *CaptureGateway) Payment(originalReference string, reference string, amount float32) (*CaptureResponse, error) {
	auth := CaptureRequest{
		ModificationAmount: Amount{Value: amount, Currency: "EUR"},
		MerchantAccount:    a.MerchantAccount,
		Reference:          reference,
		OriginalReference:  originalReference,
	}

	resp, err := a.execute(CaptureType, auth)

	if err != nil {
		return nil, err
	}

	var val CaptureResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}
