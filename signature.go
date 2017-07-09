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
	if len(adyen.Credentials.MerchantID) == 0 ||
		len(adyen.Credentials.HppSettings.SkinCode) == 0 ||
		len(adyen.Credentials.HppSettings.Hmac) == 0 {
		return errors.New("merchantID, skinCode and HMAC hash need to be specified")
	}

	// @todo write sorting for this map
	keys := "countryCode"
	values := replaceSpecialChars(r.CountryCode)

	keys += ":" + "currencyCode"
	values += ":" + replaceSpecialChars(r.CurrencyCode)

	keys += ":" + "merchantAccount"
	values += ":" + replaceSpecialChars(adyen.Credentials.MerchantID)

	keys += ":" + "merchantReference"
	values += ":" + replaceSpecialChars(r.MerchantReference)

	keys += ":" + "paymentAmount"
	values += ":" + replaceSpecialChars(strconv.Itoa(r.PaymentAmount))

	keys += ":" + "sessionValidity"
	values += ":" + replaceSpecialChars(r.SessionsValidity)

	keys += ":" + "shipBeforeDate"
	values += ":" + replaceSpecialChars(r.ShipBeforeDate)

	keys += ":" + "skinCode"
	values += ":" + replaceSpecialChars(adyen.Credentials.HppSettings.SkinCode)

	fullString := keys + ":" + values

	src, err := hex.DecodeString(adyen.Credentials.HppSettings.Hmac)

	if err != nil {
		return err
	}

	mac := hmac.New(sha256.New, src)
	mac.Write([]byte(fullString))

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return nil
}
