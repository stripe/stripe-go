//
//
// File generated from our OpenAPI spec
//
//

// Package metereventadjustment provides the /v1/billing/meter_event_adjustments APIs
package metereventadjustment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/billing/meter_event_adjustments APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a billing meter event adjustment.
func New(params *stripe.BillingMeterEventAdjustmentParams) (*stripe.BillingMeterEventAdjustment, error) {
	return getC().New(params)
}

// Creates a billing meter event adjustment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.BillingMeterEventAdjustmentParams) (*stripe.BillingMeterEventAdjustment, error) {
	metereventadjustment := &stripe.BillingMeterEventAdjustment{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/meter_event_adjustments", c.Key, params, metereventadjustment)
	return metereventadjustment, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
