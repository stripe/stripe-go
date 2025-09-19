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

// Whether or not the information of the bank account matches what you have provided. Closed enum.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult string

// List of values that V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult can take
const (
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResultMatch        V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult = "match"
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResultMismatch     V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult = "mismatch"
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResultPartialMatch V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult = "partial_match"
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResultUnavailable  V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult = "unavailable"
)

// The business type given by the bank for this account, in case of a MATCH or PARTIAL_MATCH.
// Closed enum.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessType string

// List of values that V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessType can take
const (
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessTypeBusiness V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessType = "business"
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessTypePersonal V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessType = "personal"
)

// The provided or Legal Entity business type to match against the CoP service. Closed enum.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessType string

// List of values that V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessType can take
const (
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessTypeBusiness V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessType = "business"
	V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessTypePersonal V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessType = "personal"
)

// The current state of Confirmation of Payee on this bank account. Closed enum.
type V2CoreVaultGBBankAccountConfirmationOfPayeeStatus string

// List of values that V2CoreVaultGBBankAccountConfirmationOfPayeeStatus can take
const (
	V2CoreVaultGBBankAccountConfirmationOfPayeeStatusAwaitingAcknowledgement V2CoreVaultGBBankAccountConfirmationOfPayeeStatus = "awaiting_acknowledgement"
	V2CoreVaultGBBankAccountConfirmationOfPayeeStatusConfirmed               V2CoreVaultGBBankAccountConfirmationOfPayeeStatus = "confirmed"
	V2CoreVaultGBBankAccountConfirmationOfPayeeStatusUninitiated             V2CoreVaultGBBankAccountConfirmationOfPayeeStatus = "uninitiated"
)

// The fields that CoP service matched against. Only has value if MATCH or PARTIAL_MATCH, empty otherwise.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatched struct {
	// The business type given by the bank for this account, in case of a MATCH or PARTIAL_MATCH.
	// Closed enum.
	BusinessType V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchedBusinessType `json:"business_type,omitempty"`
	// The name given by the bank for this account, in case of a MATCH or PARTIAL_MATCH.
	Name string `json:"name,omitempty"`
}

// The fields that are matched against what the network has on file.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvided struct {
	// The provided or Legal Entity business type to match against the CoP service. Closed enum.
	BusinessType V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvidedBusinessType `json:"business_type"`
	// The provided or Legal Entity name to match against the CoP service.
	Name string `json:"name"`
}

// The result of the Confirmation of Payee check, once the check has been initiated. Closed enum.
type V2CoreVaultGBBankAccountConfirmationOfPayeeResult struct {
	// When the CoP result was created.
	Created time.Time `json:"created"`
	// The fields that CoP service matched against. Only has value if MATCH or PARTIAL_MATCH, empty otherwise.
	Matched *V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatched `json:"matched"`
	// Whether or not the information of the bank account matches what you have provided. Closed enum.
	MatchResult V2CoreVaultGBBankAccountConfirmationOfPayeeResultMatchResult `json:"match_result"`
	// Human-readable message describing the match result.
	Message string `json:"message"`
	// The fields that are matched against what the network has on file.
	Provided *V2CoreVaultGBBankAccountConfirmationOfPayeeResultProvided `json:"provided"`
}

// Information around the status of Confirmation of Payee matching done on this bank account.
// Confirmation of Payee is a name matching service that must be done before making OutboundPayments in the UK.
type V2CoreVaultGBBankAccountConfirmationOfPayee struct {
	// The result of the Confirmation of Payee check, once the check has been initiated. Closed enum.
	Result *V2CoreVaultGBBankAccountConfirmationOfPayeeResult `json:"result"`
	// The current state of Confirmation of Payee on this bank account. Closed enum.
	Status V2CoreVaultGBBankAccountConfirmationOfPayeeStatus `json:"status"`
}

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
	// Information around the status of Confirmation of Payee matching done on this bank account.
	// Confirmation of Payee is a name matching service that must be done before making OutboundPayments in the UK.
	ConfirmationOfPayee *V2CoreVaultGBBankAccountConfirmationOfPayee `json:"confirmation_of_payee"`
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
	// The Sort Code of the bank account.
	SortCode string `json:"sort_code"`
}
