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

// v1BillingAnalyticsMeterUsageService is used to invoke /v1/billing/analytics/meter_usage APIs.
type v1BillingAnalyticsMeterUsageService struct {
	B   Backend
	Key string
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
func (c v1BillingAnalyticsMeterUsageService) Retrieve(ctx context.Context, params *BillingAnalyticsMeterUsageRetrieveParams) (*BillingAnalyticsMeterUsage, error) {
	if params == nil {
		params = &BillingAnalyticsMeterUsageRetrieveParams{}
	}
	params.Context = ctx
	meterusage := &BillingAnalyticsMeterUsage{}
	err := c.B.Call(
		http.MethodGet, "/v1/billing/analytics/meter_usage", c.Key, params, meterusage)
	return meterusage, err
}
