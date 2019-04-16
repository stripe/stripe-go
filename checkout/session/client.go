// Package session provides API functions related to checkout sessions.
package session

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
	err := c.B.Call(http.MethodPost, "/v1/checkout/sessions", c.Key, params, session)
	return session, err
}

// Get retrieves a session.
func Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().Get(id, params)
}

// Get retrieves a session.
func (c Client) Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	path := stripe.FormatURLPath("/v1/checkout/sessions/%s", id)
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
