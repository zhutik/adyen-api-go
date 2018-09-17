# Adyen API for Go

[![Build Status](https://travis-ci.org/zhutik/adyen-api-go.png)](https://travis-ci.org/zhutik/adyen-api-go)
[![GoDoc](http://godoc.org/github.com/zhutik/adyen-api-go?status.png)](http://godoc.org/github.com/zhutik/adyen-api-go)
[![Coverage Status](https://coveralls.io/repos/github/zhutik/adyen-api-go/badge.svg)](https://coveralls.io/github/zhutik/adyen-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhutik/adyen-api-go)](https://goreportcard.com/report/github.com/zhutik/adyen-api-go)

A Go client library for [Adyen](https://www.adyen.com/en/) payments platform.

This is *not* an official client library. Adyen has official libraries for multiple platforms [Github](https://github.com/adyen/), but not Go yet.

This package provides core functionality to perform most common types of a Payment requests to an API. 
If you see some functionality is missing, please, open an issue (or better yet, a pull request).

## Installation

```
go get github.com/zhutik/adyen-api-go
```

## Playground and examples

Please check separate repository with Adyen API playground, where you can test API
and get some usage examples for Adyen API library

https://github.com/zhutik/adyen-api-go-example

Or you can visit [Wiki page](https://github.com/zhutik/adyen-api-go/wiki) for more details and examples

## Supported API Calls

* Authorise (Encrypted in recommended)
* Authorise 3D
* Recurring payments and retrieving stored payment methods
* Capture
* Cancel
* Refund (CancelOrRefund)
* Notifications

## Usage

```go
import "github.com/zhutik/adyen-api-go"

// Configure Adyen API
instance := adyen.New(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
)

amount := &adyen.Amount{
    Value:    1000, // amount * 100, f.e. 10,30 EUR = 1030
    Currency: "EUR" // or use instance.Currency
}

// or amount := adyen.NewAmount("EUR", 10), in this case decimal points would be adjusted automatically

req := &adyen.AuthoriseEncrypted{
  Amount: amount,
  MerchantAccount: os.Getenv("ADYEN_ACCOUNT"), // your merchant account in Adyen
  AdditionalData:  &adyen.AdditionalData{Content: "encryptedData"}, // encrypted data from a form
  Reference:       "your-order-number",
}

// Perform authorise transaction
g, err := instance.Payment().AuthoriseEncrypted(req)

```

Load Client Side JS for form encryption to include on credit card form page

```go
// Configure Adyen API
instance := adyen.New(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
)

url := &adyen.ClientURL(os.Getenv("ADYEN_CLIENT_TOKEN"))
```

Currently, MerchantAccount and Currency need to be set for every request manually

To shortcut configuration, additional methods could be used to set and retrieve those settings.

```go
// Configure Adyen API
instance := adyen.New(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
)

// set parameters once for current instance
instance.Currency = "USD"
instance.MerchantAccount = "TEST_MERCHANT_ACCOUNT"

// futher, information could be retrieved to populate request 
req := &adyen.AuthoriseEncrypted{
  Amount: adyen.NewAmount(instance.Currency, 10.00),
  MerchantAccount: instance.MerchantAccount,
  AdditionalData:  &adyen.AdditionalData{Content: "encryptedData"}, // encrypted data from a form
  Reference:       "your-order-number",
}
```

### Environment configuration

Adyen's Production environment requires additional configuration to the Test environment for security reasons.  Namely, this includes a random hexadecimal string that's generated for your account and the company account name.

In the following examples, the environmenst have been hard-coded for clarity.  They would typically come from environments variables instead.

To target the Test environment:

``` go
env := adyen.TestEnvironment()
```

To target the Production environment:

``` go
env, err := adyen.ProductionEnvironment("5409c4fd1cc98a4e", "AcmeAccount123")
```

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
instance := adyen.NewWithHMAC(
  adyen.Testing,
  os.Getenv("ADYEN_USERNAME"),
  os.Getenv("ADYEN_PASSWORD"),
  os.Getenv("ADYEN_HMAC"),
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

## TODOs

* Move some constants into enum files.
* Parse URLs for environment's BaseURL, ClientURL and HppURL methods instead of string concatenation (needs to return an error as well).
* Reduced API surface by making most types and functions unexported.
