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

// v2BillingMeterEventAdjustmentService is used to invoke metereventadjustment related APIs.
type v2BillingMeterEventAdjustmentService struct {
	B   Backend
	Key string
}

// Creates a meter event adjustment to cancel a previously sent meter event.
func (c v2BillingMeterEventAdjustmentService) Create(ctx context.Context, params *V2BillingMeterEventAdjustmentCreateParams) (*V2BillingMeterEventAdjustment, error) {
	if params == nil {
		params = &V2BillingMeterEventAdjustmentCreateParams{}
	}
	params.Context = ctx
	metereventadjustment := &V2BillingMeterEventAdjustment{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/meter_event_adjustments", c.Key, params, metereventadjustment)
	return metereventadjustment, err
}
