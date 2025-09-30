//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /v1/billing_portal/sessions APIs
package session

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/billing_portal/sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a session of the customer portal.
func New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	return getC().New(params)
}

// Creates a session of the customer portal.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	session := &stripe.BillingPortalSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing_portal/sessions", c.Key, params, session)
	return session, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
