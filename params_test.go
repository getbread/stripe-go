package stripe_test

import (
	"net/url"
	"reflect"
	"testing"

	stripe "github.com/getbread/stripe-go"
)

func TestParamsWithExtras(t *testing.T) {
	testCases := []struct {
		InitialBody  url.Values
		Extras       [][2]string
		ExpectedBody url.Values
	}{
		{
			InitialBody:  url.Values{"foo": {"bar"}},
			Extras:       [][2]string{},
			ExpectedBody: url.Values{"foo": {"bar"}},
		},
		{
			InitialBody:  url.Values{"foo": {"bar"}},
			Extras:       [][2]string{{"foo", "baz"}, {"other", "thing"}},
			ExpectedBody: url.Values{"foo": {"bar", "baz"}, "other": {"thing"}},
		},
	}

	for _, testCase := range testCases {
		p := stripe.Params{}

		for _, extra := range testCase.Extras {
			p.AddExtra(extra[0], extra[1])
		}

		body := testCase.InitialBody
		p.AppendTo(&body)

		if !reflect.DeepEqual(body, testCase.ExpectedBody) {
			t.Fatalf("Expected body of %v but got %v.", testCase.ExpectedBody, body)
		}
	}
}
