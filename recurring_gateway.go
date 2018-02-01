package adyen

// RecurringGateway - Adyen recurring transaction logic
type RecurringGateway struct {
	*Adyen
}

const (
	// listRecurringDetailsType - listRecurringDetails type request, @TODO: move to enums
	listRecurringDetailsType = "listRecurringDetails"
	// disableRecurringType - disable recurring type request, @TODO: move to enums
	disableRecurringType = "disable"
)

// ListRecurringDetails - Get list of recurring payments in Adyen
func (a *RecurringGateway) ListRecurringDetails(req *RecurringDetailsRequest) (*RecurringDetailsResult, error) {
	url := a.adyenURL(RecurringService, listRecurringDetailsType, RecurringAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.listRecurringDetails()
}

// DisableRecurring - disable customer's saved payment method based on a contract type or/and payment method ID
func (a *RecurringGateway) DisableRecurring(req *RecurringDisableRequest) (*RecurringDisableResponse, error) {
	url := a.adyenURL(RecurringService, disableRecurringType, RecurringAPIVersion)

	resp, err := a.execute(url, req)

	if err != nil {
		return nil, err
	}

	return resp.disableRecurring()
}
