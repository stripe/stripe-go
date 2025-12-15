//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Object containing the amount value and currency to credit.
type V2TestHelpersFinancialAddressCreditAmountParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value *int64 `form:"value" json:"value,omitempty"`
}

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
type V2TestHelpersFinancialAddressCreditParams struct {
	Params `form:"*"`
	// Object containing the amount value and currency to credit.
	Amount *V2TestHelpersFinancialAddressCreditAmountParams `form:"amount" json:"amount"`
	// Open Enum. The network to use in simulating the funds flow. This will be the reflected in the resulting ReceivedCredit.
	Network *string `form:"network" json:"network"`
	// String explaining funds flow. Use this field to populate the statement descriptor of the ReceivedCredit created as an eventual result of this simulation.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
type V2TestHelpersFinancialAddressGenerateMicrodepositsParams struct {
	Params `form:"*"`
}
