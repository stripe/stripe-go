//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The reason the account or customer was restricted.
type V2SignalsAccountActivityAccountRestrictedReason string

// List of values that V2SignalsAccountActivityAccountRestrictedReason can take
const (
	V2SignalsAccountActivityAccountRestrictedReasonAbuse V2SignalsAccountActivityAccountRestrictedReason = "abuse"
	V2SignalsAccountActivityAccountRestrictedReasonOther V2SignalsAccountActivityAccountRestrictedReason = "other"
)

// The outcome of the merchant review.
type V2SignalsAccountActivityAccountReviewedOutcome string

// List of values that V2SignalsAccountActivityAccountReviewedOutcome can take
const (
	V2SignalsAccountActivityAccountReviewedOutcomeTrusted V2SignalsAccountActivityAccountReviewedOutcome = "trusted"
)

// The reason the customer was suspended.
type V2SignalsAccountActivityAccountSuspendedReason string

// List of values that V2SignalsAccountActivityAccountSuspendedReason can take
const (
	V2SignalsAccountActivityAccountSuspendedReasonAbuse V2SignalsAccountActivityAccountSuspendedReason = "abuse"
	V2SignalsAccountActivityAccountSuspendedReasonOther V2SignalsAccountActivityAccountSuspendedReason = "other"
)

// The type of credential involved in the confirmed takeover.
type V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType string

// List of values that V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType can take
const (
	V2SignalsAccountActivityAccountTakeoverConfirmationCredentialTypeAPIKey        V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType = "api_key"
	V2SignalsAccountActivityAccountTakeoverConfirmationCredentialTypeLoginPassword V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType = "login_password"
	V2SignalsAccountActivityAccountTakeoverConfirmationCredentialTypeOther         V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType = "other"
)

// The type of credential involved in the reported takeover.
type V2SignalsAccountActivityAccountTakeoverReportCredentialType string

// List of values that V2SignalsAccountActivityAccountTakeoverReportCredentialType can take
const (
	V2SignalsAccountActivityAccountTakeoverReportCredentialTypeAPIKey        V2SignalsAccountActivityAccountTakeoverReportCredentialType = "api_key"
	V2SignalsAccountActivityAccountTakeoverReportCredentialTypeLoginPassword V2SignalsAccountActivityAccountTakeoverReportCredentialType = "login_password"
	V2SignalsAccountActivityAccountTakeoverReportCredentialTypeOther         V2SignalsAccountActivityAccountTakeoverReportCredentialType = "other"
)

// The type of credential that changed.
type V2SignalsAccountActivityCredentialChangeCredentialType string

// List of values that V2SignalsAccountActivityCredentialChangeCredentialType can take
const (
	V2SignalsAccountActivityCredentialChangeCredentialTypeAPIKey        V2SignalsAccountActivityCredentialChangeCredentialType = "api_key"
	V2SignalsAccountActivityCredentialChangeCredentialTypeLoginPassword V2SignalsAccountActivityCredentialChangeCredentialType = "login_password"
	V2SignalsAccountActivityCredentialChangeCredentialTypeOther         V2SignalsAccountActivityCredentialChangeCredentialType = "other"
)

// The outcome of the action.
type V2SignalsAccountActivityCustomMarketingActionOutcome string

// List of values that V2SignalsAccountActivityCustomMarketingActionOutcome can take
const (
	V2SignalsAccountActivityCustomMarketingActionOutcomeFailed    V2SignalsAccountActivityCustomMarketingActionOutcome = "failed"
	V2SignalsAccountActivityCustomMarketingActionOutcomeSucceeded V2SignalsAccountActivityCustomMarketingActionOutcome = "succeeded"
)

// The authentication method used for the challenge.
type V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType string

