# Adyen API for Go

[![Build Status](https://travis-ci.org/zhutik/adyen-api-go.png)](https://travis-ci.org/zhutik/adyen-api-go)
[![GoDoc](http://godoc.org/github.com/zhutik/adyen-api-go?status.png)](http://godoc.org/github.com/zhutik/adyen-api-go)

## Install

```
go get github.com/zhutik/adyen-api-go
```

## Playgroup and examples

Please check separate repository with Adyen API playgroup, where you can test API
and get some usage example for Adyen API library

https://github.com/zhutik/adyen-api-go-example

## Usage

```go
import "github.com/zhutik/adyen-api-go"

// Configure Adyen API
instance := adyen.New(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
  os.Getenv("ADYEN_CLIENT_TOKEN"),
  os.Getenv("ADYEN_ACCOUNT"),
)

req := &adyen.Authorise{
  Amount: &adyen.Amount{
    Value:    1000, // amount * 100, f.e. 10,30 EUR = 1030
    Currency: "EUR"
  },
  MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
  AdditionalData:  &adyen.AdditionalData{Content: "encryptedData"},
  Reference:       "your-order-number",
}

// Perform authorise transaction
g, err := instance.Payment().Authorise(req)

```

Supported API Calls
* Authorise (only encrypted)
* Authorise 3D
* Recurring payments
* Capture
* Cancel
* Refund
* CancelOrRefund
* Notifications

Next:
* Add more tests

## To run example

### Expose your settings for Adyen API configuration.

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

## Hosted Payment Pages

Update your settings to include

```
$ export ADYEN_HMAC="YOUR_HMAC_KEY"
$ export ADYEN_SKINCODE="YOUR_SKINCODE_ID"
$ export ADYEN_SHOPPER_LOCALE="YOUR_SHOPPER_LOCALE"
```

Use HPP constructor to initialize new Adyen API instance

```go
import "github.com/zhutik/adyen-api-go"

// Configure Adyen API
instance := adyen.NewWithHPP(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
  os.Getenv("ADYEN_CLIENT_TOKEN"),
  os.Getenv("ADYEN_ACCOUNT"),
  os.Getenv("ADYEN_HMAC"),
  os.Getenv("ADYEN_SKINCODE"),
  os.Getenv("ADYEN_SHOPPER_LOCALE"),
)

```

Perform requests as usual:

```go
timeIn := time.Now().Local().Add(time.Minute * time.Duration(60))

req := &adyen.DirectoryLookupRequest{
    CurrencyCode:      "EUR",
    MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
    PaymentAmount:     1000,
    SkinCode:          os.Getenv("ADYEN_SKINCODE"),
    MerchantReference: "your-order-number",
    SessionsValidity:  timeIn.Format(time.RFC3339),
}

g, err := instance.Payment().DirectoryLookup(req)

```

or generate redirect URL for selected payment method

Example with iDEAL for Netherlands:

```go
timeIn := time.Now().Local().Add(time.Minute * time.Duration(60))

req := &adyen.SkipHppRequest{
    MerchantReference: "your-order-number",
    PaymentAmount:     1000,
    CurrencyCode:      "EUR",
    ShipBeforeDate:    timeIn.Format(time.RFC3339),
    SkinCode:          os.Getenv("ADYEN_SKINCODE"),
    MerchantAccount:   os.Getenv("ADYEN_ACCOUNT"),
    ShopperLocale:     "nl",
    SessionsValidity:  timeIn.Format(time.RFC3339),
		CountryCode:       "NL",
		BrandCode:         "ideal",
		IssuerID:          "1121",
}

url, err := instance.Payment().GetHPPRedirectURL(req)

http.Redirect(w, r, url, http.StatusTemporaryRedirect)
```

Supported Calls:
* Directory Lookup
* Locale payment methods redirect

### Setup playgroup

Please check separate repository with Adyen API playgroup, where you can test API
and get some usage example for Adyen API library

https://github.com/zhutik/adyen-api-go-example

### Perform payments

Open http://localhost:8080 in your browser and check implemented actions.

Test credit cards could be found https://docs.adyen.com/support/integration#testcardnumbers
