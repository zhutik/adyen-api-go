package adyen

import (
	"testing"
	"strings"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"time"
)

// TestNotificationRequest - test adyen notification JSON conversion
func TestNotificationRequest(t *testing.T) {
	t.Parallel()

	responseJSON := `
{
	"live":"false",
	"notificationItems":[
		{
			"NotificationRequestItem":{
				"additionalData":{
					"cardSummary":"7777",
					"eci":"N\/A",
					"shopperIP":"127.0.0.1",
					"totalFraudScore":"10",
					"expiryDate":"12\/2012",
					"xid":"AAE=",
					"billingAddress.street":"Nieuwezijds Voorburgwal",
					"cavvAlgorithm":"N\/A",
					"cardBin":"976543",
					"extraCostsValue":"101",
					"billingAddress.city":"Amsterdam",
					"threeDAuthenticated":"false",
					"alias":"H934380689410347",
					"paymentMethodVariant":"visa",
					"billingAddress.country":"NL",
					"fraudCheck-6-ShopperIpUsage":"10",
					"deviceType":"Other",
					" NAME1 ":"VALUE1",
					"authCode":"1234",
					"cardHolderName":"J. De Tester",
					"threeDOffered":"false",
					"billingAddress.houseNumberOrName":"21 - 5",
					"threeDOfferedResponse":"N\/A",
					"NAME2":"  VALUE2  ",
					"billingAddress.postalCode":"1012RC",
					"browserCode":"Other",
					"cavv":"AAE=",
					"issuerCountry":"unknown",
					"threeDAuthenticatedResponse":"N\/A",
					"aliasType":"Default",
					"extraCostsCurrency":"EUR",
					"captureDelayHours":"120"
				},
				"amount":{
					"currency":"EUR",
					"value":10100
				},
				"eventCode":"AUTHORISATION",
				"eventDate":"2017-12-27T14:53:06+01:00",
				"merchantAccountCode":"TestCOM148",
				"merchantReference":"8313842560770001",
				"operations":["CANCEL","CAPTURE","REFUND"],
				"paymentMethod":"visa",
				"pspReference":"test_AUTHORISATION_1",
				"reason":"1234:7777:12\/2012",
				"success":"true"
			}
		}
	]
}
`
	body := strings.NewReader(responseJSON)

	resp := &http.Response{
		Status:        "OK 200",
		StatusCode:    200,
		ContentLength: int64(body.Len()),
		Body:          ioutil.NopCloser(body),
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)

	if err != nil {
		t.Error(err)
	}

	var notification NotificationRequest

	if err := json.Unmarshal(buf.Bytes(), &notification); err != nil {
		t.Error(err)
	}

	if notification.Live != false {
		t.Errorf("Expected notification environment should not be live, %t given", notification.Live)
	}

	if len(notification.NotificationItems) != 1 {
		t.Errorf("Expected to have only one notification element in a list, %d given", len(notification.NotificationItems))
	}

	item := notification.NotificationItems[0].NotificationRequestItem

	if item.EventCode != "AUTHORISATION" {
		t.Errorf("Expected to have AUTHORISATION event code, %s given", item.EventCode)
	}

	if item.EventDate.Format(time.RFC3339) != "2017-12-27T14:53:06+01:00" {
		t.Errorf("Expected to have 2017-12-27T14:53:06+01:00 event date, %s given", item.EventDate.Format(time.RFC3339))
	}

	if item.MerchantAccountCode != "TestCOM148" {
		t.Errorf("Expected to have TestCOM148 merchant account code, %s given", item.MerchantAccountCode)
	}

	if item.MerchantReference != "8313842560770001" {
		t.Errorf("Expected to have 8313842560770001 merchant reference, %s given", item.MerchantReference)
	}

	if strings.Join(item.Operations, ",") != "CANCEL,CAPTURE,REFUND" {
		t.Errorf("Expected to have CANCEL,CAPTURE,REFUND operations available, %s given", strings.Join(item.Operations, ","))
	}

	if item.PaymentMethod != "visa" {
		t.Errorf("Expected to have visa payment, %s given", item.PaymentMethod)
	}

	if item.PspReference != "test_AUTHORISATION_1" {
		t.Errorf("Expected to have test_AUTHORISATION_1 as pspRegerence, %s given", item.PspReference)
	}

	if item.Reason != "1234:7777:12/2012" {
		t.Errorf("Expected to have 1234:7777:12/2012 reason, %s given", item.Reason)
	}

	if item.Success != true {
		t.Errorf("Expected to have successful notification, %t given", item.Success)
	}
}
