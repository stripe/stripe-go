package stripe_test

import (
	"reflect"
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
	. "github.com/stripe/stripe-go/testing"
)

func TestCheckinRangeQueryParamsAppendTo(t *testing.T) {
	{
		values := &form.Values{}

		// Try it with an empty set of parameters
		params := &stripe.RangeQueryParams{}
		params.AppendTo(values, "created")

		if !values.Empty() {
			t.Fatalf("Expected request values to be empty")
		}
	}

	{
		values := &form.Values{}

		// Try it with an empty set of parameters
		params := &stripe.RangeQueryParams{GreaterThan: 99}
		params.AppendTo(values, "created")

		value := values.Get("created[gt]")
		if len(value) != 1 || value[0] != "99" {
			t.Fatalf(
				"Expected encoded value of 99 for created[gt] but got %v.",
				value)
		}
	}

	{
		values := &form.Values{}

		// Try it with an empty set of parameters
		params := &stripe.RangeQueryParams{GreaterThanOrEqual: 99}
		params.AppendTo(values, "created")

		value := values.Get("created[gte]")
		if len(value) != 1 || value[0] != "99" {
			t.Fatalf(
				"Expected encoded value of 99 for created[gte] but got %v.",
				value)
		}
	}

	{
		values := &form.Values{}

		// Try it with an empty set of parameters
		params := &stripe.RangeQueryParams{LesserThan: 99}
		params.AppendTo(values, "created")

		value := values.Get("created[lt]")
		if len(value) != 1 || value[0] != "99" {
			t.Fatalf(
				"Expected encoded value of 99 for created[lt] but got %v.",
				value)
		}
	}

	{
		values := &form.Values{}

		// Try it with an empty set of parameters
		params := &stripe.RangeQueryParams{LesserThanOrEqual: 99}
		params.AppendTo(values, "created")

		value := values.Get("created[lte]")
		if len(value) != 1 || value[0] != "99" {
			t.Fatalf(
				"Expected encoded value of 99 for created[lte] but got %v.",
				value)
		}
	}
}

func TestParamsWithExtras(t *testing.T) {
	testCases := []struct {
		InitialBody  [][2]string
		Extras       [][2]string
		ExpectedBody [][2]string
	}{
		{
			InitialBody:  [][2]string{{"foo", "bar"}},
			Extras:       [][2]string{},
			ExpectedBody: [][2]string{{"foo", "bar"}},
		},
		{
			InitialBody:  [][2]string{{"foo", "bar"}},
			Extras:       [][2]string{{"foo", "baz"}, {"other", "thing"}},
			ExpectedBody: [][2]string{{"foo", "bar"}, {"foo", "baz"}, {"other", "thing"}},
		},
	}

	for _, testCase := range testCases {
		p := stripe.Params{}

		for _, extra := range testCase.Extras {
			p.AddExtra(extra[0], extra[1])
		}

		body := valuesFromArray(testCase.InitialBody)
		p.AppendTo(body)

		expected := valuesFromArray(testCase.ExpectedBody)
		if !reflect.DeepEqual(body, expected) {
			t.Fatalf("Expected body of %v but got %v.", expected, body)
		}
	}
}

func TestCheckinListParamsExpansion(t *testing.T) {
	testCases := []struct {
		InitialBody  [][2]string
		Expand       []string
		ExpectedBody [][2]string
	}{
		{
			InitialBody:  [][2]string{{"foo", "bar"}},
			Expand:       []string{},
			ExpectedBody: [][2]string{{"foo", "bar"}},
		},
		{
			InitialBody:  [][2]string{{"foo", "bar"}, {"foo", "baz"}},
			Expand:       []string{"data", "data.foo"},
			ExpectedBody: [][2]string{{"foo", "bar"}, {"foo", "baz"}, {"expand[]", "data"}, {"expand[]", "data.foo"}},
		},
	}

	for _, testCase := range testCases {
		p := stripe.ListParams{}

		for _, exp := range testCase.Expand {
			p.Expand(exp)
		}

		body := valuesFromArray(testCase.InitialBody)
		p.AppendTo(body)

		expected := valuesFromArray(testCase.ExpectedBody)
		if !reflect.DeepEqual(body, expected) {
			t.Fatalf("Expected body of %v but got %v.", expected, body)
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

// Converts a collection of key/value tuples in a two dimensional slice/array
// into form.Values form. The purpose of this is that it's much cleaner to
// initialize the array all at once on a single line.
func valuesFromArray(arr [][2]string) *form.Values {
	body := &form.Values{}
	for _, v := range arr {
		body.Add(v[0], v[1])
	}
	return body
}
