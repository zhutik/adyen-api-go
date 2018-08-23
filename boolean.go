package adyen

import (
	"strconv"
	"strings"
)

// StringBool allows us to unmarhsal Adyen Boolean values
// which appear as strings instead of bools.
type StringBool bool

func NewStringBool(b bool) *StringBool {
	sb := StringBool(b)
	return &sb
}

func (b *StringBool) UnmarshalJSON(data []byte) (err error) {
	str := strings.TrimFunc(strings.ToLower(string(data)), func(c rune) bool {
		return c == ' ' || c == '"'
	})

	parsed, err := strconv.ParseBool(str)
	if err != nil {
		return
	}

	*b = StringBool(parsed)
	return
}

func (b StringBool) MarshalJSON() ([]byte, error) {
	boolResult := bool(b)
	var boolString string

	if boolResult {
		boolString = `"true"`
	} else {
		boolString = `"false"`
	}

	return []byte(boolString), nil
}
