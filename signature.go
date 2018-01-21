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
// @todo: refactor this method
func (r *DirectoryLookupRequest) CalculateSignature(adyen *Adyen) error {
	if len(r.MerchantAccount) == 0 ||
		len(r.SkinCode) == 0 ||
		len(adyen.Credentials.HppSettings.Hmac) == 0 {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	var sortedKeys = map[int]string{
		0: "countryCode",
		1: "currencyCode",
		2: "merchantAccount",
		3: "merchantReference",
		4: "paymentAmount",
		5: "sessionValidity",
		6: "shipBeforeDate",
		7: "skinCode",
	}

	var sortedValues = map[int]string{
		0: replaceSpecialChars(r.CountryCode),
		1: replaceSpecialChars(r.CurrencyCode),
		2: replaceSpecialChars(r.MerchantAccount),
		3: replaceSpecialChars(r.MerchantReference),
		4: replaceSpecialChars(strconv.Itoa(r.PaymentAmount)),
		5: replaceSpecialChars(r.SessionsValidity),
		6: replaceSpecialChars(r.ShipBeforeDate),
		7: replaceSpecialChars(r.SkinCode),
	}

	keysString := ""
	valuesString := ""

	for k := 0; k <= 7; k++ {
		if k == 0 {
			keysString = sortedKeys[k]
			valuesString = sortedValues[k]
		} else {
			keysString += ":" + sortedKeys[k]
			valuesString += ":" + sortedValues[k]
		}
	}

	fullString := keysString + ":" + valuesString

	src, err := hex.DecodeString(adyen.Credentials.HppSettings.Hmac)

	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	_, err = mac.Write([]byte(fullString))

	if err != nil {
		return err
	}

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return nil
}

// CalculateSignature calculate HMAC signature for request
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
// @todo: refactor this method
func (r *SkipHppRequest) CalculateSignature(adyen *Adyen) error {
	if len(r.MerchantAccount) == 0 ||
		len(r.SkinCode) == 0 ||
		len(adyen.Credentials.HppSettings.Hmac) == 0 {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	var sortedKeys = map[int]string{
		0:  "brandCode",
		1:  "countryCode",
		2:  "currencyCode",
		3:  "issuerId",
		4:  "merchantAccount",
		5:  "merchantReference",
		6:  "paymentAmount",
		7:  "sessionValidity",
		8:  "shipBeforeDate",
		9:  "shopperLocale",
		10: "skinCode",
	}

	var sortedValues = map[int]string{
		0:  replaceSpecialChars(r.BrandCode),
		1:  replaceSpecialChars(r.CountryCode),
		2:  replaceSpecialChars(r.CurrencyCode),
		3:  replaceSpecialChars(r.IssuerID),
		4:  replaceSpecialChars(r.MerchantAccount),
		5:  replaceSpecialChars(r.MerchantReference),
		6:  replaceSpecialChars(strconv.Itoa(r.PaymentAmount)),
		7:  replaceSpecialChars(r.SessionsValidity),
		8:  replaceSpecialChars(r.ShipBeforeDate),
		9:  replaceSpecialChars(r.ShopperLocale),
		10: replaceSpecialChars(r.SkinCode),
	}

	keysString := ""
	valuesString := ""

	for k := 0; k <= 10; k++ {
		if k == 0 {
			keysString = sortedKeys[k]
			valuesString = sortedValues[k]
		} else {
			keysString += ":" + sortedKeys[k]
			valuesString += ":" + sortedValues[k]
		}
	}

	fullString := keysString + ":" + valuesString

	src, err := hex.DecodeString(adyen.Credentials.HppSettings.Hmac)

	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	_, err = mac.Write([]byte(fullString))

	if err != nil {
		return err
	}

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return nil
}
