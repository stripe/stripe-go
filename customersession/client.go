//
//
// File generated from our OpenAPI spec
//
//

// Package customersession provides the /customer_sessions APIs
package customersession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /customer_sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new customer session.
func New(params *stripe.CustomerSessionParams) (*stripe.CustomerSession, error) {
	return getC().New(params)
}

// New creates a new customer session.
func (c Client) New(params *stripe.CustomerSessionParams) (*stripe.CustomerSession, error) {
	customersession := &stripe.CustomerSession{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/customer_sessions",
		c.Key,
		params,
		customersession,
	)
	return customersession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
