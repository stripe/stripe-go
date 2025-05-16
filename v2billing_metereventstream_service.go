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

// v2BillingMeterEventStreamService is used to invoke metereventstream related APIs.
type v2BillingMeterEventStreamService struct {
	BMeterEvents Backend
	Key          string
}

// Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
func (c v2BillingMeterEventStreamService) Create(ctx context.Context, params *V2BillingMeterEventStreamCreateParams) error {
	if params == nil {
		params = &V2BillingMeterEventStreamCreateParams{}
	}
	params.Context = ctx
	metereventstream := &APIResource{}
	err := c.BMeterEvents.Call(
		http.MethodPost, "/v2/billing/meter_event_stream", c.Key, params, metereventstream)
	return err
}
