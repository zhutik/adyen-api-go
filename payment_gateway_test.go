package adyen

import (
	"os"
	"strings"
	"testing"
	"time"
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
//
// In order to have test running correctly, account should be configured to have full API permissions
// Otherwise, Adyen API will return "not allowed" error. Please check https://github.com/Adyen/adyen-php-api-library/issues/20
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

	knownError, ok := err.(APIError)
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

// TestDirectoryLookUpMissingData - DirectoryLookUp Request failing due to missing data
func TestDirectoryLookUpMissingData(t *testing.T) {
	t.Parallel()

	instance := getTestInstance()

	timeIn := time.Now().Local().Add(time.Minute * time.Duration(60))

	directoryRequest := &DirectoryLookupRequest{
		CurrencyCode:      "EUR",
		PaymentAmount:     1000,
		MerchantReference: "DE-TEST-1" + randomString(10),
		SessionsValidity:  timeIn.Format(time.RFC3339),
		CountryCode:       "NL",
	}

	_, err := instance.Payment().DirectoryLookup(directoryRequest)
	if err == nil {
		t.Error("Request should fail due to missing request data")
	}

	if err.Error() != "merchantID, skinCode and HMAC hash need to be specified" {
		t.Errorf("Error should indicate that request is missing configuration data, error - %s", err)
	}
}

// TestDirectoryLookUp - test directory lookup v2 integration
//
// In order to have test running correctly, Adyen Skin need to be configured and passed through environment variable
func TestDirectoryLookUp(t *testing.T) {
	t.Parallel()

	instance := getTestInstanceWithHPP()

	timeIn := time.Now().Local().Add(time.Minute * time.Duration(60))

	directoryRequest := &DirectoryLookupRequest{
		CurrencyCode:      "EUR",
		MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
		PaymentAmount:     1000,
		SkinCode:          os.Getenv("ADYEN_SKINCODE"),
		MerchantReference: "DE-TEST-1" + randomString(10),
		SessionsValidity:  timeIn.Format(time.RFC3339),
		CountryCode:       "NL",
	}

	response, err := instance.Payment().DirectoryLookup(directoryRequest)

	if err != nil {
		t.Errorf("DirectoryLookup response should be successful, error - %s", err)
	}

	if len(response.PaymentMethods) == 0 {
		t.Errorf("DirectoryLookup response should contain at least one payment method available, response - %s", response)
	}
}
