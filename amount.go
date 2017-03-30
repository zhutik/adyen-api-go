package adyen

// Amount value/currency representation
type Amount struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}
