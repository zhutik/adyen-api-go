package adyen

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

// noopLogger keeps the logs quite during verbose testing.
var noopLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func TestMain(m *testing.M) {
	// Set environment variables for subsequent tests.
	if err := godotenv.Load(".default.env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	os.Exit(m.Run())
}

func TestNewWithTimeout(t *testing.T) {
	const timeout = time.Second * 123

	act := New(Testing, "un", "pw", nil, WithTimeout(timeout))
	equals(t, timeout, act.ClientTimeout)
}

func TestNewWithCurrency(t *testing.T) {
	const currency = "USD"

	act := New(Testing, "un", "pw", nil, WithCurrency(currency))
	equals(t, currency, act.Currency)
}

func TestNewWithCustomOptions(t *testing.T) {
	const merchant, currency, timeout = "merch", "JPY", time.Second * 21

	f1 := func(a *Adyen) {
		a.Currency = currency
		a.ClientTimeout = timeout
	}

	f2 := func(a *Adyen) {
		a.MerchantAccount = merchant
	}

	act := New(Testing, "un", "pw", nil, f1, f2)
	equals(t, merchant, act.MerchantAccount)
	equals(t, currency, act.Currency)
	equals(t, timeout, act.ClientTimeout)
}

func equals(tb testing.TB, exp interface{}, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\n\texp: %[1]v (%[1]T)\n\tgot: %[2]v (%[2]T)\n", exp, act)
	}
}

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
