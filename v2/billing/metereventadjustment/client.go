//
//
// File generated from our OpenAPI spec
//
//

// Package metereventadjustment provides the metereventadjustment related APIs
package metereventadjustment

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke metereventadjustment related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a meter event adjustment to cancel a previously sent meter event.
func (c Client) New(params *stripe.V2BillingMeterEventAdjustmentParams) (*stripe.V2BillingMeterEventAdjustment, error) {
	metereventadjustment := &stripe.V2BillingMeterEventAdjustment{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_adjustments", c.Key, params, metereventadjustment)
	return metereventadjustment, err
}
