package adyen

// CheckoutGateway - allows you to accept all of Adyen's payment
// methods and flows.
type CheckoutGateway struct {
	*Adyen
}

const (
	paymentMethodsURL = "paymentMethods"
)

// PaymentMethods - Perform paymentMethods request in Adyen.
//
// Used to get a collection of available payment methods for a merchant.
func (a *CheckoutGateway) PaymentMethods(req *PaymentMethods) (*PaymentMethodsResponse, error) {
	url := a.checkoutURL(paymentMethodsURL, CheckoutAPIVersion)

	resp, err := a.execute(url, req)
	if err != nil {
		return nil, err
	}

	return resp.paymentMethods()
}
