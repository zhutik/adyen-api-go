package adyen

// RecurringGateway - Adyen recurring transaction logic
type RecurringGateway struct {
	*Adyen
}

// listRecurringDetailsType - listRecurringDetails type request, @TODO: move to enums
const listRecurringDetailsType = "listRecurringDetails"

// disableRecurringType - disable recurring type request, @TODO: move to enums
const disableRecurringType = "disable"

// ListRecurringDetails - Get list of recurring payments in Adyen
func (a *RecurringGateway) ListRecurringDetails(req *RecurringDetailsRequest) (*RecurringDetailsResult, error) {
	resp, err := a.execute(RecurringService, listRecurringDetailsType, req)

	if err != nil {
		return nil, err
	}

	return resp.listRecurringDetails()
}

// DisableRecurring - disable customer's saved payment method based on a contract type or/and payment method ID
func (a *RecurringGateway) DisableRecurring(req *RecurringDisableRequest) (*RecurringDisableResponse, error) {
	resp, err := a.execute(RecurringService, disableRecurringType, req)

	if err != nil {
		return nil, err
	}

	return resp.disableRecurring()
}
