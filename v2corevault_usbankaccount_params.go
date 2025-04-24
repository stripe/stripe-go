//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Archive a USBankAccount object. USBankAccount objects will not be automatically archived by Stripe.
// Archived USBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
type V2CoreVaultUSBankAccountArchiveParams struct {
	Params `form:"*"`
}

// Create a USBankAccount object.
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

// Create a USBankAccount object.
type V2CoreVaultUSBankAccountCreateParams struct {
	Params `form:"*"`
	// The account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// The fedwire routing number of the bank account. Note that certain banks have the same ACH and wire routing number.
	FedwireRoutingNumber *string `form:"fedwire_routing_number" json:"fedwire_routing_number,omitempty"`
	// The ACH routing number of the bank account. Note that certain banks have the same ACH and wire routing number.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// Retrieve a USBankAccount object.
type V2CoreVaultUSBankAccountRetrieveParams struct {
	Params `form:"*"`
}

// Update a USBankAccount object. This is limited to supplying a previously empty routing_number field.
type V2CoreVaultUSBankAccountUpdateParams struct {
	Params `form:"*"`
	// The bank account's fedwire routing number can be provided for update it was were empty previously.
	FedwireRoutingNumber *string `form:"fedwire_routing_number" json:"fedwire_routing_number,omitempty"`
	// The bank account's ACH routing number can be provided for update if it was empty previously.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}
