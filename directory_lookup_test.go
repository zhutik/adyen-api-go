package adyen

import (
	"github.com/google/go-querystring/query"
	"net/http"
	"os"
	"testing"
)

func TestDirectoryLookupRequest_CalculateSignature(t *testing.T) {
	t.Parallel()

	req := DirectoryLookupRequest{
		CurrencyCode:      "EUR",
		MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
		PaymentAmount:     1000,
		SkinCode:          "sgOgVcKV",
		MerchantReference: "DE-100100GMWJGS",
		SessionsValidity:  "2015-11-29T13:42:40+1:00",
	}

	req.CalculateSignature()

	v, _ := query.Values(req)
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}
}
