package adyen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
)

// replaceSpecialChars replace special characters according to Adyen documentation
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
func replaceSpecialChars(value string) string {
	temp := strings.Replace(value, "\\", "\\\\", -1)
	temp = strings.Replace(temp, ":", "\\:", -1)

	return temp
}

// CalculateSignature calculate HMAC signature for request
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
func (r *DirectoryLookupRequest) CalculateSignature(adyen *Adyen) error {
	if r.MerchantAccount == "" || r.SkinCode == "" || adyen.Credentials.Hmac == "" {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	keyString := strings.Join([]string{
		"countryCode",
		"currencyCode",
		"merchantAccount",
		"merchantReference",
		"paymentAmount",
		"sessionValidity",
		"shipBeforeDate",
		"skinCode",
	}, ":")

	valueString := strings.Join([]string{
		replaceSpecialChars(r.CountryCode),
		replaceSpecialChars(r.CurrencyCode),
		replaceSpecialChars(r.MerchantAccount),
		replaceSpecialChars(r.MerchantReference),
		replaceSpecialChars(strconv.Itoa(r.PaymentAmount)),
		replaceSpecialChars(r.SessionsValidity),
		replaceSpecialChars(r.ShipBeforeDate),
		replaceSpecialChars(r.SkinCode),
	}, ":")

	fullString := keyString + ":" + valueString

	src, err := hex.DecodeString(adyen.Credentials.Hmac)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	if _, err = mac.Write([]byte(fullString)); err != nil {
		return err
	}

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return nil
}

// CalculateSignature calculate HMAC signature for request
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
func (r *SkipHppRequest) CalculateSignature(adyen *Adyen) error {
	if r.MerchantAccount == "" || r.SkinCode == "" || adyen.Credentials.Hmac == "" {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	keyString := strings.Join([]string{
		"brandCode",
		"countryCode",
		"currencyCode",
		"issuerId",
		"merchantAccount",
		"merchantReference",
		"paymentAmount",
		"sessionValidity",
		"shipBeforeDate",
		"shopperLocale",
		"skinCode",
	}, ":")

	valueString := strings.Join([]string{
		replaceSpecialChars(r.BrandCode),
		replaceSpecialChars(r.CountryCode),
		replaceSpecialChars(r.CurrencyCode),
		replaceSpecialChars(r.IssuerID),
		replaceSpecialChars(r.MerchantAccount),
		replaceSpecialChars(r.MerchantReference),
		replaceSpecialChars(strconv.Itoa(r.PaymentAmount)),
		replaceSpecialChars(r.SessionsValidity),
		replaceSpecialChars(r.ShipBeforeDate),
		replaceSpecialChars(r.ShopperLocale),
		replaceSpecialChars(r.SkinCode),
	}, ":")

	fullString := keyString + ":" + valueString

	src, err := hex.DecodeString(adyen.Credentials.Hmac)
	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	if _, err = mac.Write([]byte(fullString)); err != nil {
		return err
	}

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return nil
}
