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
