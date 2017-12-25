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

// AuthoriseEncrypted - Perform authorise payment in Adyen
func (a *PaymentGateway) AuthoriseEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	resp, err := a.execute(PaymentService, authoriseType, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	resp, err := a.execute(PaymentService, authoriseType, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// DirectoryLookup - Execute directory lookup request
func (a *PaymentGateway) DirectoryLookup(req *DirectoryLookupRequest) (*DirectoryLookupResponse, error) {

	// Calculate HMAC signature to request
	err := req.CalculateSignature(a.Adyen)
	if err != nil {
		return nil, err
	}

	resp, err := a.executeHpp(directoryLookupURL, req)

	if err != nil {
		return nil, err
	}

	return resp.directoryLookup()
}

// GetHPPRedirectURL - Generates link, so customer could be redirected
// to perform Hosted Payment Page payments
func (a *PaymentGateway) GetHPPRedirectURL(req *SkipHppRequest) (string, error) {
	// Calculate HMAC signature to request
	err := req.CalculateSignature(a.Adyen)
	if err != nil {
		return "", err
	}

	url := a.createHPPUrl(skipHppURL)

	v, _ := query.Values(req)
	url = url + "?" + v.Encode()

	return url, nil
}
