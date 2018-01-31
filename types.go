package adyen

import "time"

/*********
* Amount *
*********/

// Amount value/currency representation
type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

/*********
* Cancel *
*********/

// Cancel structure for Cancel request
type Cancel struct {
	Reference         string `json:"reference"`
	MerchantAccount   string `json:"merchantAccount"`
	OriginalReference string `json:"originalReference"`
}

// CancelResponse is a response structure for Adyen cancellation
type CancelResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}

// CancelOrRefundResponse is a response structure for Adyen cancelOrRefund
type CancelOrRefundResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}

/**********
* Capture *
**********/

// Capture structure for Capture request
type Capture struct {
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	MerchantAccount    string  `json:"merchantAccount"`
	OriginalReference  string  `json:"originalReference"`
}

// CaptureResponse is a response structure for Adyen capture
type CaptureResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
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

/*******************
* Directory lookup *
*******************/

// DirectoryLookupRequest - get list of available payment methods based on skin, country and order details
//
// Description - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#directoryrequest
// CountryCode could be used to test local payment methods, if client's IP is from different country
type DirectoryLookupRequest struct {
	CurrencyCode      string `url:"currencyCode"`
	MerchantAccount   string `url:"merchantAccount"`
	PaymentAmount     int    `url:"paymentAmount"`
	SkinCode          string `url:"skinCode"`
	MerchantReference string `url:"merchantReference"`
	SessionsValidity  string `url:"sessionValidity"`
	MerchantSig       string `url:"merchantSig"`
	CountryCode       string `url:"countryCode"`
	ShipBeforeDate    string `url:"shipBeforeDate"`
}

// DirectoryLookupResponse - api response for DirectoryLookupRequest
//
// Description - https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#directoryresponse
type DirectoryLookupResponse struct {
	PaymentMethods []PaymentMethod `json:"paymentMethods"`
}

// PaymentMethod - structure for single payment method in directory look up response
//
// Part of DirectoryLookupResponse
type PaymentMethod struct {
	BrandCode string   `json:"brandCode"`
	Name      string   `json:"name"`
	Logos     logos    `json:"logos"`
	Issuers   []issuer `json:"issuers"`
}

// logos - payment method logos
//
// Part of DirectoryLookupResponse
type logos struct {
	Normal string `json:"normal"`
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
}

// issuer - bank issuer type
//
// Part of DirectoryLookupResponse
type issuer struct {
	IssuerID string `json:"issuerId"`
	Name     string `json:"name"`
}

/***************
* Notification *
***************/

// NotificationRequest contains environment specification and list of notifications to process
//
// Link - https://docs.adyen.com/developers/api-reference/notifications-api#notificationrequest
type NotificationRequest struct {
	Live              stringBool                `json:"live"`
	NotificationItems []NotificationRequestItem `json:"notificationItems"`
}

// NotificationRequestItem contains notification details
//
// Depending on notification type, different fields can be populated and send from Adyen
//
// Link - https://docs.adyen.com/developers/api-reference/notifications-api#notificationrequestitem
type NotificationRequestItem struct {
	NotificationRequestItem struct {
		AdditionalData struct {
			ShopperReference         string `json:"shopperReference,omitempty"`
			ShopperEmail             string `json:"shopperEmail,omitempty"`
			AuthCode                 string `json:"authCode,omitempty"`
			CardSummary              string `json:"cardSummary,omitempty"`
			ExpiryDate               string `json:"expiryDate,omitempty"`
			AuthorisedAmountValue    string `json:"authorisedAmountValue,omitempty"`
			AuthorisedAmountCurrency string `json:"authorisedAmountCurrency,omitempty"`
		} `json:"additionalData,omitempty"`
		Amount              Amount     `json:"amount"`
		PspReference        string     `json:"pspReference"`
		EventCode           string     `json:"eventCode"`
		EventDate           time.Time  `json:"eventDate"` // Event date in time.RFC3339 format
		MerchantAccountCode string     `json:"merchantAccountCode"`
		Operations          []string   `json:"operations"`
		MerchantReference   string     `json:"merchantReference"`
		OriginalReference   string     `json:"originalReference,omitempty"`
		PaymentMethod       string     `json:"paymentMethod"`
		Reason              string     `json:"reason,omitempty"`
		Success             stringBool `json:"success"`
	} `json:"NotificationRequestItem"`
}

