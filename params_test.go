package stripe_test

import (
	"net/url"
	"reflect"
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/testing"
)

func TestRequestValues(t *testing.T) {
	values := &stripe.RequestValues{}

	actual := values.Encode()
	expected := ""
	if expected != actual {
		t.Fatalf("Expected encoded value of %v but got %v.", expected, actual)
	}

	values = &stripe.RequestValues{}
	values.Add("foo", "bar")

	actual = values.Encode()
	expected = "foo=bar"
	if expected != actual {
		t.Fatalf("Expected encoded value of %v but got %v.", expected, actual)
	}

	values = &stripe.RequestValues{}
	values.Add("foo", "bar")
	values.Add("foo", "bar")
	values.Add("baz", "bar")

	actual = values.Encode()
	expected = "foo=bar&foo=bar&baz=bar"
	if expected != actual {
		t.Fatalf("Expected encoded value of %v but got %v.", expected, actual)
	}

	values.Set("foo", "firstbar")

	actual = values.Encode()
	expected = "foo=firstbar&foo=bar&baz=bar"
	if expected != actual {
		t.Fatalf("Expected encoded value of %v but got %v.", expected, actual)
	}

	values.Set("new", "appended")

	actual = values.Encode()
	expected = "foo=firstbar&foo=bar&baz=bar&new=appended"
	if expected != actual {
		t.Fatalf("Expected encoded value of %v but got %v.", expected, actual)
	}
}

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

		body := fromURLValues(testCase.InitialBody)
		p.AppendTo(body)

		if !reflect.DeepEqual(body, fromURLValues(testCase.ExpectedBody)) {
			t.Fatalf("Expected body of %v but got %v.", testCase.ExpectedBody, body)
		}
	}
}

func fromURLValues(urlValues url.Values) *stripe.RequestValues {
	body := &stripe.RequestValues{}
	for k, values := range urlValues {
		for _, v := range values {
			body.Add(k, v)
		}
	}
	return body
}

func TestCheckinListParamsExpansion(t *testing.T) {
	testCases := []struct {
		InitialBody  url.Values
		Expand       []string
		ExpectedBody url.Values
	}{
		{
			InitialBody:  url.Values{"foo": {"bar"}},
			Expand:       []string{},
			ExpectedBody: url.Values{"foo": {"bar"}},
		},
		{
			InitialBody:  url.Values{"foo": {"bar", "baz"}},
			Expand:       []string{"data", "data.foo"},
			ExpectedBody: url.Values{"foo": {"bar", "baz"}, "expand[]": {"data", "data.foo"}},
		},
	}

	for _, testCase := range testCases {
		p := stripe.ListParams{}

		for _, exp := range testCase.Expand {
			p.Expand(exp)
		}

		body := fromURLValues(testCase.InitialBody)
		p.AppendTo(body)

		if !reflect.DeepEqual(body, fromURLValues(testCase.ExpectedBody)) {
			t.Fatalf("Expected body of %v but got %v.", testCase.ExpectedBody, body)
		}
	}
}

func TestCheckinListParamsToParams(t *testing.T) {
	listParams := &stripe.ListParams{StripeAccount: TestMerchantID}
	params := listParams.ToParams()

	if params.StripeAccount != TestMerchantID {
		t.Fatalf("Expected StripeAccount of %v but got %v.",
			TestMerchantID, params.StripeAccount)
	}
}

func TestCheckinParamsSetAccount(t *testing.T) {
	p := &stripe.Params{}
	p.SetAccount(TestMerchantID)

	if p.Account != TestMerchantID {
		t.Fatalf("Expected Account of %v but got %v.", TestMerchantID, p.Account)
	}

	if p.StripeAccount != TestMerchantID {
		t.Fatalf("Expected StripeAccount of %v but got %v.", TestMerchantID, p.StripeAccount)
	}
}

func TestCheckinParamsSetStripeAccount(t *testing.T) {
	p := &stripe.Params{}
	p.SetStripeAccount(TestMerchantID)

	if p.StripeAccount != TestMerchantID {
		t.Fatalf("Expected Account of %v but got %v.", TestMerchantID, p.StripeAccount)
	}

	// Check that we don't set the deprecated parameter.
	if p.Account != "" {
		t.Fatalf("Expected empty Account but got %v.", TestMerchantID)
	}
}
