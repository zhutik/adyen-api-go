package adyen

import (
	"encoding/json"
	"testing"
)

type thing struct {
	Bool stringBool `json:"value"`
}

func TestStringBool_Unmarshal(t *testing.T) {
	cases := []struct {
		name   string
		json   string
		exp    bool
		expErr bool
	}{
		{
			name:   "empty",
			json:   `{ "value": "" }`,
			exp:    false,
			expErr: true,
		},
		{
			name: "true",
			json: `{ "value": "true" }`,
			exp:  true,
		},
		{
			name: "TRUE",
			json: `{ "value": "TRUE" }`,
			exp:  true,
		},
		{
			name: "1",
			json: `{ "value": "1" }`,
			exp:  true,
		},
		{
			name: "spaces",
			json: `{ "value": " true  " }`,
			exp:  true,
		},
		{
			name: "false",
			json: `{ "value": "false" }`,
			exp:  false,
		},
		{
			name: "FALSE",
			json: `{ "value": "FALSE" }`,
			exp:  false,
		},
		{
			name: "0",
			json: `{ "value": "0" }`,
			exp:  false,
		},
		{
			name: "spaces",
			json: `{ "value": " false  " }`,
			exp:  false,
		},
		{
			name: "spaces",
			json: `{ "value": " false  " }`,
			exp:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var th thing
			err := json.Unmarshal([]byte(c.json), &th)
			if !c.expErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if c.expErr && err == nil {
				t.Fatalf("expected error but didn't get one")
			}

			if stringBool(c.exp) != th.Bool {
				t.Fatalf("my exp: %v but got %v", c.exp, th.Bool)
			}
		})
	}
}
