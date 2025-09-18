//
//
// File generated from our OpenAPI spec
//
//

package stripe

// An array of meter parameters to specify which meters to include in the usage data. If not specified, usage across all meters for the customer is included.
type BillingAnalyticsMeterUsageMeterParams struct {
	// Key-value pairs used to filter usage events by meter dimension values. If specified, usage will be filtered for matching usage events.
	DimensionFilters map[string]string `form:"dimension_filters"`
	// List of meter dimension keys to group by. If specified, usage events will be grouped by the given meter dimension key's values.
	DimensionGroupByKeys []*string `form:"dimension_group_by_keys"`
	// Meter id to query usage for.
	MeterID *string `form:"meter_id"`
	// Key-value pairs used to filter usage events by high cardinality tenant dimension values. If specified, usage will be filtered for matching usage events.
	TenantFilters map[string]string `form:"tenant_filters"`
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
type BillingAnalyticsMeterUsageParams struct {
	Params `form:"*"`
	// The customer id to fetch meter usage data for.
	Customer *string `form:"customer"`
	// The timestamp from when to stop aggregating meter events (exclusive). Must be aligned with minute boundaries.
	EndTime *int64 `form:"end_time"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// An array of meter parameters to specify which meters to include in the usage data. If not specified, usage across all meters for the customer is included.
	Meters []*BillingAnalyticsMeterUsageMeterParams `form:"meters"`
	// The timestamp from when to start aggregating meter events (inclusive). Must be aligned with minute boundaries.
	StartTime *int64 `form:"start_time"`
	// The timezone to use for the start and end times. Defaults to UTC if not specified.
	Timezone *string `form:"timezone"`
	// Specifies what granularity to use when aggregating meter usage events. If not specified, a single event would be returned for the specified time range.
	ValueGroupingWindow *string `form:"value_grouping_window"`
}

// AddExpand appends a new field to expand.
func (p *BillingAnalyticsMeterUsageParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// An array of meter parameters to specify which meters to include in the usage data. If not specified, usage across all meters for the customer is included.
type BillingAnalyticsMeterUsageRetrieveMeterParams struct {
	// Key-value pairs used to filter usage events by meter dimension values. If specified, usage will be filtered for matching usage events.
	DimensionFilters map[string]string `form:"dimension_filters"`
	// List of meter dimension keys to group by. If specified, usage events will be grouped by the given meter dimension key's values.
	DimensionGroupByKeys []*string `form:"dimension_group_by_keys"`
	// Meter id to query usage for.
	MeterID *string `form:"meter_id"`
	// Key-value pairs used to filter usage events by high cardinality tenant dimension values. If specified, usage will be filtered for matching usage events.
	TenantFilters map[string]string `form:"tenant_filters"`
}

// Returns aggregated meter usage data for a customer within a specified time interval. The data can be grouped by various dimensions and can include multiple meters if specified.
type BillingAnalyticsMeterUsageRetrieveParams struct {
	Params `form:"*"`
	// The customer id to fetch meter usage data for.
	Customer *string `form:"customer"`
	// The timestamp from when to stop aggregating meter events (exclusive). Must be aligned with minute boundaries.
	EndTime *int64 `form:"end_time"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// An array of meter parameters to specify which meters to include in the usage data. If not specified, usage across all meters for the customer is included.
	Meters []*BillingAnalyticsMeterUsageRetrieveMeterParams `form:"meters"`
	// The timestamp from when to start aggregating meter events (inclusive). Must be aligned with minute boundaries.
	StartTime *int64 `form:"start_time"`
	// The timezone to use for the start and end times. Defaults to UTC if not specified.
	Timezone *string `form:"timezone"`
	// Specifies what granularity to use when aggregating meter usage events. If not specified, a single event would be returned for the specified time range.
	ValueGroupingWindow *string `form:"value_grouping_window"`
}

// AddExpand appends a new field to expand.
func (p *BillingAnalyticsMeterUsageRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A billing meter usage event represents an aggregated view of a customer's billing meter events within a specified timeframe.
type BillingAnalyticsMeterUsage struct {
	APIResource
	// The aggregated meter usage data for the specified customer and time range.
	Data []*BillingAnalyticsMeterUsageRow `json:"data"`
	// Timestamp indicating how fresh the data is. Measured in seconds since the Unix epoch.
	DataRefreshedAt int64 `json:"data_refreshed_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
