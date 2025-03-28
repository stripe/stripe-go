//
//
// File generated from our OpenAPI spec
//
//

// Package accountlink provides the accountlink related APIs
package accountlink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke accountlink related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
func (c Client) New(params *stripe.V2CoreAccountLinkParams) (*stripe.V2CoreAccountLink, error) {
	accountlink := &stripe.V2CoreAccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_links", c.Key, params, accountlink)
	return accountlink, err
}
