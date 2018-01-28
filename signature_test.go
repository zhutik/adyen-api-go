package adyen

import (
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
		t.Error(err)
	}

	v, _ := query.Values(req)

	// there is no automated way to verify full URL, cause adyen webpage require authenication first
	// to debug request signature, print URL params, login to https://ca-test.adyen.com and follow full link below
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	_, err = http.NewRequest("GET", url, nil)

	if err != nil {
		t.Error(err)
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
		t.Error(err)
	}

	v, _ := query.Values(req)

	// there is no automated way to verify full URL, cause adyen webpage require authenication first
	// to debug request signature, print URL params, login to https://ca-test.adyen.com and follow full link below
	url := "https://ca-test.adyen.com/ca/ca/skin/checkhmac.shtml" + "?" + v.Encode()

	_, err = http.NewRequest("GET", url, nil)

	if err != nil {
		t.Error(err)
	}
}
