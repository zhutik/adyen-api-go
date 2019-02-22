package adyen

import "github.com/google/go-querystring/query"

// PaymentGateway - Adyen payment transaction logic
type PaymentGateway struct {
	*Adyen
}

// authoriseType - authorise type request, @TODO: move to enums
const authoriseType = "authorise"

// directoryLookupURL - version 2 url for Directory Lookup request
const directoryLookupURL = "directory/v2"

// skipHppUrl - SkipDetails request endpoint
const skipHppURL = "skipDetails"

// authorise3DType - authorise 3DS type request, @TODO: move to enums
const authorise3DType = "authorise3d"

// authorise3DS2Type - authorise 3DS 2.0 type request, @TODO: move to enums
const authorise3DS2Type = "authorise3d2"

// AuthoriseEncrypted - Perform authorise payment in Adyen
//
// To perform recurring payment, AuthoriseEncrypted need to have contract specified and shopperReference
//
// Example:
//   &adyen.AuthoriseEncrypted{
//       Amount:           &adyen.Amount{Value: "2000", Currency: "EUR"},
//       MerchantAccount:  "merchant-account",
//       AdditionalData:   &adyen.AdditionalData{Content: r.Form.Get("adyen-encrypted-data")}, // encrypted CC data
//       ShopperReference: "unique-customer-reference",
//       Recurring:        &adyen.Recurring{Contract:adyen.RecurringPaymentRecurring}
//       Reference:        "some-merchant-reference",
//   }
//}
// adyen.Recurring{Contract:adyen.RecurringPaymentRecurring} as one of the contracts
func (a *PaymentGateway) AuthoriseEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	url := a.adyenURL(PaymentService, authoriseType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise - Perform authorise payment in Adyen
//
// Used to perform authorisation transaction without credit card data encrypted
//
// NOTE: Due to PCI compliance, it's not recommended to send credit card data to server
//
// Please use AuthoriseEncrypted instead and adyen frontend encryption library
func (a *PaymentGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	url := a.adyenURL(PaymentService, authoriseType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise3DSEncrypted - Perform authorise payment in Adyen
//
// To perform recurring payment, AuthoriseEncrypted need to have contract specified and shopperReference
//
// Example:
//   &adyen.AuthoriseEncrypted{
//       Amount:           adyen.NewAmount("EUR", 10.50),
//       MerchantAccount:  "merchant-account",
//       ThreeDS2RequestData: &adyen.ThreeDS2RequestData{
//          AuthenticationOnly: true,
//          DeviceChannel:      adyen.DeviceChannelBrowser,
//          NotificationURL:    "http://localhost:8080/3ds2/notification",
//       },
//       BrowserInfo: &adyen.BrowserInfo{
//          UserAgent:    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
//          AcceptHeader: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
//       },
//       AdditionalData:   &adyen.AdditionalData{Content: r.Form.Get("adyen-3ds2-encrypted-data")}, // encrypted CC data
//   }
//}
// adyen.Recurring{Contract:adyen.RecurringPaymentRecurring} as one of the contracts
func (a *PaymentGateway) Authorise3DSEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	url := a.adyenURL(PaymentService, authoriseType, Payment3DS2APIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise3ds2 - Perform authorise 3DS 2.0 payment in Adyen
func (a *PaymentGateway) Authorise3ds2(req *Authorise3DS2) (*AuthoriseResponse, error) {
	url := a.adyenURL(PaymentService, authorise3DS2Type, Payment3DS2APIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// DirectoryLookup - Execute directory lookup request
//
// Link - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api
func (a *PaymentGateway) DirectoryLookup(req *DirectoryLookupRequest) (*DirectoryLookupResponse, error) {

	// Calculate HMAC signature to request
	err := req.CalculateSignature(a.Adyen)
	if err != nil {
		return nil, err
	}

	url := a.createHPPUrl(directoryLookupURL)

	v, _ := query.Values(req)
	url = url + "?" + v.Encode()

	resp, err := a.executeHpp(url, req)

	if err != nil {
		return nil, err
	}

	return resp.directoryLookup()
}

// GetHPPRedirectURL - Generates link, so customer could be redirected
// to perform Hosted Payment Page payments
//
// Link - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api
func (a *PaymentGateway) GetHPPRedirectURL(req *SkipHppRequest) (string, error) {
	// Calculate HMAC signature to request
	if err := req.CalculateSignature(a.Adyen); err != nil {
		return "", err
	}

	url := a.createHPPUrl(skipHppURL)

	v, _ := query.Values(req)
	url = url + "?" + v.Encode()

	return url, nil
}

// Authorise3D - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise3D(req *Authorise3D) (*AuthoriseResponse, error) {
	url := a.adyenURL(PaymentService, authorise3DType, PaymentAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}
