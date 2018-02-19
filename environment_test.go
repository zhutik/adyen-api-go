package adyen

import "testing"

func TestBaseURLEnvironmentTesting(t *testing.T) {
	act := Testing.BaseURL("service", "version")
	exp := "https://pal-test.adyen.com/pal/servlet/service/version"

	equals(t, exp, act)
}

func TestClientURLEnvironmentTesting(t *testing.T) {
	act := Testing.ClientURL("clientID")
	exp := "https://test.adyen.com/hpp/cse/js/clientID.shtml"

	equals(t, exp, act)
}

func TestHppURLEnvironmentTest(t *testing.T) {
	act := Testing.HppURL("request")
	exp := "https://test.adyen.com/hpp/request.shtml"

	equals(t, exp, act)
}

func TestBaseURLEnvironmentProduction(t *testing.T) {
	env, _ := ParseEnvironment(EnvironmentProduction, "5409c4fd1cc98a4e", "AcmeAccount123")
	act := env.BaseURL("service", "version")
	exp := "https://5409c4fd1cc98a4e-AcmeAccount123-pal-live.adyen.com/pal/servlet/service/version"

	equals(t, exp, act)
}

func TestClientURLEnvironmentProduction(t *testing.T) {
	env, _ := ParseEnvironment(EnvironmentProduction, "5409c4fd1cc98a4e", "AcmeAccount123")
	act := env.ClientURL("clientID")
	exp := "https://live.adyen.com/hpp/cse/js/clientID.shtml"

	equals(t, exp, act)
}

func TestHppURLEnvironmentProduction(t *testing.T) {
	env, _ := ParseEnvironment(EnvironmentProduction, "5409c4fd1cc98a4e", "AcmeAccount123")
	act := env.HppURL("request")
	exp := "https://live.adyen.com/hpp/request.shtml"

	equals(t, exp, act)
}

func TestParseEnvironment(t *testing.T) {
	cases := []struct {
		name        string
		input       string
		random      string
		companyName string
		expName     string
		expErr      error
	}{
		{
			name:    EnvironmentTesting,
			input:   EnvironmentTesting,
			expName: EnvironmentTesting,
		},
		{
			name:        EnvironmentProduction,
			input:       EnvironmentProduction,
			random:      "5409c4fd1cc98a4e",
			companyName: "AcmeAccount123",
			expName:     EnvironmentProduction,
			expErr:      nil,
		},
		{
			name:        "production without random",
			input:       EnvironmentProduction,
			companyName: "AcmeAccount123",
			expName:     "",
			expErr:      errProdEnvValidation,
		},
		{
			name:    "production without company name",
			input:   EnvironmentProduction,
			random:  "5409c4fd1cc98a4e",
			expName: "",
			expErr:  errProdEnvValidation,
		},
		{
			name:    "invalid",
			input:   "blah",
			expName: "",
			expErr:  errParseEnvironment{name: "blah"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e, err := ParseEnvironment(c.input, c.random, c.companyName)
			if c.expErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if c.expErr != nil {
				if err == nil {
					t.Fatal("expected error but didn't get one")
				}
				equals(t, c.expErr, err)
				return
			}

			equals(t, c.expName, e.Name)
		})
	}
}
