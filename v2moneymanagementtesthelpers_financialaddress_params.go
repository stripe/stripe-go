//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulate crediting a FinancialAddress in a Sandbox environment. This can be used to add virtual funds and increase your balance for testing.
type V2MoneyManagementTestHelpersFinancialAddressCreditParams struct {
	Params `form:"*"`
	// Object containing the amount value and currency to credit.
	Amount *Amount `form:"amount" json:"amount"`
	// Open Enum. The network to use in simulating the funds flow. This will be the reflected in the resulting ReceivedCredit.
	Network *string `form:"network" json:"network"`
	// String explaining funds flow. Use this field to populate the statement descriptor of the ReceivedCredit created as an eventual result of this simulation.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// ACH-specific details for the simulated debit.
type V2MoneyManagementTestHelpersFinancialAddressDebitNetworkDetailsACHParams struct {
	// The ACH company ID of the originator of the simulated debit. Must contain exactly 10 alphanumeric characters. Sent to the network in uppercase.
	OriginatorCompanyID *string `form:"originator_company_id" json:"originator_company_id"`
}

// Network-specific details for the simulated debit. Only used for ACH debits; supplied details must be valid.
type V2MoneyManagementTestHelpersFinancialAddressDebitNetworkDetailsParams struct {
	// ACH-specific details for the simulated debit.
	ACH *V2MoneyManagementTestHelpersFinancialAddressDebitNetworkDetailsACHParams `form:"ach" json:"ach,omitempty"`
}

// Simulate debiting a FinancialAddress in a Sandbox environment. This can be used to remove virtual funds and decrease your balance for testing.
type V2MoneyManagementTestHelpersFinancialAddressDebitParams struct {
	Params `form:"*"`
	// Object containing the amount value and currency to debit.
	Amount *Amount `form:"amount" json:"amount"`
	// The network to use in simulating the funds flow. This will be reflected in the resulting ReceivedDebit.
	Network *string `form:"network" json:"network"`
	// Network-specific details for the simulated debit. Only used for ACH debits; supplied details must be valid.
	NetworkDetails *V2MoneyManagementTestHelpersFinancialAddressDebitNetworkDetailsParams `form:"network_details" json:"network_details,omitempty"`
	// String explaining funds flow. Use this field to populate the statement descriptor of the ReceivedDebit created as an eventual result of this simulation.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// Generates microdeposits for a FinancialAddress in a Sandbox environment.
type V2MoneyManagementTestHelpersFinancialAddressGenerateMicrodepositsParams struct {
	Params `form:"*"`
}
