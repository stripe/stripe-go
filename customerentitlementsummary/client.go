//
//
// File generated from our OpenAPI spec
//
//

// Package customerentitlementsummary provides the /customers/{customer}/entitlement_summary APIs
package customerentitlementsummary

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /customers/{customer}/entitlement_summary APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a customer entitlement summary.
func Get(id string, params *stripe.CustomerEntitlementSummaryParams) (*stripe.CustomerEntitlementSummary, error) {
	return getC().Get(id, params)
}

// Get returns the details of a customer entitlement summary.
func (c Client) Get(id string, params *stripe.CustomerEntitlementSummaryParams) (*stripe.CustomerEntitlementSummary, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/entitlement_summary", id)
	customerentitlementsummary := &stripe.CustomerEntitlementSummary{}
	err := c.B.Call(
		http.MethodGet,
		path,
		c.Key,
		params,
		customerentitlementsummary,
	)
	return customerentitlementsummary, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
