//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Adjustments that match the provided filters.
type V2MoneyManagementAdjustmentListParams struct {
	Params `form:"*"`
	// Filter for Adjustments linked to a Flow.
	AdjustedFlow *string `form:"adjusted_flow" json:"adjusted_flow,omitempty"`
	// Set of filters to query Adjustments within a range of `created` timestamps.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves the details of an Adjustment by ID.
type V2MoneyManagementAdjustmentParams struct {
	Params `form:"*"`
}

// Retrieves the details of an Adjustment by ID.
type V2MoneyManagementAdjustmentRetrieveParams struct {
	Params `form:"*"`
}