/**********
* Payment *
**********/

// One-click functionality gives the shopper the option to store their payment details with the merchant, within the Adyen environment.
//
// In this type of transaction, the shopper needs to enter the CVC code for the transaction to get through.
//
// Link: https://docs.adyen.com/developers/api-reference/payments-api#recurring
const (
	RecurringPaymentOneClick              = "ONECLICK"
	RecurringPaymentRecurring             = "RECURRING"
	ShopperInteractionContAuth            = "ContAuth"
	SelectRecurringDetailReferenceLatests = "LATEST"
)

// AuthoriseEncrypted structure for Authorisation request (with encrypted card information)
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest
type AuthoriseEncrypted struct {
	AdditionalData                   *AdditionalData `json:"additionalData,omitempty"`
	Amount                           *Amount         `json:"amount"`
	Reference                        string          `json:"reference"`
	MerchantAccount                  string          `json:"merchantAccount"`
	ShopperReference                 string          `json:"shopperReference,omitempty"` // Mandatory for recurring payment
	Recurring                        *Recurring      `json:"recurring,omitempty"`
	ShopperInteraction               string          `json:"shopperInteraction,omitempty"`
	SelectedRecurringDetailReference string          `json:"selectedRecurringDetailReference,omitempty"`
	BrowserInfo                      *BrowserInfo    `json:"browserInfo,omitempty"` // Required for a 3DS process
}

// Authorise structure for Authorisation request (card is not encrypted)
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest
type Authorise struct {
	Card                             *Card        `json:"card,omitempty"`
	Amount                           *Amount      `json:"amount"`
	Reference                        string       `json:"reference"`
	MerchantAccount                  string       `json:"merchantAccount"`
	ShopperReference                 string       `json:"shopperReference,omitempty"` // Mandatory for recurring payment
	Recurring                        *Recurring   `json:"recurring,omitempty"`
	ShopperInteraction               string       `json:"shopperInteraction,omitempty"`
	SelectedRecurringDetailReference string       `json:"selectedRecurringDetailReference,omitempty"`
	BrowserInfo                      *BrowserInfo `json:"browserInfo,omitempty"` // Required for a 3DS process
}

// AuthoriseResponse is a response structure for Adyen
//
// Link - https://docs.adyen.com/developers/api-reference/payments-api#paymentresult
type AuthoriseResponse struct {
	PspReference   string          `json:"pspReference"`
	ResultCode     string          `json:"resultCode"`
	AuthCode       string          `json:"authCode"`
	RefusalReason  string          `json:"refusalReason"`
	IssuerURL      string          `json:"issuerUrl"`
	MD             string          `json:"md"`
	PaRequest      string          `json:"paRequest"`
	AdditionalData *AdditionalData `json:"additionalData,omitempty"`
}

