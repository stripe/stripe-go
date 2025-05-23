//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2BillingMeterEventSessionService is used to invoke metereventsession related APIs.
type v2BillingMeterEventSessionService struct {
	B   Backend
	Key string
}

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
func (c v2BillingMeterEventSessionService) Create(ctx context.Context, params *V2BillingMeterEventSessionCreateParams) (*V2BillingMeterEventSession, error) {
	if params == nil {
		params = &V2BillingMeterEventSessionCreateParams{}
	}
	params.Context = ctx
	metereventsession := &V2BillingMeterEventSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_session", c.Key, params, metereventsession)
	return metereventsession, err
}
