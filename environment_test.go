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

func TestParseEnvironment(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		expName string
		expErr  bool
	}{
		{
			name:    EnvironmentTesting,
			input:   EnvironmentTesting,
			expName: EnvironmentTesting,
		},
		{
			name:    EnvironmentProduction,
			input:   EnvironmentProduction,
			expName: EnvironmentProduction,
		},
		{
			name:    "invalid",
			input:   "blah",
			expName: "",
			expErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e, err := ParseEnvironment(c.input)

			if c.expErr {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}
				equals(t, errParseEnvironment{name: c.input}, err)
				return
			}

			equals(t, c.expName, e.Name)
		})
	}
}
