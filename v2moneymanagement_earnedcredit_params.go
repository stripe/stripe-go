//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of EarnedCredits.
type V2MoneyManagementEarnedCreditListParams struct {
	Params `form:"*"`
	// The FinancialAccount to list EarnedCredits for.
	FinancialAccount *string `form:"financial_account" json:"financial_account,omitempty"`
	// The maximum number of EarnedCredits to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves an EarnedCredit.
type V2MoneyManagementEarnedCreditParams struct {
	Params `form:"*"`
}

// Retrieves an EarnedCredit.
type V2MoneyManagementEarnedCreditRetrieveParams struct {
	Params `form:"*"`
}
