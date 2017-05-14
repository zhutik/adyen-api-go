package adyen

// Card structure representation
type Card struct {
	Number      string `json:"number"`
	ExpireMonth int    `json:"expiryMonth"`
	ExpireYear  int    `json:"expiryYear"`
	Cvc         int    `json:"cvc"`
	HolderName  string `json:"holderName"`
}
