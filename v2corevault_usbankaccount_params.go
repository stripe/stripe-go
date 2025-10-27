//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List USBankAccount objects. Optionally filter by verification status.
type V2CoreVaultUSBankAccountListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optionally filter by verification status. Mutually exclusive with `unverified`, `verified`, `awaiting_verification`, and `verification_failed`.
	VerificationStatus *string `form:"verification_status" json:"verification_status,omitempty"`
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

// Archive a USBankAccount object. USBankAccount objects will not be automatically archived by Stripe.
// Archived USBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
type V2CoreVaultUSBankAccountArchiveParams struct {
	Params `form:"*"`
}

// Confirm microdeposits amounts or descriptor code that you have received from the Send Microdeposits request. Once you correctly confirm this, this US Bank Account will be verified and eligible to transfer funds with.
type V2CoreVaultUSBankAccountConfirmMicrodepositsParams struct {
	Params `form:"*"`
	// Two amounts received through Send Microdeposits must match the input to Confirm Microdeposits to verify US Bank Account.
	Amounts []*int64 `form:"amounts,flat_array" json:"amounts,omitempty"`
	// Descriptor code received through Send Microdeposits must match the input to Confirm Microdeposits to verify US Bank Account.
	DescriptorCode *string `form:"descriptor_code" json:"descriptor_code,omitempty"`
}

// Send microdeposits in order to verify your US Bank Account so it is eligible to transfer funds. This will start the verification process and you must Confirm Microdeposits to successfully verify your US Bank Account.
type V2CoreVaultUSBankAccountSendMicrodepositsParams struct {
	Params `form:"*"`
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
