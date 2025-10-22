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

// Microdeposit type can be amounts or descriptor_type.
type V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositType string

// List of values that V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositType can take
const (
	V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositTypeAmounts        V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositType = "amounts"
	V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositTypeDescriptorCode V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositType = "descriptor_code"
)

// The bank account verification status.
type V2CoreVaultUSBankAccountVerificationStatus string

// List of values that V2CoreVaultUSBankAccountVerificationStatus can take
const (
	V2CoreVaultUSBankAccountVerificationStatusAwaitingVerification V2CoreVaultUSBankAccountVerificationStatus = "awaiting_verification"
	V2CoreVaultUSBankAccountVerificationStatusUnverified           V2CoreVaultUSBankAccountVerificationStatus = "unverified"
	V2CoreVaultUSBankAccountVerificationStatusVerificationFailed   V2CoreVaultUSBankAccountVerificationStatus = "verification_failed"
	V2CoreVaultUSBankAccountVerificationStatusVerified             V2CoreVaultUSBankAccountVerificationStatus = "verified"
)

// The microdeposit verification details if the status is awaiting verification.
type V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetails struct {
	// Time when microdeposits will expire and have to be re-sent.
	Expires time.Time `json:"expires"`
	// Microdeposit type can be amounts or descriptor_type.
	MicrodepositType V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetailsMicrodepositType `json:"microdeposit_type"`
	// Time when microdeposits were sent.
	Sent time.Time `json:"sent"`
}

// The bank account verification details.
type V2CoreVaultUSBankAccountVerification struct {
	// The microdeposit verification details if the status is awaiting verification.
	MicrodepositVerificationDetails *V2CoreVaultUSBankAccountVerificationMicrodepositVerificationDetails `json:"microdeposit_verification_details,omitempty"`
	// The bank account verification status.
	Status V2CoreVaultUSBankAccountVerificationStatus `json:"status"`
}

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
	// The bank account verification details.
	Verification *V2CoreVaultUSBankAccountVerification `json:"verification"`
}
