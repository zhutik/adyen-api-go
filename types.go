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

/***************
* Account Info *
***************/

// AccountInfo constants
const (
	AccountInfoNotApplible     = "notApplicable"
	AccountInfoThisTransaction = "thisTransaction"
	AccountInfoLessThan30Days  = "lessThan30Days"
	AccountInfoFrom30To60Days  = "from30To60Days"
	AccountInfoMoreThan60Days  = "moreThan60Days"
)

// AccountInfo - Shopper account information for 3D Secure 2.0
//
// Description - https://docs.adyen.com/developers/risk-management/3d-secure-2-0/server-integration/api-reference-3d-secure-2-0#accountinfo
type AccountInfo struct {
	AccountAgeIndicator           string `json:"accountAgeIndicator,omitempty"` // one of AccountInfo constants
	AccountChangeDate             string `json:"accountChangeDate,omitempty"`
	AccountChangeIndicator        string `json:"accountChangeIndicator,omitempty"` // one of AccountInfo constants
	AccountCreationDate           string `json:"accountCreationDate,omitempty"`
	PasswordChangeDate            string `json:"passwordChangeDate,omitempty"`
	PasswordChangeDateIndicator   string `json:"passwordChangeDateIndicator,omitempty"` // one of AccountInfo constants
	PurchasesLast6Months          int    `json:"purchasesLast6Months,omitempty"`
	AddCardAttemptsDay            int    `json:"addCardAttemptsDay,omitempty"`
	PastTransactionsDay           int    `json:"pastTransactionsDay,omitempty"`
	PastTransactionsYear          int    `json:"pastTransactionsYear,omitempty"`
	PaymentAccountAge             string `json:"paymentAccountAge,omitempty"`
	PaymentAccountIndicator       string `json:"paymentAccountIndicator,omitempty"` // one of AccountInfo constants
	DeliveryAddressUsageDate      string `json:"deliveryAddressUsageDate,omitempty"`
	DeliveryAddressUsageIndicator string `json:"deliveryAddressUsageIndicator,omitempty"`
	SuspiciousAccActivity         bool   `json:"suspiciousAccActivity,omitempty"`
	HomePhone                     string `json:"homePhone,omitempty"`
	MobilePhone                   string `json:"mobilePhone,omitempty"`
	WorkPhone                     string `json:"workPhone,omitempty"`
}

/**************
* BrowserInfo *
**************/

// BrowserInfo hold information on the user browser
type BrowserInfo struct {
	AcceptHeader   string `json:"acceptHeader"`
	UserAgent      string `json:"userAgent"`
	Language       string `json:"language,omitempty"`
	ColorDepth     int    `json:"colorDepth,omitempty"`
	JavaEnabled    bool   `json:"javaEnabled,omitempty"`
	ScreenHeight   int    `json:"screenHeight,omitempty"`
	ScreenWidth    int    `json:"screenWidth,omitempty"`
	TimeZoneOffset int    `json:"timeZoneOffset,omitempty"`
}
