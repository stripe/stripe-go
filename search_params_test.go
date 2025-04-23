package stripe_test

import (
	"context"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/form"
	. "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

type testSearchParams struct {
	stripe.SearchParams `form:"*"`
	Page                *string `form:"page"`
}

func TestSearchParams_Nested(t *testing.T) {
	params := &testSearchParams{
		Page: stripe.String("page_value"),
		SearchParams: stripe.SearchParams{
			Query: "query_value",
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	assert.Equal(t, valuesFromArray([][2]string{
		{"query", "query_value"},
		{"page", "page_value"},
	}), body)
}

func TestSearchParams_Expand(t *testing.T) {
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
			ExpectedBody: [][2]string{{"foo", "bar"}, {"foo", "baz"}, {"expand[0]", "data"}, {"expand[1]", "data.foo"}},
		},
	}

	for _, testCase := range testCases {
		p := stripe.SearchParams{}

		for _, exp := range testCase.Expand {
			p.AddExpand(exp)
		}

		body := valuesFromArray(testCase.InitialBody)
		form.AppendTo(body, p)
		assert.Equal(t, valuesFromArray(testCase.ExpectedBody), body)
	}
}

func TestSearchParams_SetStripeAccount(t *testing.T) {
	p := &stripe.SearchParams{}
	p.SetStripeAccount(TestMerchantID)
	assert.Equal(t, TestMerchantID, *p.StripeAccount)
}

func TestSearchParams_ToParams(t *testing.T) {
	SearchParams := &stripe.SearchParams{
		Context: context.Background(),
	}
	SearchParams.SetStripeAccount(TestMerchantID)
	params := SearchParams.ToParams()
	assert.Equal(t, SearchParams.Context, params.Context)
	assert.Equal(t, *SearchParams.StripeAccount, *params.StripeAccount)
}
