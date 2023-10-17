//
//
// File generated from our OpenAPI spec
//
//

// Package financingsummary provides the /capital/financing_summary APIs
package financingsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /capital/financing_summary APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a capital financing summary.
func Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	return getC().Get(params)
}

// Get returns the details of a capital financing summary.
func (c Client) Get(params *stripe.CapitalFinancingSummaryParams) (*stripe.CapitalFinancingSummary, error) {
	financingsummary := &stripe.CapitalFinancingSummary{}
	err := c.B.Call(
		http.MethodGet,
		"/v1/capital/financing_summary",
		c.Key,
		params,
		financingsummary,
	)
	return financingsummary, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
