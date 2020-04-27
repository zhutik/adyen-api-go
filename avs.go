package adyen

// AVSResponse is a type definition for all possible responses from Adyen's AVS system
//
// https://docs.adyen.com/risk-management/avs-checks
type AVSResponse string

// AVSResponse hard-coded for easy comparison checking later
const (
	AVSResponse0  AVSResponse = "0 Unknown"
	AVSResponse1  AVSResponse = "1 Address matches, postal code doesn't"
	AVSResponse2  AVSResponse = "2 Neither postal code nor address match"
	AVSResponse3  AVSResponse = "3 AVS unavailable"
	AVSResponse4  AVSResponse = "4 AVS not supported for this card type"
	AVSResponse5  AVSResponse = "5 No AVS data provided"
	AVSResponse6  AVSResponse = "6 Postal code matches, but the address does not match"
	AVSResponse7  AVSResponse = "7 Both postal code and address match"
	AVSResponse8  AVSResponse = "8 Address not checked, postal code unknown"
	AVSResponse9  AVSResponse = "9 Address matches, postal code unknown"
	AVSResponse10 AVSResponse = "10 Address doesn't match, postal code unknown"
	AVSResponse11 AVSResponse = "11 Postal code not checked, address unknown"
	AVSResponse12 AVSResponse = "12 Address matches, postal code not checked"
	AVSResponse13 AVSResponse = "13 Address doesn't match, postal code not checked"
	AVSResponse14 AVSResponse = "14 Postal code matches, address unknown"
	AVSResponse15 AVSResponse = "15 Postal code matches, address not checked"
	AVSResponse16 AVSResponse = "16 Postal code doesn't match, address unknown"
	AVSResponse17 AVSResponse = "17 Postal code doesn't match, address not checked."
	AVSResponse18 AVSResponse = "18 Neither postal code nor address were checked"
	AVSResponse19 AVSResponse = "19 Name and postal code matches"
	AVSResponse20 AVSResponse = "20 Name, address and postal code matches"
	AVSResponse21 AVSResponse = "21 Name and address matches"
	AVSResponse22 AVSResponse = "22 Name matches"
	AVSResponse23 AVSResponse = "23 Postal code matches, name doesn't match"
	AVSResponse24 AVSResponse = "24 Both postal code and address matches, name doesn't match"
	AVSResponse25 AVSResponse = "25 Address matches, name doesn't match"
	AVSResponse26 AVSResponse = "26 Neither postal code, address nor name matches"
)
