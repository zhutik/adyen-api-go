package adyen

const (
	// Male to indicate "male" gender
	Male = "MALE"
	// Female to indicate "female" gender
	Female = "FEMALE"
	// Unknown to indicate "unknown" gender
	Unknown = "UNKNOWN"
)

/**********
* Address *
**********/

// Address - base address type for customer billing and delivery addresses
//
// Link - https://docs.adyen.com/developers/api-reference/common-api#address
type Address struct {
	City              string `json:"city"`
	Country           string `json:"country"`
	HouseNumberOrName string `json:"houseNumberOrName"`
	PostalCode        string `json:"postalCode,omitempty"`
	StateOrProvince   string `json:"stateOrProvince,omitempty"`
	Street            string `json:"street"`
}

/*******
* Card *
*******/

// Card structure representation
type Card struct {
	Number      string `json:"number"`
	ExpireMonth string `json:"expiryMonth"`
	ExpireYear  string `json:"expiryYear"`
	Cvc         string `json:"cvc"`
	HolderName  string `json:"holderName"`
}

/*******
* Name *
*******/

// Name - generic name structure
//
// Link - https://docs.adyen.com/developers/api-reference/common-api#name
type Name struct {
	FirstName string `json:"firstName"`
	Gender    string `json:"gender"` // Should be ENUM (Male, Female, Unknown) from a constants
	Infix     string `json:"infix,omitempty"`
	LastName  string `json:"lastName"`
}
