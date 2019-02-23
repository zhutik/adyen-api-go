package adyen

/*******************
* Payment 3 DS 2.0 *
*                  *
* Requires v40 API *
*******************/

// Authorise 3DS 2.0 Device channel list
const (
	DeviceChannelApp     = "app"
	DeviceChannelBrowser = "browser"
)

// Authorise3DS2 - Authorise 3DS 2.0 request
//
// Description - https://docs.adyen.com/api-explorer/#/Payment/v40/authorise3ds2
type Authorise3DS2 struct {
	AccountInfo         *AccountInfo         `json:"accountInfo,omitempty"`
	Amount              *Amount              `json:"amount"`
	BillingAddress      *Address             `json:"billingAddress,omitempty"`
	DeliveryAddress     *Address             `json:"deliveryAddress,omitempty"`
	BrowserInfo         *BrowserInfo         `json:"browserInfo,omitempty"`
	MerchantAccount     string               `json:"merchantAccount"`
	ShopperReference    string               `json:"shopperReference,omitempty"` // Mandatory for recurring payment
	Recurring           *Recurring           `json:"recurring,omitempty"`
	ShopperEmail        string               `json:"shopperEmail,omitempty"`
	ShopperInteraction  string               `json:"shopperInteraction,omitempty"`
	ShopperIP           string               `json:"shopperIP,omitempty"`
	ShopperLocale       string               `json:"shopperLocale,omitempty"`
	ShopperName         *Name                `json:"shopperName,omitempty"`
	ThreeDS2RequestData *ThreeDS2RequestData `json:"threeDS2RequestData,omitempty"`
	ThreeDS2Result      *ThreeDS2Result      `json:"threeDS2Result,omitempty"`
	ThreeDS2Token       string               `json:"threeDS2Token,omitempty"`
}

// ThreeDS2RequestData - Request fields for 3D Secure 2.0
//
// Description - https://docs.adyen.com/developers/risk-management/3d-secure-2-0/server-integration/api-reference-3d-secure-2-0#threeds2requestdata
type ThreeDS2RequestData struct {
	AuthenticationOnly bool   `json:"authenticationOnly,omitempty"`
	ChallengeIndicator string `json:"challengeIndicator,omitempty"`
	DeviceChannel      string `json:"deviceChannel"`
	NotificationURL    string `json:"notificationURL,omitempty"` // Only for deviceChannel browser
	ThreeDSCompInd     string `json:"threeDSCompInd,omitempty"`
}

// ThreeDS2Result - 3DS 2.0 Results
//
// Description - https://docs.adyen.com/developers/risk-management/3d-secure-2-0/server-integration/api-reference-3d-secure-2-0#threeds2result
type ThreeDS2Result struct {
	AuthenticationValue  string `json:"authenticationValue"`
	Eci                  string `json:"eci"`
	ThreeDSServerTransID string `json:"threeDSServerTransID"`
	Timestamp            string `json:"timestamp"`
	TransStatus          string `json:"transStatus"`
	TransStatusReason    string `json:"transStatusReason"`
}
