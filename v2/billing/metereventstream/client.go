//
//
// File generated from our OpenAPI spec
//
//

// Package metereventstream provides the metereventstream related APIs
package metereventstream

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke metereventstream related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	BMeterEvents stripe.Backend
	Key          string
}

// Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingMeterEventStreamParams) error {
	metereventstream := &stripe.APIResource{}
	err := c.BMeterEvents.Call(
		http.MethodPost, "/v2/billing/meter_event_stream", c.Key, params, metereventstream)
	return err
}
