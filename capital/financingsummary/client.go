//
//
// File generated from our OpenAPI spec
//
//

// Package financingsummary provides the /v1/capital/financing_summary APIs
package financingsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/capital/financing_summary APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve the financing summary object for the account.
func Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	return getC().Get(params)
}

// Retrieve the financing summary object for the account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	financingsummary := &stripe.CapitalFinancingSummary{}
	err := c.B.Call(
		http.MethodGet, "/v1/capital/financing_summary", c.Key, params, financingsummary)
	return financingsummary, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
