package adyen

// Adyen Modification actions
const (
	captureType        = "capture"
	cancelType         = "cancel"
	cancelOrRefundType = "cancelOrRefund"
	refundType         = "refund"
)

// ModificationGateway - Adyen modification transaction logic, capture, cancel, refunds and e.t.c
type ModificationGateway struct {
	*Adyen
}

// Capture - Perform capture payment in Adyen
func (a *ModificationGateway) Capture(req *Capture) (*CaptureResponse, error) {
	resp, err := a.execute(PaymentService, captureType, req)

	if err != nil {
		return nil, err
	}

	return resp.capture()
}

// Cancel - Perform cancellation of the authorised transaction
func (a *ModificationGateway) Cancel(req *Cancel) (*CancelResponse, error) {
	resp, err := a.execute(PaymentService, cancelType, req)

	if err != nil {
		return nil, err
	}

	return resp.cancel()
}

// CancelOrRefund - Perform cancellation for not captured transaction
// otherwise perform refund action
func (a *ModificationGateway) CancelOrRefund(req *Cancel) (*CancelOrRefundResponse, error) {
	resp, err := a.execute(PaymentService, cancelOrRefundType, req)

	if err != nil {
		return nil, err
	}

	return resp.cancelOrRefund()
}

// Refund - perform refund for already captured request
func (a *ModificationGateway) Refund(req *Refund) (*RefundResponse, error) {
	resp, err := a.execute(PaymentService, refundType, req)

	if err != nil {
		return nil, err
	}

	return resp.refund()
}
