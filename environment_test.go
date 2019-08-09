package adyen

import "testing"

func TestTestEnvironment(t *testing.T) {
	act := TestEnvironment()
	equals(t, Testing.apiURL, act.apiURL)
	equals(t, Testing.clientURL, act.clientURL)
	equals(t, Testing.hppURL, act.hppURL)
}

func TestBaseURLEnvironmentTesting(t *testing.T) {
	env := TestEnvironment()
	act := env.BaseURL("service", "version")
	exp := "https://pal-test.adyen.com/pal/servlet/service/version"

	equals(t, exp, act)
}

func TestClientURLEnvironmentTesting(t *testing.T) {
	env := TestEnvironment()
	act := env.ClientURL("clientID")
	exp := "https://test.adyen.com/hpp/cse/js/clientID.shtml"

	equals(t, exp, act)
}

func TestHppURLEnvironmentTest(t *testing.T) {
	env := TestEnvironment()
	act := env.HppURL("request")
	exp := "https://test.adyen.com/hpp/request.shtml"

	equals(t, exp, act)
}

func TestCheckoutURLEnvironmentTesting(t *testing.T) {
	env := TestEnvironment()
	act := env.CheckoutURL("service", "version")
	exp := "https://checkout-test.adyen.com/services/PaymentSetupAndVerification/version/service"

	equals(t, exp, act)
}

func TestEnvironmentProductionValidation(t *testing.T) {
	cases := []struct {
		name        string
		random      string
		companyName string
	}{
		{
			name:        "missing random",
			random:      "",
			companyName: "AcmeAccount123",
		},
		{
			name:        "missing company name",
			random:      "5409c4fd1cc98a4e",
			companyName: "",
		},
		{
			name:        "missing random and company name",
			random:      "",
			companyName: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := ProductionEnvironment(c.random, c.companyName)
			equals(t, errProdEnvValidation, err)
		})
	}
}

func TestBaseURLEnvironmentProduction(t *testing.T) {
	env, err := ProductionEnvironment("5409c4fd1cc98a4e", "AcmeAccount123")
	if err != nil {
		t.Fatalf("error creating production environment: %v", err)
	}

	act := env.BaseURL("service", "version")
	exp := "https://5409c4fd1cc98a4e-AcmeAccount123pal-live.adyenpayments.com/pal/servlet/service/version"

	equals(t, exp, act)
}

func TestClientURLEnvironmentProduction(t *testing.T) {
	env, err := ProductionEnvironment("5409c4fd1cc98a4e", "AcmeAccount123")
	if err != nil {
		t.Fatalf("error creating production environment: %v", err)
	}

	act := env.ClientURL("clientID")
	exp := "https://live.adyen.com/hpp/cse/js/clientID.shtml"

	equals(t, exp, act)
}

func TestHppURLEnvironmentProduction(t *testing.T) {
	env, err := ProductionEnvironment("5409c4fd1cc98a4e", "AcmeAccount123")
	if err != nil {
		t.Fatalf("error creating production environment: %v", err)
	}

	act := env.HppURL("request")
	exp := "https://live.adyen.com/hpp/request.shtml"

	equals(t, exp, act)
}

func TestCheckoutURLEnvironmentProduction(t *testing.T) {
	env, err := ProductionEnvironment("5409c4fd1cc98a4e", "AcmeAccount123")
	if err != nil {
		t.Fatalf("error creating production environment: %v", err)
	}

	act := env.CheckoutURL("service", "version")
	exp := "https://5409c4fd1cc98a4e-AcmeAccount123checkout-live.adyenpayments.com/services/PaymentSetupAndVerification/version/service"

	equals(t, exp, act)
}
