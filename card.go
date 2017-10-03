package adyen

// Card structure representation
type Card struct {
	Number      string `json:"number"`
	ExpireMonth string `json:"expiryMonth"`
	ExpireYear  string `json:"expiryYear"`
	Cvc         string `json:"cvc"`
	HolderName  string `json:"holderName"`
}
