//
//
// File generated from our OpenAPI spec
//
//

// Package event provides the /entitlements/events APIs
package event

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /entitlements/events APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new entitlements event.
func New(params *stripe.EntitlementsEventParams) (*stripe.EntitlementsEvent, error) {
	return getC().New(params)
}

// New creates a new entitlements event.
func (c Client) New(params *stripe.EntitlementsEventParams) (*stripe.EntitlementsEvent, error) {
	event := &stripe.EntitlementsEvent{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/entitlements/events",
		c.Key,
		params,
		event,
	)
	return event, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
