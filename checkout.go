package adyen

// PaymentMethods contains the fields required by the checkout
// API's /paymentMethods endpoint.  See the following for more
// information:
//
// https://docs.adyen.com/api-explorer/#/PaymentSetupAndVerificationService/v32/paymentMethods
type PaymentMethods struct {
	Amount           *Amount `json:"amount"`
	Channel          string  `json:"channel"`
	CountryCode      string  `json:"countryCode"`
	MerchantAccount  string  `json:"merchantAccount"`
	ShopperLocale    string  `json:"shopperLocale"`
	ShopperReference string  `json:"shopperReference"`
}

// PaymentMethodsResponse is returned by Adyen in response to
// a PaymentMethods request.
type PaymentMethodsResponse struct {
	PaymentMethods         []PaymentMethodDetails         `json:"paymentMethods"`
	OneClickPaymentMethods []OneClickPaymentMethodDetails `json:"oneClickPaymentMethods,omitempty"`
}

// PaymentMethodDetails describes the PaymentMethods part of
// a PaymentMethodsResponse.
type PaymentMethodDetails struct {
	Details []PaymentMethodDetailsInfo `json:"details,omitempty"`
	Name    string                     `json:"name"`
	Type    string                     `json:"type"`
}

// PaymentMethodDetailsInfo describes the collection of all
// payment methods.
type PaymentMethodDetailsInfo struct {
	Items []PaymentMethodItems `json:"items"`
	Key   string               `json:"key"`
	Type  string               `json:"type"`
}

// PaymentMethodItems describes a single payment method.
type PaymentMethodItems struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OneClickPaymentMethodDetails describes the OneClickPayment part of
// a PaymentMethods response.
type OneClickPaymentMethodDetails struct {
	Details       []PaymentMethodTypes       `json:"details"`
	Name          string                     `json:"name"`
	Type          string                     `json:"type"`
	StoredDetails PaymentMethodStoredDetails `json:"storedDetails"`
}

// PaymentMethodTypes describes any additional information associated
// with a OneClick payment.
type PaymentMethodTypes struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// PaymentMethodStoredDetails describes the information stored for a
// OneClick payment.
type PaymentMethodStoredDetails struct {
	Card PaymentMethodCard `json:"card"`
}

// PaymentMethodCard describes the card information associated with a
// OneClick payment.
type PaymentMethodCard struct {
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	HolderName  string `json:"holderName"`
	Number      string `json:"number"`
}
