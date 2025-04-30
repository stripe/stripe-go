//
//
// File generated from our OpenAPI spec
//
//

// Package metereventsession provides the metereventsession related APIs
package metereventsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke metereventsession related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingMeterEventSessionParams) (*stripe.V2BillingMeterEventSession, error) {
	metereventsession := &stripe.V2BillingMeterEventSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_session", c.Key, params, metereventsession)
	return metereventsession, err
}
