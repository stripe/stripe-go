//
//
// File generated from our OpenAPI spec
//
//

// Package meterevent provides the /v1/billing/meter_events APIs
package meterevent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/billing/meter_events APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing meter event.
func New(params *stripe.BillingMeterEventParams) (*stripe.BillingMeterEvent, error) {
	return getC().New(params)
}

// Creates a billing meter event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.BillingMeterEventParams) (*stripe.BillingMeterEvent, error) {
	meterevent := &stripe.BillingMeterEvent{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/meter_events", c.Key, params, meterevent)
	return meterevent, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
