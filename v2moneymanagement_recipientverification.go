//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of resource that consumed the RecipientVerification.
type V2MoneyManagementRecipientVerificationConsumedByType string

// List of values that V2MoneyManagementRecipientVerificationConsumedByType can take
const (
	V2MoneyManagementRecipientVerificationConsumedByTypeOutboundPayment V2MoneyManagementRecipientVerificationConsumedByType = "outbound_payment"
)

// Closed Enum. Match level of the RecipientVerification: `match`, `close_match`, `no_match`, `unavailable`.
type V2MoneyManagementRecipientVerificationMatchResult string

// List of values that V2MoneyManagementRecipientVerificationMatchResult can take
const (
	V2MoneyManagementRecipientVerificationMatchResultCloseMatch  V2MoneyManagementRecipientVerificationMatchResult = "close_match"
	V2MoneyManagementRecipientVerificationMatchResultMatch       V2MoneyManagementRecipientVerificationMatchResult = "match"
	V2MoneyManagementRecipientVerificationMatchResultNoMatch     V2MoneyManagementRecipientVerificationMatchResult = "no_match"
	V2MoneyManagementRecipientVerificationMatchResultUnavailable V2MoneyManagementRecipientVerificationMatchResult = "unavailable"
)

// The business type matched by the provider. Only present for Confirmation of Payee (CoP) verifications.
type V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessType string

// List of values that V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessType can take
const (
	V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessTypeBusiness V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessType = "business"
	V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessTypePersonal V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessType = "personal"
)

// The business type provided by the user. Only present for Confirmation of Payee (CoP) verifications.
type V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessType string

// List of values that V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessType can take
const (
	V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessTypeBusiness V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessType = "business"
	V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessTypePersonal V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessType = "personal"
)

// Current status of the recipient verification.
type V2MoneyManagementRecipientVerificationStatus string

// List of values that V2MoneyManagementRecipientVerificationStatus can take
const (
	V2MoneyManagementRecipientVerificationStatusAcknowledged            V2MoneyManagementRecipientVerificationStatus = "acknowledged"
	V2MoneyManagementRecipientVerificationStatusAwaitingAcknowledgement V2MoneyManagementRecipientVerificationStatus = "awaiting_acknowledgement"
	V2MoneyManagementRecipientVerificationStatusConsumed                V2MoneyManagementRecipientVerificationStatus = "consumed"
	V2MoneyManagementRecipientVerificationStatusExpired                 V2MoneyManagementRecipientVerificationStatus = "expired"
	V2MoneyManagementRecipientVerificationStatusVerified                V2MoneyManagementRecipientVerificationStatus = "verified"
)

// Open Enum. The type of verification performed: `verification_of_payee`, `confirmation_of_payee`.
type V2MoneyManagementRecipientVerificationVerificationType string

// List of values that V2MoneyManagementRecipientVerificationVerificationType can take
const (
	V2MoneyManagementRecipientVerificationVerificationTypeConfirmationOfPayee V2MoneyManagementRecipientVerificationVerificationType = "confirmation_of_payee"
	V2MoneyManagementRecipientVerificationVerificationTypeVerificationOfPayee V2MoneyManagementRecipientVerificationVerificationType = "verification_of_payee"
)

// The resource that consumed this verification.
type V2MoneyManagementRecipientVerificationConsumedBy struct {
	// The ID of the OutboundPayment that consumed the RecipientVerification.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// The type of resource that consumed the RecipientVerification.
	Type V2MoneyManagementRecipientVerificationConsumedByType `json:"type"`
}

// Details for the match result.
type V2MoneyManagementRecipientVerificationMatchResultDetails struct {
	// The business type matched by the provider. Only present for Confirmation of Payee (CoP) verifications.
	MatchedBusinessType V2MoneyManagementRecipientVerificationMatchResultDetailsMatchedBusinessType `json:"matched_business_type,omitempty"`
	// The account name associated with the bank account as provided by the VoP provider, only present if there is a match or close match.
	MatchedName string `json:"matched_name,omitempty"`
	// A message describing the match result.
	Message string `json:"message"`
	// The business type provided by the user. Only present for Confirmation of Payee (CoP) verifications.
	ProvidedBusinessType V2MoneyManagementRecipientVerificationMatchResultDetailsProvidedBusinessType `json:"provided_business_type,omitempty"`
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
	// The resource that consumed this verification.
	ConsumedBy *V2MoneyManagementRecipientVerificationConsumedBy `json:"consumed_by,omitempty"`
	// Time at which the RecipientVerification was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Time at which the RecipientVerification expires, 5 minutes after the creation. Only applicable for Verification of Payee (VoP) verifications.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
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
	// ID of the payout method that was verified.
	PayoutMethod string `json:"payout_method"`
	// ID of the recipient account. Omitted for same-entity payouts.
	Recipient string `json:"recipient,omitempty"`
	// Current status of the recipient verification.
	Status V2MoneyManagementRecipientVerificationStatus `json:"status"`
	// Hash containing timestamps of when the object transitioned to a particular status.
	StatusTransitions *V2MoneyManagementRecipientVerificationStatusTransitions `json:"status_transitions,omitempty"`
	// Open Enum. The type of verification performed: `verification_of_payee`, `confirmation_of_payee`.
	VerificationType V2MoneyManagementRecipientVerificationVerificationType `json:"verification_type"`
}
