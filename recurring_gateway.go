package adyen

// RecurringGateway - Adyen recurring transaction logic
type RecurringGateway struct {
	*Adyen
}

// listRecurringDetailsType - listRecurringDetails type request, @TODO: move to enums
const listRecurringDetailsType = "listRecurringDetails"

// ListRecurringDetails - Get list of recurring payments in Adyen
func (a *RecurringGateway) ListRecurringDetails(req *RecurringDetailsRequest) (*RecurringDetailsResult, error) {
	resp, err := a.execute(RecurringService, listRecurringDetailsType, req)

	if err != nil {
		return nil, err
	}

	return resp.listRecurringDetails()
}