// List of values that V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType can take
const (
	V2SignalsAccountActivityLoginDecisionMfaChallengeMfaTypeAuthenticatorApp V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType = "authenticator_app"
	V2SignalsAccountActivityLoginDecisionMfaChallengeMfaTypeEmail            V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType = "email"
	V2SignalsAccountActivityLoginDecisionMfaChallengeMfaTypeHardwareKey      V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType = "hardware_key"
	V2SignalsAccountActivityLoginDecisionMfaChallengeMfaTypeOther            V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType = "other"
	V2SignalsAccountActivityLoginDecisionMfaChallengeMfaTypeSms              V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType = "sms"
)

// The outcome of the authentication challenge.
type V2SignalsAccountActivityLoginDecisionMfaChallengeStatus string

// List of values that V2SignalsAccountActivityLoginDecisionMfaChallengeStatus can take
const (
	V2SignalsAccountActivityLoginDecisionMfaChallengeStatusFailed    V2SignalsAccountActivityLoginDecisionMfaChallengeStatus = "failed"
	V2SignalsAccountActivityLoginDecisionMfaChallengeStatusOther     V2SignalsAccountActivityLoginDecisionMfaChallengeStatus = "other"
	V2SignalsAccountActivityLoginDecisionMfaChallengeStatusSucceeded V2SignalsAccountActivityLoginDecisionMfaChallengeStatus = "succeeded"
)

// The action the merchant took following the evaluation.
type V2SignalsAccountActivityLoginDecisionStatus string

// List of values that V2SignalsAccountActivityLoginDecisionStatus can take
const (
	V2SignalsAccountActivityLoginDecisionStatusAllowed    V2SignalsAccountActivityLoginDecisionStatus = "allowed"
	V2SignalsAccountActivityLoginDecisionStatusBlocked    V2SignalsAccountActivityLoginDecisionStatus = "blocked"
	V2SignalsAccountActivityLoginDecisionStatusRestricted V2SignalsAccountActivityLoginDecisionStatus = "restricted"
)

// The authentication method added to the account.
type V2SignalsAccountActivityMfaAdditionMfaType string

// List of values that V2SignalsAccountActivityMfaAdditionMfaType can take
const (
	V2SignalsAccountActivityMfaAdditionMfaTypeAuthenticatorApp V2SignalsAccountActivityMfaAdditionMfaType = "authenticator_app"
	V2SignalsAccountActivityMfaAdditionMfaTypeEmail            V2SignalsAccountActivityMfaAdditionMfaType = "email"
	V2SignalsAccountActivityMfaAdditionMfaTypeHardwareKey      V2SignalsAccountActivityMfaAdditionMfaType = "hardware_key"
	V2SignalsAccountActivityMfaAdditionMfaTypeOther            V2SignalsAccountActivityMfaAdditionMfaType = "other"
	V2SignalsAccountActivityMfaAdditionMfaTypeSms              V2SignalsAccountActivityMfaAdditionMfaType = "sms"
)

// The authentication method removed from the account.
type V2SignalsAccountActivityMfaRemovalMfaType string

// List of values that V2SignalsAccountActivityMfaRemovalMfaType can take
const (
	V2SignalsAccountActivityMfaRemovalMfaTypeAuthenticatorApp V2SignalsAccountActivityMfaRemovalMfaType = "authenticator_app"
	V2SignalsAccountActivityMfaRemovalMfaTypeEmail            V2SignalsAccountActivityMfaRemovalMfaType = "email"
	V2SignalsAccountActivityMfaRemovalMfaTypeHardwareKey      V2SignalsAccountActivityMfaRemovalMfaType = "hardware_key"
	V2SignalsAccountActivityMfaRemovalMfaTypeOther            V2SignalsAccountActivityMfaRemovalMfaType = "other"
	V2SignalsAccountActivityMfaRemovalMfaTypeSms              V2SignalsAccountActivityMfaRemovalMfaType = "sms"
)

// The action the merchant took following the evaluation.
type V2SignalsAccountActivityRegistrationDecisionStatus string

