//
//
// File generated from our OpenAPI spec
//
//

// Package issuedtoken provides the /v1/shared_payment/issued_tokens APIs
package issuedtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/shared_payment/issued_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new SharedPaymentIssuedToken object
func New(params *stripe.SharedPaymentIssuedTokenParams) (*stripe.SharedPaymentIssuedToken, error) {
	return getC().New(params)
}

// Creates a new SharedPaymentIssuedToken object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.SharedPaymentIssuedTokenParams) (*stripe.SharedPaymentIssuedToken, error) {
	issuedtoken := &stripe.SharedPaymentIssuedToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/shared_payment/issued_tokens", c.Key, params, issuedtoken)
	return issuedtoken, err
}

// Retrieves an existing SharedPaymentIssuedToken object
func Get(id string, params *stripe.SharedPaymentIssuedTokenParams) (*stripe.SharedPaymentIssuedToken, error) {
	return getC().Get(id, params)
}

// Retrieves an existing SharedPaymentIssuedToken object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.SharedPaymentIssuedTokenParams) (*stripe.SharedPaymentIssuedToken, error) {
	path := stripe.FormatURLPath("/v1/shared_payment/issued_tokens/%s", id)
	issuedtoken := &stripe.SharedPaymentIssuedToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, issuedtoken)
	return issuedtoken, err
}

// Revokes a SharedPaymentIssuedToken
func Revoke(id string, params *stripe.SharedPaymentIssuedTokenRevokeParams) (*stripe.SharedPaymentIssuedToken, error) {
	return getC().Revoke(id, params)
}

// Revokes a SharedPaymentIssuedToken
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Revoke(id string, params *stripe.SharedPaymentIssuedTokenRevokeParams) (*stripe.SharedPaymentIssuedToken, error) {
	path := stripe.FormatURLPath("/v1/shared_payment/issued_tokens/%s/revoke", id)
	issuedtoken := &stripe.SharedPaymentIssuedToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, issuedtoken)
	return issuedtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
