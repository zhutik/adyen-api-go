package adyen

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AuthoriseGateway struct {
	*Adyen
}

const RequestType = "authorise"

func (a *AuthoriseGateway) Payment(encryptedData string, reference string, amount float32) (*AuthoriseResponse, error) {
	auth := AuthoriseRequest{
		Amount:          Amount{Value: amount, Currency: "EUR"},
		MerchantAccount: a.MerchantAccount,
		AdditionalData:  AdditionalData{Content: encryptedData},
		Reference:       reference,
	}

	// move request logic to Adyen
	body, _ := json.Marshal(auth)

	req, err := http.NewRequest("POST", a.AdyenUrl(RequestType), bytes.NewBuffer(body))
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
