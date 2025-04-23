//
//
// File generated from our OpenAPI spec
//
//

// Package metereventstream provides the metereventstream related APIs
package metereventstream

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke metereventstream related APIs.
type Client struct {
	BMeterEvents stripe.Backend
	Key          string
}

// Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
func (c Client) New(params *stripe.V2BillingMeterEventStreamParams) error {
	metereventstream := &stripe.APIResource{}
	err := c.BMeterEvents.Call(
		http.MethodPost, "/v2/billing/meter_event_stream", c.Key, params, metereventstream)
	return err
}
