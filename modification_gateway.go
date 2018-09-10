package adyen

// Adyen Modification actions
const (
	captureType         = "capture"
	cancelType          = "cancel"
	cancelOrRefundType  = "cancelOrRefund"
	refundType          = "refund"
	adjustAuthorisation = "adjustAuthorisation"
	technicalCancel     = "technicalCancel"
)

// ModificationGateway - Adyen modification transaction logic, capture, cancel, refunds and e.t.c
type ModificationGateway struct {
	*Adyen
}

// Capture - Perform capture payment in Adyen
func (a *ModificationGateway) Capture(req *Capture) (*CaptureResponse, error) {
	url := a.adyenURL(PaymentService, captureType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.capture()
}

// Cancel - Perform cancellation of the authorised transaction
func (a *ModificationGateway) Cancel(req *Cancel) (*CancelResponse, error) {
	url := a.adyenURL(PaymentService, cancelType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.cancel()
}

// CancelOrRefund - Perform cancellation for not captured transaction
// otherwise perform refund action
func (a *ModificationGateway) CancelOrRefund(req *Cancel) (*CancelOrRefundResponse, error) {
	url := a.adyenURL(PaymentService, cancelOrRefundType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.cancelOrRefund()
}

// Refund - perform refund for already captured request
func (a *ModificationGateway) Refund(req *Refund) (*RefundResponse, error) {
	url := a.adyenURL(PaymentService, refundType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.refund()
}

// AdjustAuthorisation - perform adjustAuthorisation request to modify already authorised amount
//
// Link - https://docs.adyen.com/developers/payment-modifications#adjustauthorisation
func (a *ModificationGateway) AdjustAuthorisation(req *AdjustAuthorisation) (*AdjustAuthorisationResponse, error) {
	url := a.adyenURL(PaymentService, adjustAuthorisation, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.adjustAuthorisation()
}

// TechnicalCancel - perform cancellation without knowing orinal payment reference (PSP), f.e. in case of technical error
//
// Link - https://docs.adyen.com/developers/payment-modifications#technicalcancel
func (a *ModificationGateway) TechnicalCancel(req *TechnicalCancel) (*TechnicalCancelResponse, error) {
	url := a.adyenURL(PaymentService, technicalCancel, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.technicalCancel()
}
