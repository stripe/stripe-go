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

// v2DataAnalyticsMetricQueryService is used to invoke metricquery related APIs.
type v2DataAnalyticsMetricQueryService struct {
	B   Backend
	Key string
}

// Run a synchronous metric query against one or more metrics within the same metric namespace.
func (c v2DataAnalyticsMetricQueryService) Create(ctx context.Context, params *V2DataAnalyticsMetricQueryCreateParams) (*V2DataAnalyticsMetricQueryResult, error) {
	if params == nil {
		params = &V2DataAnalyticsMetricQueryCreateParams{}
	}
	params.Context = ctx
	metricqueryresult := &V2DataAnalyticsMetricQueryResult{}
	err := c.B.Call(
		http.MethodPost, "/v2/data/analytics/metric_query", c.Key, params, metricqueryresult)
	return metricqueryresult, err
}
