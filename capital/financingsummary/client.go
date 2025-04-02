//
//
// File generated from our OpenAPI spec
//
//

// Package financingsummary provides the /v1/capital/financing_summary APIs
package financingsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/capital/financing_summary APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve the financing state for the account that was authenticated in the request.
func Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	return getC().Get(params)
}

// Retrieve the financing state for the account that was authenticated in the request.
func (c Client) Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	financingsummary := &stripe.CapitalFinancingSummary{}
	err := c.B.Call(
		http.MethodGet, "/v1/capital/financing_summary", c.Key, params, financingsummary)
	return financingsummary, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
