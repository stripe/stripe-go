//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. The type of the bank account (checking or savings).
type V2CoreVaultGBBankAccountBankAccountType string

// List of values that V2CoreVaultGBBankAccountBankAccountType can take
const (
	V2CoreVaultGBBankAccountBankAccountTypeChecking V2CoreVaultGBBankAccountBankAccountType = "checking"
	V2CoreVaultGBBankAccountBankAccountTypeSavings  V2CoreVaultGBBankAccountBankAccountType = "savings"
)

// Use the GBBankAccounts API to create and manage GB bank account objects
type V2CoreVaultGBBankAccount struct {
	APIResource
	// Whether this bank account object was archived. Bank account objects can be archived through
	// the /archive API, and they will not be automatically archived by Stripe. Archived bank account objects
	// cannot be used as outbound destinations and will not appear in the outbound destination list.
	Archived bool `json:"archived"`
	// Closed Enum. The type of the bank account (checking or savings).
	BankAccountType V2CoreVaultGBBankAccountBankAccountType `json:"bank_account_type"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// Creation time.
	Created time.Time `json:"created"`
	// The ID of the GBBankAccount object.
	ID string `json:"id"`
	// The last 4 digits of the account number or IBAN.
	Last4 string `json:"last4"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Whether the GB Bank Account is currently unusable for money movement, despite potentially being correctly set up. Please reach out to Stripe Support for more information.
	Restricted bool `json:"restricted"`
	// The Sort Code of the bank account.
	SortCode string `json:"sort_code,omitempty"`
	// The list of currencies supported by this bank account.
	SupportedCurrencies []Currency `json:"supported_currencies"`
}
