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
