//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Returns a list of Adjustments that match the provided filters.
type V2MoneyManagementAdjustmentListParams struct {
	Params `form:"*"`
	// Filter for Adjustments linked to a Flow.
	AdjustedFlow *string `form:"adjusted_flow" json:"adjusted_flow,omitempty"`
	// Filter for objects created at the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for objects created after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for objects created on or after the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for objects created before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for objects created on or before the specified timestamp.
	// Must be an RFC 3339 date & time value, for example: 2022-09-18T13:22:00Z.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves the details of an Adjustment by ID.
type V2MoneyManagementAdjustmentParams struct {
	Params `form:"*"`
}
