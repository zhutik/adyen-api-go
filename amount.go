package adyen

import "math"

// Amount value/currency representation
type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

var (
	// DefaultCurrencyDecimals - default currency decimals
	DefaultCurrencyDecimals uint = 2

	// CurrencyDecimals - https://docs.adyen.com/developers/currency-codes
	// currencies with 2 decimals stripped out
	CurrencyDecimals = map[string]uint{
		"BHD": 3,
		"CVE": 0,
		"DJF": 0,
		"GNF": 0,
		"IDR": 0,
		"JOD": 3,
		"JPY": 0,
		"KMF": 0,
		"KRW": 0,
		"KWD": 3,
		"LYD": 3,
		"OMR": 3,
		"PYG": 0,
		"RWF": 0,
		"TND": 3,
		"UGX": 0,
		"VND": 0,
		"VUV": 0,
		"XAF": 0,
		"XOF": 0,
		"XPF": 0,
	}
)

// NewAmount - creates Amount instance
//
// Automatically adjust decimal points for the float value
// Link - https://docs.adyen.com/developers/development-resources/currency-codes
func NewAmount(currency string, amount float32) *Amount {
	decimals, ok := CurrencyDecimals[currency]
	if !ok {
		decimals = DefaultCurrencyDecimals
	}

	if decimals == 0 {
		return &Amount{
			Currency: currency,
			Value:    amount,
		}
	}

	coef := float32(math.Pow(10, float64(decimals)))

	return &Amount{
		Currency: currency,
		Value:   float32(math.Round(float64(amount * coef))),
	}
}
