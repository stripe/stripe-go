//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Array of metric values returned from this query.
type V2DataAnalyticsMetricQueryResultDataResult struct {
	// If this is a monetary metric, the currency it is returned in. Otherwise null.
	Currency Currency `json:"currency,omitempty"`
	// The Gen6 ID of this metric.
	Metric string `json:"metric"`
	// The common name of this metric.
	Name string `json:"name"`
	// The numeric value of this metric.
	Value int64 `json:"value,string"`
}

// An array of timeseries data rows.
type V2DataAnalyticsMetricQueryResultData struct {
	// A hash of dimension type to dimension instance, if group_by was specified.
	Dimensions map[string]string `json:"dimensions"`
	// A unique identifier for this row.
	ID string `json:"id"`
	// Array of metric values returned from this query.
	Results []*V2DataAnalyticsMetricQueryResultDataResult `json:"results"`
	// Timestamp denoting the start of this time bucket.
	Timestamp time.Time `json:"timestamp"`
}

// The result of a metric query.
type V2DataAnalyticsMetricQueryResult struct {
	APIResource
	// The timestamp at which this metric query result was created.
	Created time.Time `json:"created"`
	// An array of timeseries data rows.
	Data []*V2DataAnalyticsMetricQueryResultData `json:"data"`
	// The unique identifier of this metric query result.
	ID string `json:"id"`
	// Whether this query was run in live mode.
	Livemode bool `json:"livemode"`
	// Pagination future-proofing: URL to fetch the next page; always null for now.
	NextPageURL string `json:"next_page_url,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Pagination future-proofing: URL to fetch the previous page; always null for now.
	PreviousPageURL string `json:"previous_page_url,omitempty"`
	// A timestamp representing the freshness of the data this metric is aggregated from.
	RefreshedAt time.Time `json:"refreshed_at"`
}
