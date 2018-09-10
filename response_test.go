package adyen

import (
	"testing"
)

// TestResponseErrorResponseStatus - response is valid JSON, but status is > 299, error should be returned
func TestResponseErrorResponseStatus(t *testing.T) {
	t.Parallel()

	responseJSON := `
{
	"errorType" : "authorise",
	"errorCode" : "501",
	"message"   : "sample error",
	"status"    : 501
}
`
	providerResponse, err := createTestResponse(responseJSON, "OK 200", 200)

	if err != nil {
		t.Fatal(err)
	}

	err = providerResponse.handleHTTPError()
	if err == nil {
		t.Fatal("Response should raise an error - got nil")
	}

	errorContent := "[authorise][501]: (501) sample error"
	if err.Error() != errorContent {
		t.Fatal("Expected error message ", errorContent, " got ", err.Error())
	}
}

// TestResponseNotValidJson - response is an empty script, error should be returned
func TestResponseNotValidJson(t *testing.T) {
	t.Parallel()

	providerResponse, err := createTestResponse("", "503", 503)

	err = providerResponse.handleHTTPError()
	if err == nil {
		t.Fatal("Response should raise an error - got nil")
	}
}

func TestAuthorizeResponse(t *testing.T) {
	cases := []struct {
		name       string
		input      string
		reference  string
		resultCode string
		authCode   string
		expErr     bool
	}{
		{
			name: "authorize response",
			input: `{
				"pspReference" : "8413547924770610",
				"ResultCode" : "Authorised",
				"AuthCode" : "53187"
			}`,
			reference:  "8413547924770610",
			resultCode: "Authorised",
			authCode:   "53187",
		},
		{
			name:       "authorize returns errors",
			input:      "some error string",
			reference:  "",
			resultCode: "",
			authCode:   "",
			expErr:     true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.authorize()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.resultCode, res.ResultCode)
			equals(t, c.authCode, res.AuthCode)
		})
	}
}

func TestCaptureResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "capture response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[capture-received]"
			}`,
			reference: "8413547924770610",
			response:  "[capture-received]",
		},
		{
			name:      "capture returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.capture()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}

func TestCancelResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "cancel response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[cancel-received]"
			}`,
			reference: "8413547924770610",
			response:  "[cancel-received]",
		},
		{
			name:      "cancel returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.cancel()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}

func TestCancelOrRefundResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "cancelOrRefund response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[cancelOrRefund-received]"
			}`,
			reference: "8413547924770610",
			response:  "[cancelOrRefund-received]",
		},
		{
			name:      "cancelOrRefund returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.cancelOrRefund()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}

func TestRefundResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "refund response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[refund-received]"
			}`,
			reference: "8413547924770610",
			response:  "[refund-received]",
		},
		{
			name:      "refund returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.refund()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}

func TestListRecurringDetailsResponse(t *testing.T) {
	cases := []struct {
		name                     string
		input                    string
		shopperReference         string
		recurringDetailReference string
		cartHolderName           string
		expErr                   bool
	}{
		{
			name: "listRecurringDetails response",
			input: `{
				"creationDate": "2018-05-23T15:25:40+02:00",
				"details": [
					{
						"RecurringDetail": {
							"additionalData": {
								"cardBin": "411111"
							},
							"alias": "K333136193308394",
							"aliasType": "Default",
							"card": {
								"expiryMonth": "8",
								"expiryYear": "2018",
								"holderName": "John Smith",
								"number": "1111"
							},
							"contractTypes": [
								"PAYOUT",
								"RECURRING",
								"ONECLICK"
							],
							"creationDate": "2018-08-10T10:28:43+02:00",
							"firstPspReference": "8815338897222637",
							"paymentMethodVariant": "visa",
							"recurringDetailReference": "8415336862463792",
							"variant": "visa"
						}
					}
				],
				"shopperReference": "yourShopperId_IOfW3k9G2PvXFu2j"
			}`,
			shopperReference:         "yourShopperId_IOfW3k9G2PvXFu2j",
			recurringDetailReference: "8415336862463792",
			cartHolderName:           "John Smith",
		},
		{
			name:                     "listRecurringDetails error response",
			input:                    "some error string",
			shopperReference:         "",
			recurringDetailReference: "",
			cartHolderName:           "",
			expErr:                   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.listRecurringDetails()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.shopperReference, res.ShopperReference)
			equals(t, 1, len(res.Details))
			equals(t, c.recurringDetailReference, res.Details[0].RecurringDetail.RecurringDetailReference)
			equals(t, c.cartHolderName, res.Details[0].RecurringDetail.Card.HolderName)
		})
	}
}

func TestDisableRecurringResponse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		response string
		expErr   bool
	}{
		{
			name: "disableRecurring response",
			input: `{
				"response" : "[detail-successfully-disabled]"
			}`,
			response: "[detail-successfully-disabled]",
		},
		{
			name:     "disableRecurring returns errors",
			input:    "some error string",
			response: "",
			expErr:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.disableRecurring()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.response, res.Response)
		})
	}
}

func TestAdjustAuthorisationResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "adjustAuthorisation response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[adjustAuthorisation-received]"
			}`,
			reference: "8413547924770610",
			response:  "[adjustAuthorisation-received]",
		},
		{
			name:      "adjustAuthorisation returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.adjustAuthorisation()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}

func TestTechnicalCancelResponse(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		reference string
		response  string
		expErr    bool
	}{
		{
			name: "technicalCancel response",
			input: `{
				"pspReference" : "8413547924770610",
				"response" : "[technical-cancel-received]"
			}`,
			reference: "8413547924770610",
			response:  "[technical-cancel-received]",
		},
		{
			name:      "technicalCancel returns errors",
			input:     "some error string",
			reference: "",
			response:  "",
			expErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			response, err := createTestResponse(c.input, "OK 200", 200)

			if err != nil {
				t.Fatal(err)
			}

			res, err := response.technicalCancel()

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}

				return
			}

			equals(t, c.reference, res.PspReference)
			equals(t, c.response, res.Response)
		})
	}
}
