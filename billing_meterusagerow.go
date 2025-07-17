//
//
// File generated from our OpenAPI spec
//
//

package stripe

type BillingMeterUsageRow struct {
	// Timestamp indicating the end of the bucket. Measured in seconds since the Unix epoch.
	BucketEndTime int64 `json:"bucket_end_time"`
	// Timestamp indicating the start of the bucket. Measured in seconds since the Unix epoch.
	BucketStartTime int64 `json:"bucket_start_time"`
	// The aggregated meter usage value for the specified bucket.
	BucketValue float64 `json:"bucket_value"`
	// A set of key-value pairs representing the dimensions of the meter usage.
	Dimensions map[string]string `json:"dimensions"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The unique identifier for the meter.
	MeterID string `json:"meter_id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
