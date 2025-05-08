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

// v2BillingMeterEventService is used to invoke meterevent related APIs.
type v2BillingMeterEventService struct {
	B   Backend
	Key string
}

// Creates a meter event. Events are validated synchronously, but are processed asynchronously. Supports up to 1,000 events per second in livemode. For higher rate-limits, please use meter event streams instead.
func (c v2BillingMeterEventService) Create(ctx context.Context, params *V2BillingMeterEventCreateParams) (*V2BillingMeterEvent, error) {
	if params == nil {
		params = &V2BillingMeterEventCreateParams{}
	}
	params.Context = ctx
	meterevent := &V2BillingMeterEvent{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_events", c.Key, params, meterevent)
	return meterevent, err
}
