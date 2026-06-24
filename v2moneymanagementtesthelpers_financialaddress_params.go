//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulate debiting a FinancialAddress in a Sandbox environment. This can be used to remove virtual funds and decrease your balance for testing.
type V2MoneyManagementTestHelpersFinancialAddressDebitParams struct {
	Params `form:"*"`
	// Object containing the amount value and currency to debit.
	Amount *Amount `form:"amount" json:"amount"`
	// The network to use in simulating the funds flow. This will be reflected in the resulting ReceivedDebit.
	Network *string `form:"network" json:"network"`
	// String explaining funds flow. Use this field to populate the statement descriptor of the ReceivedDebit created as an eventual result of this simulation.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}
