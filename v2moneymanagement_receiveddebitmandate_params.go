//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Deprecated: Use ListSpendingMandates instead.
// Returns a list of ReceivedDebitMandates.
type V2MoneyManagementReceivedDebitMandateListParams struct {
	Params `form:"*"`
	// The ID of the FinancialAccount to filter by.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The page limit.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by mandate status.
	Statuses []*string `form:"statuses" json:"statuses,omitempty"`
	// The type of ReceivedDebitMandate to filter by.
	Type *string `form:"type" json:"type,omitempty"`
}

// Deprecated: Use RetrieveSpendingMandate instead.
// Retrieves the details of an existing ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateParams struct {
	Params `form:"*"`
}

// Deprecated: Use CancelSpendingMandate instead.
// Cancels an active ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateCancelParams struct {
	Params `form:"*"`
}

// Deprecated: Use RetrieveSpendingMandate instead.
// Retrieves the details of an existing ReceivedDebitMandate.
type V2MoneyManagementReceivedDebitMandateRetrieveParams struct {
	Params `form:"*"`
}