// List of values that V2SignalsAccountActivityRegistrationDecisionStatus can take
const (
	V2SignalsAccountActivityRegistrationDecisionStatusAllowed    V2SignalsAccountActivityRegistrationDecisionStatus = "allowed"
	V2SignalsAccountActivityRegistrationDecisionStatusBlocked    V2SignalsAccountActivityRegistrationDecisionStatus = "blocked"
	V2SignalsAccountActivityRegistrationDecisionStatusRestricted V2SignalsAccountActivityRegistrationDecisionStatus = "restricted"
)

// The type of activity.
type V2SignalsAccountActivityType string

// List of values that V2SignalsAccountActivityType can take
const (
	V2SignalsAccountActivityTypeAccountRestricted           V2SignalsAccountActivityType = "account_restricted"
	V2SignalsAccountActivityTypeAccountReviewed             V2SignalsAccountActivityType = "account_reviewed"
	V2SignalsAccountActivityTypeAccountSuspended            V2SignalsAccountActivityType = "account_suspended"
	V2SignalsAccountActivityTypeAccountTakeoverConfirmation V2SignalsAccountActivityType = "account_takeover_confirmation"
	V2SignalsAccountActivityTypeAccountTakeoverMitigation   V2SignalsAccountActivityType = "account_takeover_mitigation"
	V2SignalsAccountActivityTypeAccountTakeoverReport       V2SignalsAccountActivityType = "account_takeover_report"
	V2SignalsAccountActivityTypeCredentialChange            V2SignalsAccountActivityType = "credential_change"
	V2SignalsAccountActivityTypeCustomMarketingAction       V2SignalsAccountActivityType = "custom_marketing_action"
	V2SignalsAccountActivityTypeLoginAttempt                V2SignalsAccountActivityType = "login_attempt"
	V2SignalsAccountActivityTypeLoginDecision               V2SignalsAccountActivityType = "login_decision"
	V2SignalsAccountActivityTypeMfaAddition                 V2SignalsAccountActivityType = "mfa_addition"
	V2SignalsAccountActivityTypeMfaRemoval                  V2SignalsAccountActivityType = "mfa_removal"
	V2SignalsAccountActivityTypeRegistrationAttempt         V2SignalsAccountActivityType = "registration_attempt"
	V2SignalsAccountActivityTypeRegistrationDecision        V2SignalsAccountActivityType = "registration_decision"
)

