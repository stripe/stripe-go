package stripe_test

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
	. "github.com/stripe/stripe-go/testing"
)

func TestCheckinRangeQueryParamsAppendTo(t *testing.T) {
	type testParams struct {
		CreatedRange *stripe.RangeQueryParams `form:"created"`
	}

	{
		body := &form.Values{}

		// Try it with an empty set of parameters
		params := testParams{
			CreatedRange: &stripe.RangeQueryParams{},
		}
		form.AppendTo(body, params)
		assert.True(t, body.Empty())
	}

	{
		body := &form.Values{}

		params := testParams{
			CreatedRange: &stripe.RangeQueryParams{GreaterThan: 99},
		}
		form.AppendTo(body, params)
		assert.Equal(t, []string{"99"}, body.Get("created[gt]"))
	}

	{
		body := &form.Values{}

		params := testParams{
			CreatedRange: &stripe.RangeQueryParams{GreaterThanOrEqual: 99},
		}
		form.AppendTo(body, params)
		assert.Equal(t, []string{"99"}, body.Get("created[gte]"))
	}

	{
		body := &form.Values{}

		params := testParams{
			CreatedRange: &stripe.RangeQueryParams{LesserThan: 99},
		}
		form.AppendTo(body, params)
		assert.Equal(t, []string{"99"}, body.Get("created[lt]"))
	}

	{
		body := &form.Values{}

		params := testParams{
			CreatedRange: &stripe.RangeQueryParams{LesserThanOrEqual: 99},
		}
		form.AppendTo(body, params)
		assert.Equal(t, []string{"99"}, body.Get("created[lte]"))
	}
}

type testListParams struct {
	stripe.ListParams `form:"*"`
	Field             string `form:"field"`
}

func TestListParams_Nested(t *testing.T) {
	params := &testListParams{
		Field: "field_value",
		ListParams: stripe.ListParams{
			End:   "acct_123",
			Start: "acct_123",
			Limit: 100,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	assert.Equal(t, valuesFromArray([][2]string{
		{"starting_after", "acct_123"},
		{"ending_before", "acct_123"},
		{"limit", "100"},
		{"field", "field_value"},
	}), body)
}

func TestParams_AppendTo_Extra(t *testing.T) {
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
			Extras:       [][2]string{{"foo", "baz"}},
			ExpectedBody: [][2]string{{"foo", "bar"}, {"foo", "baz"}},
		},
	}

	for _, testCase := range testCases {
		p := &stripe.Params{}

		for _, extra := range testCase.Extras {
			p.AddExtra(extra[0], extra[1])
		}

		body := valuesFromArray(testCase.InitialBody)
		form.AppendTo(body, p)
		assert.Equal(t, valuesFromArray(testCase.ExpectedBody), body)
	}
}

type testParams struct {
	stripe.Params `form:"*"`
	Field         string         `form:"field"`
	SubParams     *testSubParams `form:"sub_params"`
}

type testSubParams struct {
	stripe.Params `form:"*"`
	SubField      string `form:"sub_field"`
}

func TestParams_AppendTo_Nested(t *testing.T) {
	params := &testParams{
		Field: "field_value",
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
		SubParams: &testSubParams{
			Params: stripe.Params{
				Meta: map[string]string{
					"sub_foo": "bar",
				},
			},
			SubField: "sub_field_value",
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	assert.Equal(t, valuesFromArray([][2]string{
		{"metadata[foo]", "bar"},
		{"field", "field_value"},
		{"sub_params[metadata][sub_foo]", "bar"},
		{"sub_params[sub_field]", "sub_field_value"},
	}), body)
}

func TestCheckinListParams_Expand(t *testing.T) {
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
		form.AppendTo(body, p)
		assert.Equal(t, valuesFromArray(testCase.ExpectedBody), body)
	}
}

func TestCheckinListParams_ToParams(t *testing.T) {
	listParams := &stripe.ListParams{StripeAccount: TestMerchantID}
	params := listParams.ToParams()

	if params.StripeAccount != TestMerchantID {
		t.Fatalf("Expected StripeAccount of %v but got %v.",
			TestMerchantID, params.StripeAccount)
	}
}

func TestCheckinParams_SetAccount(t *testing.T) {
	p := &stripe.Params{}
	p.SetAccount(TestMerchantID)

	if p.Account != TestMerchantID {
		t.Fatalf("Expected Account of %v but got %v.", TestMerchantID, p.Account)
	}

	if p.StripeAccount != TestMerchantID {
		t.Fatalf("Expected StripeAccount of %v but got %v.", TestMerchantID, p.StripeAccount)
	}
}

func TestCheckinParams_SetStripeAccount(t *testing.T) {
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

//
// ---
//

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
