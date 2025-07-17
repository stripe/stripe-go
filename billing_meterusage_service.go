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

// v1BillingMeterUsageService is used to invoke /v1/billing/analytics/meter_usage APIs.
type v1BillingMeterUsageService struct {
	B   Backend
	Key string
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
func (c v1BillingMeterUsageService) Retrieve(ctx context.Context, params *BillingMeterUsageRetrieveParams) (*BillingMeterUsage, error) {
	if params == nil {
		params = &BillingMeterUsageRetrieveParams{}
	}
	params.Context = ctx
	meterusage := &BillingMeterUsage{}
	err := c.B.Call(
		http.MethodGet, "/v1/billing/analytics/meter_usage", c.Key, params, meterusage)
	return meterusage, err
}