// Account profile data.
type V2SignalsAccountActivityAccountDetailsDataDefaultsProfile struct {
	// The business URL.
	BusinessURL string `json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs string `json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription string `json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountActivityAccountDetailsDataDefaults struct {
	// Account profile data.
	Profile *V2SignalsAccountActivityAccountDetailsDataDefaultsProfile `json:"profile"`
}

// Business details for identity data.
type V2SignalsAccountActivityAccountDetailsDataIdentityBusinessDetails struct {
	// Registered business name.
	RegisteredName string `json:"registered_name,omitempty"`
}

// Identity data.
type V2SignalsAccountActivityAccountDetailsDataIdentity struct {
	// Business details for identity data.
	BusinessDetails *V2SignalsAccountActivityAccountDetailsDataIdentityBusinessDetails `json:"business_details"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountActivityAccountDetailsData struct {
	// The account's contact email.
	ContactEmail string `json:"contact_email,omitempty"`
	// Default account settings.
	Defaults *V2SignalsAccountActivityAccountDetailsDataDefaults `json:"defaults,omitempty"`
	// Identity data.
	Identity *V2SignalsAccountActivityAccountDetailsDataIdentity `json:"identity,omitempty"`
}

// The account, customer, or inline account data associated with the activity.
type V2SignalsAccountActivityAccountDetails struct {
	// The v2 account ID of the account.
	Account string `json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2 accounts.
	Customer string `json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountActivityAccountDetailsData `json:"data,omitempty"`
}

// Details for the account restriction. Present only when type is account_restricted. The activity
// requires an existing account_details.account or account_details.customer; inline data is unsupported.
type V2SignalsAccountActivityAccountRestricted struct {
	// The reason the account or customer was restricted.
	Reason V2SignalsAccountActivityAccountRestrictedReason `json:"reason"`
}

// Details for the account review. Present only when type is account_reviewed. The activity
// requires an existing account_details.account or account_details.customer; inline data is unsupported.
type V2SignalsAccountActivityAccountReviewed struct {
	// The outcome of the merchant review.
	Outcome V2SignalsAccountActivityAccountReviewedOutcome `json:"outcome"`
}

// Details for the account suspension. Present only when type is account_suspended. The activity
// requires an existing account_details.customer; account_details.account and inline data are unsupported.
type V2SignalsAccountActivityAccountSuspended struct {
	// The reason the customer was suspended.
	Reason V2SignalsAccountActivityAccountSuspendedReason `json:"reason"`
}

// Details for the merchant's takeover confirmation. Present only when type is account_takeover_confirmation.
type V2SignalsAccountActivityAccountTakeoverConfirmation struct {
	// The type of credential involved in the confirmed takeover.
	CredentialType V2SignalsAccountActivityAccountTakeoverConfirmationCredentialType `json:"credential_type"`
}

// Details for the account holder's takeover report. Present only when type is account_takeover_report.
type V2SignalsAccountActivityAccountTakeoverReport struct {
	// The type of credential involved in the reported takeover.
	CredentialType V2SignalsAccountActivityAccountTakeoverReportCredentialType `json:"credential_type"`
}

// Details for the credential change. Present only when type is credential_change.
type V2SignalsAccountActivityCredentialChange struct {
	// The type of credential that changed.
	CredentialType V2SignalsAccountActivityCredentialChangeCredentialType `json:"credential_type"`
}

// Details for the custom marketing action. Present only when type is custom_marketing_action.
type V2SignalsAccountActivityCustomMarketingAction struct {
	// A description of the action.
	Description string `json:"description"`
	// The outcome of the action.
	Outcome V2SignalsAccountActivityCustomMarketingActionOutcome `json:"outcome"`
	// Additional user-defined properties of the action.
	Properties map[string]string `json:"properties"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityLoginAttemptClientDetailsData struct {
	// The IP address associated with the activity.
	IP string `json:"ip"`
	// The referrer associated with the activity.
	Referrer string `json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent string `json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityLoginAttemptClientDetails struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityLoginAttemptClientDetailsData `json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession string `json:"radar_session,omitempty"`
}

// Details for the login attempt. Present only when type is login_attempt.
type V2SignalsAccountActivityLoginAttempt struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityLoginAttemptClientDetails `json:"client_details"`
}

// Multi-factor authentication challenges performed during the login process.
type V2SignalsAccountActivityLoginDecisionMfaChallenge struct {
	// The authentication method used for the challenge.
	MfaType V2SignalsAccountActivityLoginDecisionMfaChallengeMfaType `json:"mfa_type"`
	// The outcome of the authentication challenge.
	Status V2SignalsAccountActivityLoginDecisionMfaChallengeStatus `json:"status"`
}

// Details for the login decision. Present only when type is login_decision.
type V2SignalsAccountActivityLoginDecision struct {
	// Multi-factor authentication challenges performed during the login process.
	MfaChallenges []*V2SignalsAccountActivityLoginDecisionMfaChallenge `json:"mfa_challenges,omitempty"`
	// The action the merchant took following the evaluation.
	Status V2SignalsAccountActivityLoginDecisionStatus `json:"status"`
}

// Details for the MFA addition. Present only when type is mfa_addition.
type V2SignalsAccountActivityMfaAddition struct {
	// The authentication method added to the account.
	MfaType V2SignalsAccountActivityMfaAdditionMfaType `json:"mfa_type"`
}

// Details for the MFA removal. Present only when type is mfa_removal.
type V2SignalsAccountActivityMfaRemoval struct {
	// The authentication method removed from the account.
	MfaType V2SignalsAccountActivityMfaRemovalMfaType `json:"mfa_type"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityRegistrationAttemptClientDetailsData struct {
	// The IP address associated with the activity.
	IP string `json:"ip"`
	// The referrer associated with the activity.
	Referrer string `json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent string `json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityRegistrationAttemptClientDetails struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityRegistrationAttemptClientDetailsData `json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession string `json:"radar_session,omitempty"`
}

