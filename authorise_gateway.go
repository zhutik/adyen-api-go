package adyen

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AuthoriseGateway struct {
	*Adyen
}

func (a *AuthoriseGateway) Payment(encryptedData string, reference string, amount float32) (*AuthoriseResponse, error) {
	auth := AuthoriseRequest{
		Amount:          Amount{Value: amount, Currency: "EUR"},
		MerchantAccount: a.MerchantAccount,
		AdditionalData:  AdditionalData{Content: encryptedData},
		Reference:       reference,
	}

	body, _ := json.Marshal(auth)
	url := "https://pal-test.adyen.com/pal/servlet/Payment/v25/authorise"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(a.Username, a.Password)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var val AuthoriseResponse
	json.NewDecoder(resp.Body).Decode(&val)

	return &val, nil
}
