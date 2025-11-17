//
//
// File generated from our OpenAPI spec
//
//

// Package accounttoken provides the accounttoken related APIs
package accounttoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke accounttoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an Account Token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountTokenParams) (*stripe.V2CoreAccountToken, error) {
	accounttoken := &stripe.V2CoreAccountToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_tokens", c.Key, params, accounttoken)
	return accounttoken, err
}

// Retrieves an Account Token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountTokenParams) (*stripe.V2CoreAccountToken, error) {
	path := stripe.FormatURLPath("/v2/core/account_tokens/%s", id)
	accounttoken := &stripe.V2CoreAccountToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accounttoken)
	return accounttoken, err
}
