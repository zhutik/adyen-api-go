package adyen

import "time"

// NotificationRequest contains environment specification and list of notifications to process
//
// Link - https://docs.adyen.com/developers/api-reference/notifications-api#notificationrequest
type NotificationRequest struct {
	Live              convertibleBoolean        `json:"live"`
	NotificationItems []NotificationRequestItem `json:"notificationItems"`
}

// NotificationRequestItem contains notification details
//
// Depending on notification type, different fields can be populated and send from Adyen
//
// Link - https://docs.adyen.com/developers/api-reference/notifications-api#notificationrequestitem
type NotificationRequestItem struct {
	NotificationRequestItem struct {
		AdditionalData struct {
			ShopperReference         string `json:"shopperReference,omitempty"`
			ShopperEmail             string `json:"shopperEmail,omitempty"`
			AuthCode                 string `json:"authCode,omitempty"`
			CardSummary              string `json:"cardSummary,omitempty"`
			ExpiryDate               string `json:"expiryDate,omitempty"`
			AuthorisedAmountValue    string `json:"authorisedAmountValue,omitempty"`
			AuthorisedAmountCurrency string `json:"authorisedAmountCurrency,omitempty"`
		} `json:"additionalData,omitempty"`
		Amount              Amount             `json:"amount"`
		PspReference        string             `json:"pspReference"`
		EventCode           string             `json:"eventCode"`
		// Event date in time.RFC3339 format
		EventDate           time.Time          `json:"eventDate"`
		MerchantAccountCode string             `json:"merchantAccountCode"`
		Operations          []string           `json:"operations"`
		MerchantReference   string             `json:"merchantReference"`
		OriginalReference   string             `json:"originalReference,omitempty"`
		PaymentMethod       string             `json:"paymentMethod"`
		Reason              string             `json:"reason,omitempty"`
		Success             convertibleBoolean `json:"success"`
	} `json:"NotificationRequestItem"`
}