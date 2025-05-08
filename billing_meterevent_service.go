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

// v1BillingMeterEventService is used to invoke /v1/billing/meter_events APIs.
type v1BillingMeterEventService struct {
	B   Backend
	Key string
}

// Creates a billing meter event.
func (c v1BillingMeterEventService) Create(ctx context.Context, params *BillingMeterEventCreateParams) (*BillingMeterEvent, error) {
	if params == nil {
		params = &BillingMeterEventCreateParams{}
	}
	params.Context = ctx
	meterevent := &BillingMeterEvent{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/meter_events", c.Key, params, meterevent)
	return meterevent, err
}
