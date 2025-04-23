//
//
// File generated from our OpenAPI spec
//
//

// Package metereventsession provides the metereventsession related APIs
package metereventsession

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke metereventsession related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
func (c Client) New(params *stripe.V2BillingMeterEventSessionParams) (*stripe.V2BillingMeterEventSession, error) {
	metereventsession := &stripe.V2BillingMeterEventSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_session", c.Key, params, metereventsession)
	return metereventsession, err
}
