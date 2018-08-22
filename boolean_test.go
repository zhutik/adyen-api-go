package adyen

import (
	"encoding/json"
	"testing"
)

type thing struct {
	Bool stringBool `json:"value"`
}

type thingWithEmpty struct {
	Bool *stringBool `json:"value,omitempty"`
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

func TestStringBool_Marshal(t *testing.T) {
	cases := []struct {
		name     string
		object   thing
		expected string
	}{
		{
			name:     "true",
			object:   thing{Bool: stringBool(true)},
			expected: `{"value":"true"}`,
		},
		{
			name:     "false",
			object:   thing{Bool: stringBool(false)},
			expected: `{"value":"false"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, _ := json.Marshal(c.object)
			str := string(res)

			if c.expected != str {
				t.Fatalf("my exp: %v but got %v", c.expected, str)
			}
		})
	}
}

func TestStringBool_MarshalWithOmitempty(t *testing.T) {
	valueTrue := stringBool(true)
	valueFalse := stringBool(false)

	cases := []struct {
		name     string
		object   thingWithEmpty
		expected string
	}{
		{
			name:     "true",
			object:   thingWithEmpty{Bool: &valueTrue},
			expected: `{"value":"true"}`,
		},
		{
			name:     "false",
			object:   thingWithEmpty{Bool: &valueFalse},
			expected: `{"value":"false"}`,
		},
		{
			name:     "false",
			object:   thingWithEmpty{},
			expected: `{}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, _ := json.Marshal(c.object)
			str := string(res)

			if c.expected != str {
				t.Fatalf("my exp: %v but got %v", c.expected, str)
			}
		})
	}
}
