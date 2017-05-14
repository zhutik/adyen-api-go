package adyen

import "encoding/json"

// PaymentGateway - Adyen payment transaction logic
type PaymentGateway struct {
	*Adyen
}

// AuthoriseType - authorise type request, @TODO: move to enums
const AuthoriseType = "authorise"

// AuthoriseEncrypted - Perform authorise payment in Adyen
func (a *PaymentGateway) AuthoriseEncrypted(req *AuthoriseEncrypted) (*AuthoriseResponse, error) {
	resp, err := a.execute(AuthoriseType, req)

	if err != nil {
		return nil, err
	}

	var val AuthoriseResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}

// Authorise - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise(req *Authorise) (*AuthoriseResponse, error) {
	resp, err := a.execute(AuthoriseType, req)

	if err != nil {
		return nil, err
	}

	var val AuthoriseResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}
