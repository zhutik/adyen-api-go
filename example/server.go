/*
	export ADYEN_ENC_URL="YOUR_ADYEN_ENCRYPTED_URL"
	export ADYEN_USERNAME="YOUR_ADYEN_API_USERNAME"
	export ADYEN_PASSWORD="YOUR_API_PASSWORD"
	export ADYEN_ACCOUNT="YOUR_MERCHANT_ACCOUNT"
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

type AdditionalData struct {
	Content string `json:"card.encrypted.json"`
}

type AdyenRequest struct {
	AdditionalData  AdditionalData `json:"additionalData"`
	Amount          Amount         `json:"amount"`
	Reference       string         `json:"reference"`
	MerchantAccount string         `json:"merchantAccount"`
}

type AdyenResponse struct {
	PspReference  string `json:"pspReference"`
	ResultCode    string `json:"resultCode"`
	AuthCode      string `json:"authCode"`
	RefusalReason string `json:"refusalReason"`
}

type AdyenConfig struct {
	ApiURL          string
	EncURL          string
	Username        string
	Password        string
	MerchantAccount string
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

/**
 * Show Adyen Payment form
 */
func showForm(w http.ResponseWriter, r *http.Request) {
	var config AdyenConfig
	config.EncURL = os.Getenv("ADYEN_ENC_URL")

	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, config)
}

func performPayment(w http.ResponseWriter, r *http.Request) {
	var config AdyenConfig

	config.ApiURL = "https://pal-test.adyen.com/pal/servlet/Payment/v25/authorise"
	config.EncURL = os.Getenv("ADYEN_ENC_URL")
	config.Username = os.Getenv("ADYEN_USERNAME")
	config.Password = os.Getenv("ADYEN_PASSWORD")
	config.MerchantAccount = os.Getenv("ADYEN_ACCOUNT")

	r.ParseForm()

	var adyenReq AdyenRequest
	var amount Amount
	amount.Currency = "EUR"
	amount.Value = 1000 // should be amount * 100, f.e. 20.00 = 2000
	var addInfo AdditionalData
	addInfo.Content = r.Form.Get("adyen-encrypted-data")

	rand.Seed(time.Now().UTC().UnixNano())

	adyenReq.Amount = amount
	adyenReq.Reference = "DE-100" + randomString(6)
	adyenReq.MerchantAccount = config.MerchantAccount
	adyenReq.AdditionalData = addInfo

	body, _ := json.Marshal(adyenReq)

	req, err := http.NewRequest("POST", config.ApiURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.Username, config.Password)

	client := &http.Client{}
	resp, err := client.Do(req)

	var responseValues AdyenResponse

	json.NewDecoder(resp.Body).Decode(&responseValues)

	if err == nil {
		fmt.Fprintf(w, "<h1>Success!</h1><code><pre>"+responseValues.AuthCode+" "+responseValues.PspReference+"</pre></code>")
	} else {
		fmt.Fprintf(w, "<h1>Something went wrong: "+err.Error()+"</h1>")
	}
}

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/perform_payment", performPayment)
	http.ListenAndServe(":8080", nil)
}
