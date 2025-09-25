//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. Match level of the RecipientVerification: `match`, `close_match`, `no_match`, `unavailable`.
type V2MoneyManagementRecipientVerificationMatchResult string

// List of values that V2MoneyManagementRecipientVerificationMatchResult can take
const (
	V2MoneyManagementRecipientVerificationMatchResultCloseMatch  V2MoneyManagementRecipientVerificationMatchResult = "close_match"
	V2MoneyManagementRecipientVerificationMatchResultMatch       V2MoneyManagementRecipientVerificationMatchResult = "match"
	V2MoneyManagementRecipientVerificationMatchResultNoMatch     V2MoneyManagementRecipientVerificationMatchResult = "no_match"
	V2MoneyManagementRecipientVerificationMatchResultUnavailable V2MoneyManagementRecipientVerificationMatchResult = "unavailable"
)

// Closed Enum. Current status of the RecipientVerification: `verified`, `consumed`, `expired`, `awaiting_acknowledgement`, `acknowledged`.
type V2MoneyManagementRecipientVerificationStatus string

// List of values that V2MoneyManagementRecipientVerificationStatus can take
const (
	V2MoneyManagementRecipientVerificationStatusAcknowledged            V2MoneyManagementRecipientVerificationStatus = "acknowledged"
	V2MoneyManagementRecipientVerificationStatusAwaitingAcknowledgement V2MoneyManagementRecipientVerificationStatus = "awaiting_acknowledgement"
	V2MoneyManagementRecipientVerificationStatusConsumed                V2MoneyManagementRecipientVerificationStatus = "consumed"
	V2MoneyManagementRecipientVerificationStatusExpired                 V2MoneyManagementRecipientVerificationStatus = "expired"
	V2MoneyManagementRecipientVerificationStatusVerified                V2MoneyManagementRecipientVerificationStatus = "verified"
)

// Details for the match result.
type V2MoneyManagementRecipientVerificationMatchResultDetails struct {
	// The account name associated with the bank account as provided by the VoP provider, only present if there is a match or close match.
	MatchedName string `json:"matched_name,omitempty"`
	// A message describing the match result.
	Message string `json:"message"`
	// The name associated with the provided recipient.
	ProvidedName string `json:"provided_name"`
}

// Hash containing timestamps of when the object transitioned to a particular status.
type V2MoneyManagementRecipientVerificationStatusTransitions struct {
	// Timestamp describing when a RecipientVerification changed status to `acknowledged`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	AcknowledgedAt time.Time `json:"acknowledged_at,omitempty"`
	// Timestamp describing when a RecipientVerification changed status to `consumed`.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ConsumedAt time.Time `json:"consumed_at,omitempty"`
}

// RecipientVerification represents a verification of recipient you intend to send funds to.
type V2MoneyManagementRecipientVerification struct {
	APIResource
	// The OBP/OBT ID that consumed this verification, present if one is successfully created.
	ConsumedBy string `json:"consumed_by,omitempty"`
	// Time at which the RecipientVerification was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Time at which the RecipientVerification expires, 5 minutes after the creation.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpiresAt time.Time `json:"expires_at"`
	// The ID of the RecipientVerification.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Closed Enum. Match level of the RecipientVerification: `match`, `close_match`, `no_match`, `unavailable`.
	MatchResult V2MoneyManagementRecipientVerificationMatchResult `json:"match_result"`
	// Details for the match result.
	MatchResultDetails *V2MoneyManagementRecipientVerificationMatchResultDetails `json:"match_result_details"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. Current status of the RecipientVerification: `verified`, `consumed`, `expired`, `awaiting_acknowledgement`, `acknowledged`.
	Status V2MoneyManagementRecipientVerificationStatus `json:"status"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementRecipientVerificationStatusTransitions `json:"status_transitions,omitempty"`
}
