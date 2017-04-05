# [WIP] Adyen API for Go

## Usage

```
// Configure Adyen API
instance := adyen.New(
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
  os.Getenv("ADYEN_CLIENT_TOKEN"),
  os.Getenv("ADYEN_ACCOUNT"),
)

req := &adyen.Authorise{
  Amount:          &adyen.Amount{
    Value: 1000, / amount * 100, f.e. 10,30 EUR = 1030
    Currency: "EUR"
  },
  MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
  AdditionalData:  &adyen.AdditionalData{Content: "encryptedData"},
  Reference:       "your-order-number",
}

// Perform authorise transaction
g, err := instance.Payment().Authorise(req)

Supported API Calls
* Authorise (only encrypted)
* Capture
* Cancel
* [NEXT] Refund
* [NEXT] CancelOrRefund

## To run example

### Expose your settings for Adyen API configuration.

```server.go``` script will use those variables to communicate with API

```
$ export ADYEN_CLIENT_TOKEN="YOUR_ADYEN_CLIENT_TOKEN"
$ export ADYEN_USERNAME="YOUR_ADYEN_API_USERNAME"
$ export ADYEN_PASSWORD="YOUR_API_PASSWORD"
$ export ADYEN_ACCOUNT="YOUR_MERCHANT_ACCOUNT"
```

Settings explanation:
* ADYEN_CLIENT_TOKEN - Library token in Adyen, used to load external JS file from Adyen to validate Credit Card information
* ADYEN_USERNAME - Adyen API username, usually starts with ws@
* ADYEN_PASSWORD - Adyen API password for username
* ADYEN_ACCOUNT - Selected Merchant Account

### Run example application
```
$ cd example
$ go run server.go
```

### Perform payments

Open http://localhost:8080 in your browser
Put credit card information.

Test credit cards could be found https://docs.adyen.com/support/integration#testcardnumbers
