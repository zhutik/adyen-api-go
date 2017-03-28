# [WIP] Adyen API for Go

## To run example

Expose your settings for Adyen API configuration.

```server.go``` script will use those variables to communicate with API

```
$ export ADYEN_ENC_URL="YOUR_ADYEN_ENCRYPTED_URL"
$ export ADYEN_USERNAME="YOUR_ADYEN_API_USERNAME"
$ export ADYEN_PASSWORD="YOUR_API_PASSWORD"
$ export ADYEN_ACCOUNT="YOUR_MERCHANT_ACCOUNT"
```

Run example application
```
$ cd example
$ go run server.go
```

Open http://localhost:8080 in your browser
Put credit card information.

Test credit cards could be found https://docs.adyen.com/support/integration#testcardnumbers
