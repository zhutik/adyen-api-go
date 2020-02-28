package adyen

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-querystring/query"
)

func TestSignatureCalculateSignature(t *testing.T) {
	t.Parallel()

	instance := getTestInstanceWithHPP()

	req := DirectoryLookupRequest{
		CurrencyCode:      "EUR",
		MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
		ShipBeforeDate:    "2015-11-31T13:42:40+1:00",
		PaymentAmount:     1000,
		SkinCode:          os.Getenv("ADYEN_SKINCODE"),
		MerchantReference: "DE-100100GMWJGS",
		SessionsValidity:  "2015-11-29T13:42:40+1:00",
	}

	err := req.CalculateSignature(instance)
	if err != nil {
		t.Fatal(err)
	}

	v, _ := query.Values(req)

	// there is no automated way to verify full URL, cause adyen webpage require authenication first
	// to debug request signature, print URL params, login to https://ca-test.adyen.com and follow full link below
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	if _, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		t.Fatal(err)
	}
}

func TestSignatureCalculateSignatureForSkipHppRequest(t *testing.T) {
	t.Parallel()

	instance := getTestInstanceWithHPP()

	req := SkipHppRequest{
		MerchantReference: "DE-100100GMWJGS",
		PaymentAmount:     1000,
		CurrencyCode:      instance.Currency,
		ShipBeforeDate:    "2015-11-31T13:42:40+1:00",
		SkinCode:          os.Getenv("ADYEN_SKINCODE"),
		MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
		ShopperLocale:     "en_GB",
		SessionsValidity:  "2015-11-29T13:42:40+1:00",
		CountryCode:       "NL",
		BrandCode:         "ideal",
	}

	err := req.CalculateSignature(instance)
	if err != nil {
		t.Fatal(err)
	}

	v, _ := query.Values(req)

	// there is no automated way to verify full URL, cause adyen webpage require authenication first
	// to debug request signature, print URL params, login to https://ca-test.adyen.com and follow full link below
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	if _, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		t.Fatal(err)
	}
}

func TestSignatureNotification(t *testing.T) {
	t.Parallel()

	//ref: https://github.com/Adyen/adyen-ruby-api-library/blob/53d9a03ab09d58927ec34e65d3d2acc1c5dc1ea7/spec/utils/hmac_validator_spec.rb
	itemDataJSON := `
{
	"additionalData": {
		"authCode": "1234",
	    "cardSummary": "7777"
	},
	"amount": {
		"currency": "EUR",
		"value": 1130
	},
	"eventCode": "AUTHORISATION",
	"eventDate": "2020-01-01T10:00:00+05:00",
	"merchantAccountCode": "TestMerchant",
	"merchantReference": "TestPayment-1407325143704",
	"operations": ["CANCEL", "CAPTURE", "REFUND"],
	"paymentMethod": "visa",
	"pspReference": "7914073381342284",
	"reason": "1234:7777:12\/2012",
	"success": "true"
}
`

	validSignature := "coqCmt/IZ4E3CzPvMY8zTjQVL5hYJUiBRg8UU+iCWo0="
	cases := []struct {
		name               string
		signatureInMessage string
		hmacKey            string
		exp                bool
		expErr             bool
	}{
		{
			name:               "no sig",
			signatureInMessage: "",
			hmacKey:            "",
			exp:                false,
			expErr:             true,
		},
		{
			name:               "sig, no key",
			signatureInMessage: validSignature,
			hmacKey:            "",
			exp:                false,
			expErr:             true,
		},
		{
			name:               "sig, wrong key",
			signatureInMessage: validSignature,
			hmacKey:            "DFB1EB5485895CFA84146406857104ABB4CBCABDC8AAF103A624C8F6A3EAAB00",
			exp:                false,
			expErr:             false,
		},
		{
			name:               "sig, correct key",
			signatureInMessage: validSignature,
			hmacKey:            "44782DEF547AAA06C910C43932B1EB0C71FC68D9D0C057550C48EC2ACF6BA056",
			exp:                true,
			expErr:             false,
		},
		{
			name:               "sig, malformed key",
			signatureInMessage: validSignature,
			hmacKey:            "invalid_hmac",
			exp:                false,
			expErr:             true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var itemData NotificationRequestItemData
			err := json.Unmarshal([]byte(itemDataJSON), &itemData)
			if err != nil {
				t.Fatalf("unmarshal error: %v", err)
			}
			if c.hmacKey != "" {
				itemData.AdditionalData.HmacSignature = c.signatureInMessage
			}
			config := NewWithHMAC(Testing, "username", "fake_password", c.hmacKey)
			res, err := itemData.ValidateSignature(config)
			if (err != nil) != c.expErr {
				t.Fatalf("expected error?: %t, actual error: %v", c.expErr, err)
			}
			if res != c.exp {
				t.Fatalf("expected result: %t, actual result: %t", c.exp, res)
			}
		})
	}
}
