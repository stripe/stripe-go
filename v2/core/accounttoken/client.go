//
//
// File generated from our OpenAPI spec
//
//

// Package accounttoken provides the accounttoken related APIs
package accounttoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke accounttoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an account token with a publishable key and pass it to the Accounts v2 API to
// create or update an account without its data touching your server.
// Learn more about [account tokens](https://docs.stripe.com/connect/account-tokens).
// In live mode, you can only create account tokens with your application's publishable key.
// In test mode, you can create account tokens with your secret key or publishable key.
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
