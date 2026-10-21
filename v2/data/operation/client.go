//
//
// File generated from our OpenAPI spec
//
//

// Package operation provides the operation related APIs
package operation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke operation related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Resolve a natural-language query into ranked public analytical resources
// (schemas, reports, and metrics) with execution context.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.V2DataOperationSearchParams) (*stripe.V2DataSearchResultList, error) {
	searchresultlist := &stripe.V2DataSearchResultList{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/operations/search", c.Key, params, searchresultlist)
	return searchresultlist, err
}
