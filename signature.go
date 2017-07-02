package adyen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"encoding/base64"
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
func (r *DirectoryLookupRequest) CalculateSignature() {
	var pairs = make(map[string]string)

	// this map should be sorted by keys
	// @todo write sorting for this map
	pairs["countryCode"] = ""
	pairs["currencyCode"] = r.CurrencyCode
	pairs["merchantAccount"] = r.MerchantAccount
	pairs["merchantReference"] = r.MerchantReference
	pairs["paymentAmount"] = strconv.Itoa(r.PaymentAmount)
	pairs["sessionValidity"] = r.SessionsValidity
	pairs["shipBeforeDate"] = "2018-07-30"
	pairs["shopperLocale"] = "en_GB"
	pairs["skinCode"] = r.SkinCode

	r.ShipBeforeDate = pairs["shipBeforeDate"]

	keys := "countryCode"
	values := ReplaceSpecialChars(pairs["countryCode"])

	keys += ":" + "currencyCode"
	values += ":" + ReplaceSpecialChars(pairs["currencyCode"])

	keys += ":" + "merchantAccount"
	values += ":" + ReplaceSpecialChars(pairs["merchantAccount"])

	keys += ":" + "merchantReference"
	values += ":" + ReplaceSpecialChars(pairs["merchantReference"])

	keys += ":" + "paymentAmount"
	values += ":" + ReplaceSpecialChars(pairs["paymentAmount"])

	keys += ":" + "sessionValidity"
	values += ":" + ReplaceSpecialChars(pairs["sessionValidity"])

	keys += ":" + "shipBeforeDate"
	values += ":" + ReplaceSpecialChars(pairs["shipBeforeDate"])

	keys += ":" + "skinCode"
	values += ":" + ReplaceSpecialChars(pairs["skinCode"])

	fullString := keys + ":" + values

	src, _ := hex.DecodeString("46AE75207D01236D1DAE55AF004F09CD18EDC303FC7F459038B01CED70D8A595")

	mac := hmac.New(sha256.New, src)
	mac.Write([]byte(fullString))

	r.MerchantSig = base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
