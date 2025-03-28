//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Archive a UsBankAccount object. UsBankAccount objects will not be automatically archived by Stripe.
// Archived UsBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
type V2CoreVaultUSBankAccountArchiveParams struct {
	Params `form:"*"`
}

// Create a UsBankAccount object.
type V2CoreVaultUSBankAccountParams struct {
	Params `form:"*"`
	// The account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// The fedwire routing number of the bank account. Note that certain banks have the same ACH and wire routing number.
	FedwireRoutingNumber *string `form:"fedwire_routing_number" json:"fedwire_routing_number,omitempty"`
	// The ACH routing number of the bank account. Note that certain banks have the same ACH and wire routing number.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}
