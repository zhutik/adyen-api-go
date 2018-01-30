package adyen

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Set environment variables for subsequent tests.
	if err := godotenv.Load(".default.env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	os.Exit(m.Run())
}

var noopLogger = log.New(ioutil.Discard, "", log.LstdFlags)

// getTestInstance - instanciate adyen for tests
func getTestInstance() *Adyen {
	instance := New(
		Testing,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		noopLogger)

	return instance
}

// getTestInstanceWithHPP - instanciate adyen for tests
func getTestInstanceWithHPP() *Adyen {
	instance := NewWithHMAC(
		Testing,
		os.Getenv("ADYEN_USERNAME"),
		os.Getenv("ADYEN_PASSWORD"),
		os.Getenv("ADYEN_HMAC"),
		noopLogger)

	return instance
}

// randInt - get random integer from a given range
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// randomString - generate randorm string of given length
// note: not for use in live code
func randomString(l int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
