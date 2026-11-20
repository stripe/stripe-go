//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a list of DisputeAdjustments.
type V2MoneyManagementDisputeAdjustmentListParams struct {
	Params `form:"*"`
	// Filter results by the ID of the FinancialAccount associated with the DisputeAdjustment.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// Maximum number of results to return, between 1 and 100. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter results by the ID of the Dispute associated with the DisputeAdjustment.
	RelatedDispute *string `form:"related_dispute" json:"related_dispute,omitempty"`
}

// Retrieves a DisputeAdjustment by ID.
type V2MoneyManagementDisputeAdjustmentParams struct {
	Params `form:"*"`
}

// Retrieves a DisputeAdjustment by ID.
type V2MoneyManagementDisputeAdjustmentRetrieveParams struct {
	Params `form:"*"`
}
