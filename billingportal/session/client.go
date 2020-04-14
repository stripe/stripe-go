// Package session provides API functions related to billing portal sessions.
package session

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
)

// Client is used to invoke /billing_portal/sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new session.
func New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	return getC().New(params)
}

// New creates a new session.
func (c Client) New(params *stripe.BillingPortalSessionParams) (*stripe.BillingPortalSession, error) {
	session := &stripe.BillingPortalSession{}
	err := c.B.Call(http.MethodPost, "/v1/billing_portal/sessions", c.Key, params, session)
	return session, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
