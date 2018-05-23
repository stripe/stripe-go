package stripe_test

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
	. "github.com/stripe/stripe-go/testing"
)

func TestRangeQueryParamsAppendTo(t *testing.T) {
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
			EndingBefore:  "acct_123",
			Limit:         100,
			StartingAfter: "acct_123",
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	assert.Equal(t, valuesFromArray([][2]string{
		{"ending_before", "acct_123"},
		{"limit", "100"},
		{"starting_after", "acct_123"},
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
		p := &testParams{}

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

// AppendTo is implemented for testParams so that we can verify that Params is
// encoded properly even in the case where a leaf struct does a custom
// override.
func (p *testParams) AppendTo(body *form.Values, keyParts []string) {
}

type testSubParams struct {
	stripe.Params `form:"*"`
	SubField      string `form:"sub_field"`
}

func TestParams_AppendTo_Nested(t *testing.T) {
	params := &testParams{
		Field: "field_value",
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
		SubParams: &testSubParams{
			Params: stripe.Params{
				Metadata: map[string]string{
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

func TestListParams_Filters(t *testing.T) {
	p := &testListParams{}
	p.Filters.AddFilter("created", "gt", "123")

	body := &form.Values{}
	form.AppendTo(body, p)

	assert.Equal(t, valuesFromArray([][2]string{
		{"created[gt]", "123"},
	}), body)
}

func TestListParams_Expand(t *testing.T) {
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
			p.AddExpand(exp)
		}

		body := valuesFromArray(testCase.InitialBody)
		form.AppendTo(body, p)
		assert.Equal(t, valuesFromArray(testCase.ExpectedBody), body)
	}
}

func TestListParams_ToParams(t *testing.T) {
	listParams := &stripe.ListParams{
		Context:       context.Background(),
		StripeAccount: TestMerchantID,
	}
	params := listParams.ToParams()
	assert.Equal(t, listParams.Context, params.Context)
	assert.Equal(t, listParams.StripeAccount, params.StripeAccount)
}

func TestParams_SetStripeAccount(t *testing.T) {
	p := &stripe.Params{}
	p.SetStripeAccount(TestMerchantID)

	if p.StripeAccount != TestMerchantID {
		t.Fatalf("Expected Account of %v but got %v.", TestMerchantID, p.StripeAccount)
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
