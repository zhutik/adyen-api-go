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

// ReplaceSpecialChars replace special characters according to Adyen documentation
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
func ReplaceSpecialChars(value string) string {
	temp := strings.Replace(value, "\\", "\\\\", -1)
	temp = strings.Replace(temp, ":", "\\:", -1)

	return temp
}

// CalculateSignature calculate HMAC signature for request
//
// Link: https://docs.adyen.com/developers/payments/accepting-payments/hmac-signature-calculation
// @todo: refactor this method
func (r *DirectoryLookupRequest) CalculateSignature(adyen *Adyen) error {
	if len(adyen.Credentials.merchantID) == 0 ||
		len(adyen.Credentials.hppSettings.skinCode) == 0 ||
		len(adyen.Credentials.hppSettings.hmac) == 0 {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	// @todo write sorting for this map
	keys := "countryCode"
	values := ReplaceSpecialChars(r.CountryCode)

	keys += ":" + "currencyCode"
	values += ":" + ReplaceSpecialChars(r.CurrencyCode)

	keys += ":" + "merchantAccount"
	values += ":" + ReplaceSpecialChars(adyen.Credentials.merchantID)

	keys += ":" + "merchantReference"
	values += ":" + ReplaceSpecialChars(r.MerchantReference)

	keys += ":" + "paymentAmount"
	values += ":" + ReplaceSpecialChars(strconv.Itoa(r.PaymentAmount))

	keys += ":" + "sessionValidity"
	values += ":" + ReplaceSpecialChars(r.SessionsValidity)

	keys += ":" + "shipBeforeDate"
	values += ":" + ReplaceSpecialChars(r.ShipBeforeDate)

	keys += ":" + "skinCode"
	values += ":" + ReplaceSpecialChars(adyen.Credentials.hppSettings.skinCode)

	fullString := keys + ":" + values

	src, err := hex.DecodeString(adyen.Credentials.hppSettings.hmac)

	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	mac.Write([]byte(fullString))

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return nil
}
