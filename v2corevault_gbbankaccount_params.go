//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List objects that can be used as destinations for outbound money movement via OutboundPayment.
type V2CoreVaultGBBankAccountListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Whether or not to automatically perform Confirmation of Payee to verify the users information
// against what was provided by the bank. Doing so is required for all bank accounts not owned
// by you before making domestic UK OutboundPayments.
type V2CoreVaultGBBankAccountConfirmationOfPayeeParams struct {
	// The business type to be checked against. Legal entity information will be used if unspecified.
	// Closed enum.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// User specifies whether Confirmation of Payee is automatically initiated when creating the bank account.
	Initiate *bool `form:"initiate" json:"initiate"`
	// The name to be checked against. Legal entity information will be used if unspecified.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create a GB bank account.
type V2CoreVaultGBBankAccountParams struct {
	Params `form:"*"`
	// The Account Number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// Whether or not to automatically perform Confirmation of Payee to verify the users information
	// against what was provided by the bank. Doing so is required for all bank accounts not owned
	// by you before making domestic UK OutboundPayments.
	ConfirmationOfPayee *V2CoreVaultGBBankAccountConfirmationOfPayeeParams `form:"confirmation_of_payee" json:"confirmation_of_payee,omitempty"`
	// The currency of the bank account.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The IBAN of the bank account.
	IBAN *string `form:"iban" json:"iban,omitempty"`
	// The Sort Code of the bank account.
	SortCode *string `form:"sort_code" json:"sort_code,omitempty"`
}

// Archive a GBBankAccount object. Archived GBBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
type V2CoreVaultGBBankAccountArchiveParams struct {
	Params `form:"*"`
}

// Whether or not to automatically perform Confirmation of Payee to verify the users information
// against what was provided by the bank. Doing so is required for all bank accounts not owned
// by you before making domestic UK OutboundPayments.
type V2CoreVaultGBBankAccountCreateConfirmationOfPayeeParams struct {
	// The business type to be checked against. Legal entity information will be used if unspecified.
	// Closed enum.
	BusinessType *string `form:"business_type" json:"business_type,omitempty"`
	// User specifies whether Confirmation of Payee is automatically initiated when creating the bank account.
	Initiate *bool `form:"initiate" json:"initiate"`
	// The name to be checked against. Legal entity information will be used if unspecified.
	Name *string `form:"name" json:"name,omitempty"`
}

// Create a GB bank account.
type V2CoreVaultGBBankAccountCreateParams struct {
	Params `form:"*"`
	// The Account Number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
	// Whether or not to automatically perform Confirmation of Payee to verify the users information
	// against what was provided by the bank. Doing so is required for all bank accounts not owned
	// by you before making domestic UK OutboundPayments.
	ConfirmationOfPayee *V2CoreVaultGBBankAccountCreateConfirmationOfPayeeParams `form:"confirmation_of_payee" json:"confirmation_of_payee,omitempty"`
	// The currency of the bank account.
	Currency *string `form:"currency" json:"currency"`
	// The IBAN of the bank account.
	IBAN *string `form:"iban" json:"iban,omitempty"`
	// The Sort Code of the bank account.
	SortCode *string `form:"sort_code" json:"sort_code,omitempty"`
}

// Retrieve a GB bank account.
type V2CoreVaultGBBankAccountRetrieveParams struct {
	Params `form:"*"`
}