// AdditionalData stores encrypted information about customer's credit card
type AdditionalData struct {
	Content   string `json:"card.encrypted.json,omitempty"`
	AliasType string `json:"aliasType,omitempty"`
	Alias     string `json:"alias,omitempty"`
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

/*************
* Payment 3D *
*************/

// Authorise3D structure for Authorisation request (card is not encrypted)
//
// https://docs.adyen.com/developers/api-reference/payments-api#paymentrequest3d
type Authorise3D struct {
	MD              string       `json:"md"`
	MerchantAccount string       `json:"merchantAccount"`
	BrowserInfo     *BrowserInfo `json:"browserInfo"`
	PaResponse      string       `json:"paResponse"`
	ShopperIP       string       `json:"shopperIP,omitempty"`
}

/************
* Recurring *
************/

// RecurringDetailsRequest structure to list all recurring payment associated to a shopperReference
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#recurringdetailsrequest
type RecurringDetailsRequest struct {
	MerchantAccount  string `json:"merchantAccount"`
	ShopperReference string `json:"shopperReference,omitempty"`
	// Not mandatory
	Recurring *Recurring `json:"recurring,omitempty"`
}

// RecurringDetailsResult structure to hold the RecurringDetails
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#recurringdetailsresult
type RecurringDetailsResult struct {
	CreationDate string `json:"creationDate"`
	Details      []struct {
		RecurringDetail RecurringDetail `json:"RecurringDetail"`
	} `json:"details"`
	InvalidOneclickContracts string `json:"invalidOneclickContracts"`
	ShopperReference         string `json:"shopperReference"`
}

// RecurringDetail structure to hold information associated to a recurring payment
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#recurringdetail
type RecurringDetail struct {
	Acquirer        string `json:"acquirer"`
	AcquirerAccount string `json:"acquirerAccount"`
	AdditionalData  struct {
		CardBin string `json:"cardBin"`
	} `json:"additionalData"`
	Alias                    string   `json:"alias"`
	AliasType                string   `json:"aliasType"`
	Card                     Card     `json:"card,omitempty"`
	ContractTypes            []string `json:"contractTypes"`
	CreationDate             string   `json:"creationDate"`
	FirstPspReference        string   `json:"firstPspReference"`
	PaymentMethodVariant     string   `json:"paymentMethodVariant"`
	RecurringDetailReference string   `json:"recurringDetailReference"`
	Variant                  string   `json:"variant"`
}

// RecurringDisableRequest structure to hold information regarding disable recurring request
//
// If `RecurringDetailReference` is specified, specific payment ID will be disabled
// otherwise all customer saved payment methods will be disabled
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#disablerequest
type RecurringDisableRequest struct {
	MerchantAccount  string `json:"merchantAccount"`
	ShopperReference string `json:"shopperReference"`
	// Type of a contract ONECLICK, RECURRING, PAYOUT or combination of them
	Contract string `json:"contract,omitempty"`
	// ID of a customer saved payment method, all will be disabled if none is specified
	RecurringDetailReference string `json:"recurringDetailReference,omitempty"`
}

// RecurringDisableResponse structure to hold response for disable recurring request
//
// Link - https://docs.adyen.com/developers/api-reference/recurring-api#disableresult
type RecurringDisableResponse struct {
	Response string `json:"response"`
}

/*********
* Refund *
*********/

// Refund structure for refund request
type Refund struct {
	ModificationAmount *Amount `json:"modificationAmount"`
	Reference          string  `json:"reference"`
	MerchantAccount    string  `json:"merchantAccount"`
	OriginalReference  string  `json:"originalReference"`
}

// RefundResponse is a response structure for Adyen refund request
type RefundResponse struct {
	PspReference string `json:"pspReference"`
	Response     string `json:"response"`
}

/***********
* Skip HPP *
***********/

// SkipHppRequest contains data that would be used to create Adyen HPP redirect URL
//
// Link: https://docs.adyen.com/developers/ecommerce-integration/local-payment-methods
//
// Request description: https://docs.adyen.com/developers/api-reference/hosted-payment-pages-api#skipdetailsrequest
type SkipHppRequest struct {
	MerchantReference string `url:"merchantReference"`
	PaymentAmount     int    `url:"paymentAmount"`
	CurrencyCode      string `url:"currencyCode"`
	ShipBeforeDate    string `url:"shipBeforeDate"`
	SkinCode          string `url:"skinCode"`
	MerchantAccount   string `url:"merchantAccount"`
	ShopperLocale     string `url:"shopperLocale"`
	SessionsValidity  string `url:"sessionValidity"`
	MerchantSig       string `url:"merchantSig"`
	CountryCode       string `url:"countryCode"`
	BrandCode         string `url:"brandCode"`
	IssuerID          string `url:"issuerId"`
}
