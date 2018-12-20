// Package checkoutsession provides the /checkout_sessions APIs
package checkoutsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /checkout_sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new session.
func New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().New(params)
}

// New creates a new session.
func (c Client) New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodPost, "/v1/checkout_sessions", c.Key, params, session)
	return session, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
