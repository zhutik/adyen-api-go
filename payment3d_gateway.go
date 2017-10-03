package adyen

// Payment3DGateway - Adyen payment transaction logic
type Payment3DGateway struct {
	*Adyen
}

// authorise3DType - authorise type request, @TODO: move to enums
const authorise3DType = "authorise3d"

// Authorise3D - Perform authorise payment in Adyen
func (a *PaymentGateway) Authorise3D(req *Authorise3D) (*AuthoriseResponse, error) {
	resp, err := a.execute(authorise3DType, PaymentService, req)

	if err != nil {
		return nil, err
	}

	return resp.authorize()
}
