package adyen

// RecurringGateway - Adyen recurring transaction logic
type RecurringGateway struct {
	*Adyen
}

// listRecurringDetailsType - listRecurringDetails type request, @TODO: move to enums
const listRecurringDetailsType = "listRecurringDetails"

// AuthoriseEncrypted - Perform authorise payment in Adyen
func (a *RecurringGateway) ListRecurringDetails(req *RecurringDetailsRequest) (*RecurringDetailsResult, error) {
	resp, err := a.execute(listRecurringDetailsType, RecurringService, req)

	if err != nil {
		return nil, err
	}

	return resp.listRecurringDetails()
}
