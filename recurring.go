package adyen

// RecurringDetailsRequest structure to list all recurring payment associated to a shopperReference
type RecurringDetailsRequest struct {
	MerchantAccount  string `json:"merchantAccount"`
	ShopperReference string `json:"shopperReference,omitempty"`
	// Not mandatory
	Recurring *Recurring `json:"recurring,omitempty"`
}

// RecurringDetailsResult structure to hold the RecurringDetails
type RecurringDetailsResult struct {
	CreationDate string `json:"creationDate"`
	Details      []struct {
		RecurringDetail RecurringDetail `json:"RecurringDetail"`
	} `json:"details"`
	InvalidOneclickContracts string `json:"invalidOneclickContracts"`
	ShopperReference         string `json:"shopperReference"`
}

// RecurringDetail structure to hold information associated to a recurring payment
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
type RecurringDisableRequest struct {
	MerchantAccount          string `json:"merchantAccount"`
	ShopperReference         string `json:"shopperReference"`
	// Type of a contract ONECLICK, RECURRING, PAYOUT or combination of them
	Contract                 string `json:"contract,omitempty"`
	// ID of a customer saved payment method, all will be disabled if none is specified
	RecurringDetailReference string `json:"recurringDetailReference,omitempty"`
}

// RecurringDisableResponse structure to hold response for disable recurring request
type RecurringDisableResponse struct {
	Response string `json:"response"`
}