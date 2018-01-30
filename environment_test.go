package adyen

import "testing"

func TestBaseURL(t *testing.T) {
	act := Testing.BaseURL("service", "version")
	exp := "https://pal-test.adyen.com/pal/servlet/service/version"

	if act != exp {
		t.Fatalf("exp %q but got %q", exp, act)
	}
}

func TestClientURL(t *testing.T) {
	act := Testing.ClientURL("clientID")
	exp := "https://test.adyen.com/hpp/cse/js/clientID.shtml"

	if act != exp {
		t.Fatalf("exp %q but got %q", exp, act)
	}
}

func TestHppURL(t *testing.T) {
	act := Testing.HppURL("request")
	exp := "https://test.adyen.com/hpp/request.shtml"

	if act != exp {
		t.Fatalf("exp %q but got %q", exp, act)
	}
}
