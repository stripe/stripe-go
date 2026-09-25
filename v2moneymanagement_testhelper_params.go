//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates an EarnedCredit in a Sandbox environment for testing purposes.
type V2MoneyManagementTestHelperEarnedCreditsParams struct {
	Params `form:"*"`
	// The amount and currency of the EarnedCredit.
	Amount *Amount `form:"amount" json:"amount"`
	// The FinancialAccount to simulate the EarnedCredit for.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// The type of EarnedCredit to create. Currently only interest is supported.
	Type *string `form:"type" json:"type"`
}
