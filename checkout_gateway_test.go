package adyen

import (
	"os"
	"testing"
)

// TestPaymentMethods - test for https://docs.adyen.com/developers/checkout/api-integration
//
// This test requires CheckoutAPI access.  To obtain, visit https://docs.adyen.com/developers/user-management/how-to-get-the-checkout-api-key.
func TestPaymentMethods(t *testing.T) {
	t.Parallel()

	instance := getTestInstance()

	request := &PaymentMethods{
		MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
	}

	_, err := instance.Checkout().PaymentMethods(request)

	knownError, ok := err.(apiError)
	if ok {
		t.Errorf("Response should be succesfull. Known API Error: Code - %s, Message - %s, Type - %s", knownError.ErrorCode, knownError.Message, knownError.ErrorType)
	}

	if err != nil {
		t.Errorf("Response should be succesfull, error - %s", err.Error())
	}
}
