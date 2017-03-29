/*
	export ADYEN_CLIENT_TOKEN="YOUR_ADYEN_ENCRYPTED_URL"
	export ADYEN_USERNAME="YOUR_ADYEN_API_USERNAME"
	export ADYEN_PASSWORD="YOUR_API_PASSWORD"
	export ADYEN_ACCOUNT="YOUR_MERCHANT_ACCOUNT"
*/

package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"

	adyen "github.com/zhutik/adyen-api-go"
)

type TempateConfig struct {
	EncURL string
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
	adyen := adyen.New(
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_CLIENT_TOKEN"),
		os.Getenv("ADYEN_ACCOUNT"),
	)
	config := TempateConfig{
		EncURL: adyen.ClientURL(),
	}

	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, config)
}

/**
 * Handle post request and perform payment authosization
 */
func performPayment(w http.ResponseWriter, r *http.Request) {
	adyen := adyen.New(
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_CLIENT_TOKEN"),
		os.Getenv("ADYEN_ACCOUNT"),
	)

	r.ParseForm()

	rand.Seed(time.Now().UTC().UnixNano())

	g, err := adyen.Authorise().Payment(
		r.Form.Get("adyen-encrypted-data"),
		"DE-100"+randomString(6),
		1000,
	)

	if err == nil {
		fmt.Fprintf(w, "<h1>Success!</h1><code><pre>"+g.AuthCode+" "+g.PspReference+"</pre></code>")
	} else {
		fmt.Fprintf(w, "<h1>Something went wrong: "+err.Error()+"</h1>")
	}
}

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/perform_payment", performPayment)
	http.ListenAndServe(":8080", nil)
}
