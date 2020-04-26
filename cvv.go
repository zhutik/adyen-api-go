package adyen

// CVCResult represents the Adyen translation of CVC codes from issuer
// https://docs.adyen.com/development-resources/test-cards/cvc-cvv-result-testing
type CVCResult string
// Constants represented by numerical code they are assigned
const (
	CVCResult0 CVCResult = "0 Unknown"
	CVCResult1 CVCResult = "1 Matches"
	CVCResult2 CVCResult = "2 Doesn't Match"
	CVCResult3 CVCResult = "3 Not Checked"
	CVCResult4 CVCResult = "4 No CVC/CVV provided, but was required"
	CVCResult5 CVCResult = "5 Issuer not certified for CVC/CVV"
	CVCResult6 CVCResult = "6 No CVC/CVV provided"
)


