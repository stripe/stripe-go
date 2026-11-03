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

// Create a GB bank account.
type V2CoreVaultGBBankAccountParams struct {
	Params `form:"*"`
	// The Account Number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
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

// Unarchive a GBBankAccount object.
type V2CoreVaultGBBankAccountUnarchiveParams struct {
	Params `form:"*"`
}

// Create a GB bank account.
type V2CoreVaultGBBankAccountCreateParams struct {
	Params `form:"*"`
	// The Account Number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType *string `form:"bank_account_type" json:"bank_account_type,omitempty"`
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
