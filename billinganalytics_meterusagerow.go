//
//
// File generated from our OpenAPI spec
//
//

package stripe

type BillingAnalyticsMeterUsageRow struct {
	// A set of key-value pairs representing the dimensions of the meter usage.
	Dimensions map[string]string `json:"dimensions"`
	// Timestamp indicating the end of the bucket. Measured in seconds since the Unix epoch.
	EndsAt int64 `json:"ends_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The unique identifier for the meter. Null if no meters were provided and usage was aggregated across all meters.
	Meter string `json:"meter"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Timestamp indicating the start of the bucket. Measured in seconds since the Unix epoch.
	StartsAt int64 `json:"starts_at"`
	// The aggregated meter usage value for the specified bucket.
	Value float64 `json:"value"`
}

// BillingAnalyticsMeterUsageRowList is a list of MeterUsageRows as retrieved from a list endpoint.
type BillingAnalyticsMeterUsageRowList struct {
	APIResource
	ListMeta
	Data []*BillingAnalyticsMeterUsageRow `json:"data"`
}
