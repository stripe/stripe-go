//
//
// File generated from our OpenAPI spec
//
//

// Package meterevent provides the meterevent related APIs
package meterevent

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke meterevent related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a meter event. Events are validated synchronously, but are processed asynchronously. Supports up to 1,000 events per second in livemode. For higher rate-limits, please use meter event streams instead.
func (c Client) New(params *stripe.V2BillingMeterEventParams) (*stripe.V2BillingMeterEvent, error) {
	meterevent := &stripe.V2BillingMeterEvent{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_events", c.Key, params, meterevent)
	return meterevent, err
}
