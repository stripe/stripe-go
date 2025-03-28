//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
type V2TestHelpersFinancialAddressCreditParams struct {
	Params `form:"*"`
	// Object containing the amount value and currency to credit.
	Amount *Amount `form:"amount" json:"amount"`
	// Open Enum. The network to use in simulating the funds flow. This will be the reflected in the resulting ReceivedCredit.
	Network *string `form:"network" json:"network"`
	// String explaining funds flow. Use this field to populate the statement descriptor of the ReceivedCredit created as an eventual result of this simulation.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
type V2TestHelpersFinancialAddressGenerateMicrodepositsParams struct {
	Params `form:"*"`
}
