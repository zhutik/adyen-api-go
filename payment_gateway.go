package adyen

// PaymentGateway - Adyen payment transaction logic
type PaymentGateway struct {
	*Adyen
}

// authoriseType - authorise type request, @TODO: move to enums
const authoriseType = "authorise"

// directoryLookupURL - version 2 url for Directory Lookup request
const directoryLookupURL = "directory/v2"

// AuthoriseEncrypted - Perform authorise payment in Adyen
func (a *PaymentGateway) AuthoriseEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	resp, err := a.execute(authoriseType, PaymentService, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	resp, err := a.execute(authoriseType, PaymentService, req)

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
