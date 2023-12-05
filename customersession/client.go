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
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/customer_sessions", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, customersession)
	return customersession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
