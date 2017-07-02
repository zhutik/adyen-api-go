package adyen

// PaymentGateway - Adyen payment transaction logic
type PaymentGateway struct {
	*Adyen
}

// AuthoriseType - authorise type request, @TODO: move to enums
const AuthoriseType = "authorise"

// DirectoryLookupURL - version 2 url for Directory Lookup request
const DirectoryLookupURL = "directory/v2"

// AuthoriseEncrypted - Perform authorise payment in Adyen
func (a *PaymentGateway) AuthoriseEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	resp, err := a.execute(AuthoriseType, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// Authorise - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	resp, err := a.execute(AuthoriseType, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}

// DirectoryLookup - Execute directory lookup request
func (a *PaymentGateway) DirectoryLookup(req *DirectoryLookupRequest) (*DirectoryLookupResponse, error) {

	// Calculate HMAC signature to request
	req.CalculateSignature(a.Adyen)
	resp, err := a.executeHpp(DirectoryLookupURL, req)

	if err != nil {
		return nil, err
	}

	return resp.directoryLookup()
}
