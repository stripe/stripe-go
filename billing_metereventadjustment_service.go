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

// v1BillingMeterEventAdjustmentService is used to invoke /v1/billing/meter_event_adjustments APIs.
type v1BillingMeterEventAdjustmentService struct {
	B   Backend
	Key string
}

// Creates a billing meter event adjustment.
func (c v1BillingMeterEventAdjustmentService) Create(ctx context.Context, params *BillingMeterEventAdjustmentCreateParams) (*BillingMeterEventAdjustment, error) {
	if params == nil {
		params = &BillingMeterEventAdjustmentCreateParams{}
	}
	params.Context = ctx
	metereventadjustment := &BillingMeterEventAdjustment{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/meter_event_adjustments", c.Key, params, metereventadjustment)
	return metereventadjustment, err
}
