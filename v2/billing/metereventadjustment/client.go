//
//
// File generated from our OpenAPI spec
//
//

// Package metereventadjustment provides the metereventadjustment related APIs
package metereventadjustment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke metereventadjustment related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a meter event adjustment to cancel a previously sent meter event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingMeterEventAdjustmentParams) (*stripe.V2BillingMeterEventAdjustment, error) {
	metereventadjustment := &stripe.V2BillingMeterEventAdjustment{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_adjustments", c.Key, params, metereventadjustment)
	return metereventadjustment, err
}
