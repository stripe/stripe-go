//
//
// File generated from our OpenAPI spec
//
//

// Package claimablesandbox provides the claimablesandbox related APIs
package claimablesandbox

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke claimablesandbox related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreClaimableSandboxParams) (*stripe.V2CoreClaimableSandbox, error) {
	claimablesandbox := &stripe.V2CoreClaimableSandbox{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/claimable_sandboxes", c.Key, params, claimablesandbox)
	return claimablesandbox, err
}

// Retrieves the details of a claimable sandbox that was previously been created.
// Supply the unique claimable sandbox ID that was returned from your creation request,
// and Stripe will return the corresponding sandbox information.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreClaimableSandboxParams) (*stripe.V2CoreClaimableSandbox, error) {
	path := stripe.FormatURLPath("/v2/core/claimable_sandboxes/%s", id)
	claimablesandbox := &stripe.V2CoreClaimableSandbox{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, claimablesandbox)
	return claimablesandbox, err
}
