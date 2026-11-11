//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of SpendingMandates.
type V2MoneyManagementSpendingMandateListParams struct {
	Params `form:"*"`
	// The ID of the FinancialAccount to filter by.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by mandate status.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
	// The type of SpendingMandate to filter by.
	Type *string `form:"type" json:"type,omitempty"`
}

// Retrieves the details of an existing SpendingMandate.
type V2MoneyManagementSpendingMandateParams struct {
	Params `form:"*"`
}

// Cancels an active SpendingMandate.
type V2MoneyManagementSpendingMandateCancelParams struct {
	Params `form:"*"`
}

// Retrieves the details of an existing SpendingMandate.
type V2MoneyManagementSpendingMandateRetrieveParams struct {
	Params `form:"*"`
}
