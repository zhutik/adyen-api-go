package adyen

import (
	"os"
	"strings"
	"testing"
)

// TestAuthoriseFailed
func TestAuthoriseFailed(t *testing.T) {
	t.Parallel()

	instance := getTestInstance()

	authRequest := &Authorise{
		Card: &Card{
			Number:      "4111111111111111",
			ExpireMonth: "08",
			ExpireYear:  "2018",
			Cvc:         "737",
			HolderName:  "John Smith",
		},
		Amount: &Amount{
			Value:    1000,
			Currency: "EUR",
		},
		Reference:       "",
		MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
	}

	_, err := instance.Payment().Authorise(authRequest)

	if err == nil {
		t.Error("Request should fail, due to missing reference error")
	}

	if !strings.Contains(err.Error(), "Reference Missing") {
		t.Errorf("Response should contain missing reference error, response - %s", err.Error())
	}
}

// TestAuthorise
func TestAuthorise(t *testing.T) {
	t.Parallel()

	instance := getTestInstance()

	authRequest := &Authorise{
		Card: &Card{
			Number:      "4111111111111111",
			ExpireMonth: "08",
			ExpireYear:  "2018",
			Cvc:         "737",
			HolderName:  "John Smith",
		},
		Amount: &Amount{
			Value:    1000,
			Currency: "EUR",
		},
		Reference:       "DE-TEST-1" + randomString(10),
		MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
	}

	response, err := instance.Payment().Authorise(authRequest)

	knownError, ok := err.(apiError)

	if ok {
		t.Errorf("Response should be succesfull. Known API Error: Code - %s, Message - %s, Type - %s", knownError.ErrorCode, knownError.Message, knownError.ErrorType)
	}

	if err != nil {
		t.Errorf("Response should be succesfull, error - %s", err.Error())
	}

	if response.PspReference == "" {
		t.Errorf("Response should contain PSP Reference. Response - %s", response)
	}

	if response.ResultCode != "Authorised" {
		t.Errorf("Response resultCode should be Authorised, Response - %s", response)
	}
}
