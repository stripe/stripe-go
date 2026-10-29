//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Adjustments that match the provided filters.
type V2MoneyManagementAdjustmentListParams struct {
	Params `form:"*"`
	// Filter for Adjustments linked to a specific flow object. This filter only supports filtering by
	// the flow object's ID, such as `ibt_123` for an InboundTransfer.
	AdjustedFlow *string `form:"adjusted_flow" json:"adjusted_flow,omitempty"`
	// Filter for Adjustments by their creation time using `gt`, `gte`, `lt`, or `lte`.
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
