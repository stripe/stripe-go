//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. The type of bank account (checking or savings).
type V2CoreVaultUSBankAccountBankAccountType string

// List of values that V2CoreVaultUSBankAccountBankAccountType can take
const (
	V2CoreVaultUSBankAccountBankAccountTypeChecking V2CoreVaultUSBankAccountBankAccountType = "checking"
	V2CoreVaultUSBankAccountBankAccountTypeSavings  V2CoreVaultUSBankAccountBankAccountType = "savings"
)

// Use the USBankAccounts API to create and manage US bank accounts objects that you can use to receive funds. Note that these are not interchangeable with v1 Tokens.
type V2CoreVaultUSBankAccount struct {
	APIResource
	// Whether this USBankAccount object was archived.
	Archived bool `json:"archived"`
	// Closed Enum. The type of bank account (checking or savings).
	BankAccountType V2CoreVaultUSBankAccountBankAccountType `json:"bank_account_type"`
	// The name of the bank this bank account belongs to. This field is populated automatically by Stripe based on the routing number.
	BankName string `json:"bank_name"`
	// Creation time of the object.
	Created time.Time `json:"created"`
	// The fedwire routing number of the bank account.
	FedwireRoutingNumber string `json:"fedwire_routing_number,omitempty"`
	// The ID of the USBankAccount object.
	ID string `json:"id"`
	// The last 4 digits of the account number.
	Last4 string `json:"last4"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ACH routing number of the bank account.
	RoutingNumber string `json:"routing_number,omitempty"`
}
