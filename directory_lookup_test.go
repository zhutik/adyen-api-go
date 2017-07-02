package adyen

import (
	"github.com/google/go-querystring/query"
	"net/http"
	"os"
	"testing"
)

func initAdyenWithHPP() *Adyen {
	instance := NewWithHPP(
		Testing,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_CLIENT_TOKEN"),
		os.Getenv("ADYEN_ACCOUNT"),
		os.Getenv("ADYEN_HMAC"),
		os.Getenv("ADYEN_SKINCODE"),
		os.Getenv("ADYEN_SHOPPER_LOCALE"),
	)

	return instance
}

func TestDirectoryLookupRequest_CalculateSignature(t *testing.T) {
	t.Parallel()

	instance := initAdyenWithHPP()

	req := DirectoryLookupRequest{
		CurrencyCode:      "EUR",
		MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
		PaymentAmount:     1000,
		SkinCode:          "sgOgVcKV",
		MerchantReference: "DE-100100GMWJGS",
		SessionsValidity:  "2015-11-29T13:42:40+1:00",
	}

	err := req.CalculateSignature(instance)

	if err != nil {
		t.Error(err)
	}

	v, _ := query.Values(req)
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	_, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}
}
