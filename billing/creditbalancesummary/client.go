//
//
// File generated from our OpenAPI spec
//
//

// Package creditbalancesummary provides the /v1/billing/credit_balance_summary APIs
package creditbalancesummary

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke /v1/billing/credit_balance_summary APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the credit balance summary for a customer.
func Get(params *stripe.BillingCreditBalanceSummaryParams) (*stripe.BillingCreditBalanceSummary, error) {
	return getC().Get(params)
}

// Retrieves the credit balance summary for a customer.
func (c Client) Get(params *stripe.BillingCreditBalanceSummaryParams) (*stripe.BillingCreditBalanceSummary, error) {
	creditbalancesummary := &stripe.BillingCreditBalanceSummary{}
	err := c.B.Call(
		http.MethodGet, "/v1/billing/credit_balance_summary", c.Key, params, creditbalancesummary)
	return creditbalancesummary, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
