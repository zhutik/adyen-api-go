package adyen

/* One-click functionality gives the shopper the option to store their payment details with the merchant, within the Adyen environment. In this type of transaction, the shopper needs to enter the CVC code for the transaction to get through.
Same than for ONECLICK except than the shopper doesn't need to enter the CVC */
const (
	ONECLICK  = "ONECLICK"
	RECURRING = "RECURRING"
)

// AuthoriseEncrypted structure for Authorisation request (with encrypted card information)
type AuthoriseEncrypted struct {
	AdditionalData  *AdditionalData `json:"additionalData"`
	Amount          *Amount         `json:"amount"`
	Reference       string          `json:"reference"`
	MerchantAccount string          `json:"merchantAccount"`
	// Mandatory for recurring payment
	ShopperReference string     `json:"shopperReference,omitempty"`
	Recurring        *Recurring `json:"recurring,omitempty"`
	// Required for a 3DS process
	BrowserInfo *BrowserInfo `json:"browserInfo,omitempty"`
}

// Authorise structure for Authorisation request (card is not encrypted)
type Authorise struct {
	Card            *Card   `json:"card"`
	Amount          *Amount `json:"amount"`
	Reference       string  `json:"reference"`
	MerchantAccount string  `json:"merchantAccount"`
	// Mandatory for recurring payment
	ShopperReference string     `json:"shopperReference,omitempty"`
	Recurring        *Recurring `json:"recurring,omitempty"`
	// Required for a 3DS process
	BrowserInfo *BrowserInfo `json:"browserInfo,omitempty"`
}

// AuthoriseResponse is a response structure for Adyen
type AuthoriseResponse struct {
	PspReference  string `json:"pspReference"`
	ResultCode    string `json:"resultCode"`
	AuthCode      string `json:"authCode"`
	RefusalReason string `json:"refusalReason"`
	IssuerURL     string `json:"issuerUrl"`
	MD            string `json:"md"`
	PaRequest     string `json:"paRequest"`
}

// AdditionalData stores encrypted information about customer's credit card
type AdditionalData struct {
	Content string `json:"card.encrypted.json"`
}

// BrowserInfo hold information on the user browser
type BrowserInfo struct {
	AcceptHeader string `json:"acceptHeader"`
	UserAgent    string `json:"userAgent"`
}

// Recurring hold the behavior for a future payment : could be ONECLICK or RECURRING
type Recurring struct {
	Contract string `json:"contract"`
}
