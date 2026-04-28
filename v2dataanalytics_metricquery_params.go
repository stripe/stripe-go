//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A list of the metrics to be returned; all metrics must share the same metric namespace.
type V2DataAnalyticsMetricQueryMetricParams struct {
	// The Gen6 ID for this metric, e.g. metric_61Sud3n5oAGVCWiSr5.
	ID *string `form:"id" json:"id,omitempty"`
	// The common name for this metric, e.g. mrr_minor_units.
	Name *string `form:"name" json:"name,omitempty"`
}

// Run a synchronous metric query against one or more metrics within the same metric namespace.
type V2DataAnalyticsMetricQueryParams struct {
	Params `form:"*"`
	// Which currency to return monetary metric results in.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Timestamp denoting the end of the time range to request data for.
	EndsAt *time.Time `form:"ends_at" json:"ends_at"`
	// Which dimension values to filter on; keys are dimension names, values are arrays of dimension values to filter on.
	Filters map[string]any `form:"filters" json:"filters,omitempty"`
	// The time granularity to aggregate results by.
	Granularity *string `form:"granularity" json:"granularity"`
	// Which dimension keys to group by; if not specified no grouping is performed.
	GroupBy []*string `form:"group_by" json:"group_by,omitempty"`
	// The maximum number of rows in the response.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// A list of the metrics to be returned; all metrics must share the same metric namespace.
	Metrics []*V2DataAnalyticsMetricQueryMetricParams `form:"metrics" json:"metrics"`
	// Pagination future-proofing: page token for navigating to next/previous batch of rows.
	Page *string `form:"page" json:"page,omitempty"`
	// Timestamp denoting the beginning of the time range to request data for.
	StartsAt *time.Time `form:"starts_at" json:"starts_at"`
	// The timezone results should be in; defaults to the merchant's timezone.
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// A list of the metrics to be returned; all metrics must share the same metric namespace.
type V2DataAnalyticsMetricQueryCreateMetricParams struct {
	// The Gen6 ID for this metric, e.g. metric_61Sud3n5oAGVCWiSr5.
	ID *string `form:"id" json:"id,omitempty"`
	// The common name for this metric, e.g. mrr_minor_units.
	Name *string `form:"name" json:"name,omitempty"`
}

// Run a synchronous metric query against one or more metrics within the same metric namespace.
type V2DataAnalyticsMetricQueryCreateParams struct {
	Params `form:"*"`
	// Which currency to return monetary metric results in.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Timestamp denoting the end of the time range to request data for.
	EndsAt *time.Time `form:"ends_at" json:"ends_at"`
	// Which dimension values to filter on; keys are dimension names, values are arrays of dimension values to filter on.
	Filters map[string]any `form:"filters" json:"filters,omitempty"`
	// The time granularity to aggregate results by.
	Granularity *string `form:"granularity" json:"granularity"`
	// Which dimension keys to group by; if not specified no grouping is performed.
	GroupBy []*string `form:"group_by" json:"group_by,omitempty"`
	// The maximum number of rows in the response.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// A list of the metrics to be returned; all metrics must share the same metric namespace.
	Metrics []*V2DataAnalyticsMetricQueryCreateMetricParams `form:"metrics" json:"metrics"`
	// Pagination future-proofing: page token for navigating to next/previous batch of rows.
	Page *string `form:"page" json:"page,omitempty"`
	// Timestamp denoting the beginning of the time range to request data for.
	StartsAt *time.Time `form:"starts_at" json:"starts_at"`
	// The timezone results should be in; defaults to the merchant's timezone.
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}
