package adyen

import (
	"math/rand"
	"os"
	"time"
)

// getTestInstance - instanciate adyen for tests
func getTestInstance() *Adyen {
	instance := New(
		Testing,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_CLIENT_TOKEN"),
		os.Getenv("ADYEN_ACCOUNT"),
	)

	return instance
}

// getTestInstanceWithHPP - instanciate adyen for tests
func getTestInstanceWithHPP() *Adyen {
	instance := NewWithHPP(
		Testing,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_CLIENT_TOKEN"),
		os.Getenv("ADYEN_ACCOUNT"),
		os.Getenv("ADYEN_HMAC"),
		os.Getenv("ADYEN_SKINCODE"),
		os.Getenv("ADYEN_SHOPPER_LOCALE"),
	)

	return instance
}

// randInt - get random integer from a given range
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// randomString - generate randorm string of given length
func randomString(l int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
