//
//
// File generated from our OpenAPI spec
//
//

// Package accountlink provides the accountlink related APIs
package accountlink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke accountlink related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountLinkParams) (*stripe.V2CoreAccountLink, error) {
	accountlink := &stripe.V2CoreAccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_links", c.Key, params, accountlink)
	return accountlink, err
}
