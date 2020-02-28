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

// ValidateSignature validate HMAC signature for notification event
//
// Link: https://docs.adyen.com/development-resources/notifications/verify-hmac-signatures#verify-using-your-own-solution
func (n *NotificationRequestItemData) ValidateSignature(adyen *Adyen) (bool, error) {
	var precondition error
	providedSig := n.AdditionalData.HmacSignature
	if len(providedSig) == 0 {
		precondition = errors.New("no HMAC signature in message")
	} else if len(adyen.Credentials.Hmac) == 0 {
		precondition = errors.New("no HMAC key configured; cannot validate signature")
	}
	if precondition != nil {
		return false, precondition
	}

	valueInJavaFloatToStringStyle := strconv.FormatFloat(float64(n.Amount.Value), 'f', -1, 32)
	valueString := strings.Join([]string{
		replaceSpecialChars(n.PspReference),
		replaceSpecialChars(n.OriginalReference),
		replaceSpecialChars(n.MerchantAccountCode),
		replaceSpecialChars(n.MerchantReference),
		valueInJavaFloatToStringStyle,
		replaceSpecialChars(n.Amount.Currency),
		replaceSpecialChars(n.EventCode),
		strconv.FormatBool(bool(n.Success)),
	}, ":")

	src, err := hex.DecodeString(adyen.Credentials.Hmac)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, src)
	if _, err = mac.Write([]byte(valueString)); err != nil {
		return false, err
	}
	expectedSig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return expectedSig == providedSig, nil
}