// Details for the registration attempt. Present only when type is registration_attempt.
type V2SignalsAccountActivityRegistrationAttempt struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityRegistrationAttemptClientDetails `json:"client_details"`
}

// Details for the registration decision. Present only when type is registration_decision.
type V2SignalsAccountActivityRegistrationDecision struct {
	// The action the merchant took following the evaluation.
	Status V2SignalsAccountActivityRegistrationDecisionStatus `json:"status"`
}

// Account Activity resource for the Signals API.
type V2SignalsAccountActivity struct {
	APIResource
	// The account, customer, or inline account data associated with the activity.
	AccountDetails *V2SignalsAccountActivityAccountDetails `json:"account_details,omitempty"`
	// The account evaluation this activity is associated with, when applicable.
	AccountEvaluation string `json:"account_evaluation,omitempty"`
	// Details for the account restriction. Present only when type is account_restricted. The activity
	// requires an existing account_details.account or account_details.customer; inline data is unsupported.
	AccountRestricted *V2SignalsAccountActivityAccountRestricted `json:"account_restricted,omitempty"`
	// Details for the account review. Present only when type is account_reviewed. The activity
	// requires an existing account_details.account or account_details.customer; inline data is unsupported.
	AccountReviewed *V2SignalsAccountActivityAccountReviewed `json:"account_reviewed,omitempty"`
	// Details for the account suspension. Present only when type is account_suspended. The activity
	// requires an existing account_details.customer; account_details.account and inline data are unsupported.
	AccountSuspended *V2SignalsAccountActivityAccountSuspended `json:"account_suspended,omitempty"`
	// Details for the merchant's takeover confirmation. Present only when type is account_takeover_confirmation.
	AccountTakeoverConfirmation *V2SignalsAccountActivityAccountTakeoverConfirmation `json:"account_takeover_confirmation,omitempty"`
	// Details for the account holder's takeover report. Present only when type is account_takeover_report.
	AccountTakeoverReport *V2SignalsAccountActivityAccountTakeoverReport `json:"account_takeover_report,omitempty"`
	// Timestamp at which the account activity was created.
	Created time.Time `json:"created"`
	// Details for the credential change. Present only when type is credential_change.
	CredentialChange *V2SignalsAccountActivityCredentialChange `json:"credential_change,omitempty"`
	// Details for the custom marketing action. Present only when type is custom_marketing_action.
	CustomMarketingAction *V2SignalsAccountActivityCustomMarketingAction `json:"custom_marketing_action,omitempty"`
	// Unique identifier for the account activity.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Details for the login attempt. Present only when type is login_attempt.
	LoginAttempt *V2SignalsAccountActivityLoginAttempt `json:"login_attempt,omitempty"`
	// Details for the login decision. Present only when type is login_decision.
	LoginDecision *V2SignalsAccountActivityLoginDecision `json:"login_decision,omitempty"`
	// Additional information about the activity.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Details for the MFA addition. Present only when type is mfa_addition.
	MfaAddition *V2SignalsAccountActivityMfaAddition `json:"mfa_addition,omitempty"`
	// Details for the MFA removal. Present only when type is mfa_removal.
	MfaRemoval *V2SignalsAccountActivityMfaRemoval `json:"mfa_removal,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt time.Time `json:"occurred_at"`
	// Details for the registration attempt. Present only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountActivityRegistrationAttempt `json:"registration_attempt,omitempty"`
	// Details for the registration decision. Present only when type is registration_decision.
	RegistrationDecision *V2SignalsAccountActivityRegistrationDecision `json:"registration_decision,omitempty"`
	// The AccountSignal that prompted this activity, when applicable. This historical reference
	// may coexist with account_evaluation and remains present after the signal expires.
	TriggeringSignal string `json:"triggering_signal,omitempty"`
	// The type of activity.
	Type V2SignalsAccountActivityType `json:"type"`
}
