package adyen

// AdyenAVSResponse is a type definition for all possible responses from Adyen's AVS system
//
// https://docs.adyen.com/risk-management/avs-checks
type AdyenAVSResponse string

const (
	AdyenAVSResponse0  AdyenAVSResponse = "0 Unknown"
	AdyenAVSResponse1  AdyenAVSResponse = "1 Address matches, but the postal code does not match"
	AdyenAVSResponse2  AdyenAVSResponse = "2 Neither postal code nor address match"
	AdyenAVSResponse3  AdyenAVSResponse = "3 AVS unavailable"
	AdyenAVSResponse4  AdyenAVSResponse = "4 AVS not supported for this card type"
	AdyenAVSResponse5  AdyenAVSResponse = "5 No AVS data provided"
	AdyenAVSResponse6  AdyenAVSResponse = "6 Postal code matches, but the address does not match"
	AdyenAVSResponse7  AdyenAVSResponse = "7 Both postal code and address match"
	AdyenAVSResponse8  AdyenAVSResponse = "8 Address not checked, postal code unknown"
	AdyenAVSResponse9  AdyenAVSResponse = "9 Address matches, postal code unknown"
	AdyenAVSResponse10 AdyenAVSResponse = "10 Address doesn't match, postal code unknown"
	AdyenAVSResponse11 AdyenAVSResponse = "11 Postal code not checked, address unknown"
	AdyenAVSResponse12 AdyenAVSResponse = "12 Address matches, postal code not checked"
	AdyenAVSResponse13 AdyenAVSResponse = "13 Address doesn't match, postal code not checked"
	AdyenAVSResponse14 AdyenAVSResponse = "14 Postal code matches, address unknown"
	AdyenAVSResponse15 AdyenAVSResponse = "15 Postal code matches, address not checked"
	AdyenAVSResponse16 AdyenAVSResponse = "16 Postal code doesn't match, address unknown"
	AdyenAVSResponse17 AdyenAVSResponse = "17 Postal code doesn't match, address not checked."
	AdyenAVSResponse18 AdyenAVSResponse = "18 Neither postal code nor address were checked"
	AdyenAVSResponse19 AdyenAVSResponse = "19 Name and postal code matches"
	AdyenAVSResponse20 AdyenAVSResponse = "20 Name, address and postal code matches"
	AdyenAVSResponse21 AdyenAVSResponse = "21 Name and address matches"
	AdyenAVSResponse22 AdyenAVSResponse = "22 Name matches"
	AdyenAVSResponse23 AdyenAVSResponse = "23 Postal code matches, name doesn't match"
	AdyenAVSResponse24 AdyenAVSResponse = "24 Both postal code and address matches, name doesn't match"
	AdyenAVSResponse25 AdyenAVSResponse = "25 Address matches, name doesn't match"
	AdyenAVSResponse26 AdyenAVSResponse = "26 Neither postal code, address nor name matches"
)
