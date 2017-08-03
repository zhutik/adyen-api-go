package adyen

// Authorise3D structure for Authorisation request (card is not encrypted)
type Authorise3D struct {
	MD string `json:"md"`
	MerchantAccount string     `json:"merchantAccount"`
	BrowserInfo *BrowserInfo `json:"browserInfo"`
	PaResponse string `json:"paResponse"`
	ShopperIP string `json:"shopperIP,omitempty"`
}
