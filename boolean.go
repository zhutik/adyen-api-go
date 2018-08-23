package adyen

import (
	"strconv"
	"strings"
)

// stringBool allows us to unmarhsal Adyen Boolean values
// which appear as strings instead of bools.
type stringBool bool

func NewStringBool(b bool) *stringBool {
	sb := stringBool(b)
	return &sb
}

func (b *stringBool) UnmarshalJSON(data []byte) (err error) {
	str := strings.TrimFunc(strings.ToLower(string(data)), func(c rune) bool {
		return c == ' ' || c == '"'
	})

	parsed, err := strconv.ParseBool(str)
	if err != nil {
		return
	}

	*b = stringBool(parsed)
	return
}

func (b stringBool) MarshalJSON() ([]byte, error) {
	boolResult := bool(b)
	var boolString string

	if boolResult {
		boolString = `"true"`
	} else {
		boolString = `"false"`
	}

	return []byte(boolString), nil
}
