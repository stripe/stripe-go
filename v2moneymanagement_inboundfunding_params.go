//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists InboundFundings.
type V2MoneyManagementInboundFundingListParams struct {
	Params `form:"*"`
	// The maximum number of results to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves an InboundFunding by ID.
type V2MoneyManagementInboundFundingParams struct {
	Params `form:"*"`
}

// Retrieves an InboundFunding by ID.
type V2MoneyManagementInboundFundingRetrieveParams struct {
	Params `form:"*"`
}
