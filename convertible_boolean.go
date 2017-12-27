package adyen

import (
	"fmt"
	"strings"
)

type convertibleBoolean bool

// UnmarshalJSON used to convert adyen JSON boolean to a real boolean values
func (bit *convertibleBoolean) UnmarshalJSON(data []byte) error {
	oldString := strings.ToLower(strings.TrimSpace(string(data)))
	asString := strings.Replace(oldString, "\"", "", -1)

	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return fmt.Errorf("boolean unmarshal error: invalid input %s", asString)
	}

	return nil
}
