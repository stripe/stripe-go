//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The action the merchant took following the evaluation.
type V2SignalsAccountActivityLoginDecisionStatus string

// List of values that V2SignalsAccountActivityLoginDecisionStatus can take
const (
	V2SignalsAccountActivityLoginDecisionStatusAllowed    V2SignalsAccountActivityLoginDecisionStatus = "allowed"
	V2SignalsAccountActivityLoginDecisionStatusBlocked    V2SignalsAccountActivityLoginDecisionStatus = "blocked"
	V2SignalsAccountActivityLoginDecisionStatusRestricted V2SignalsAccountActivityLoginDecisionStatus = "restricted"
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
	V2SignalsAccountActivityTypeLoginAttempt         V2SignalsAccountActivityType = "login_attempt"
	V2SignalsAccountActivityTypeLoginDecision        V2SignalsAccountActivityType = "login_decision"
	V2SignalsAccountActivityTypeRegistrationAttempt  V2SignalsAccountActivityType = "registration_attempt"
	V2SignalsAccountActivityTypeRegistrationDecision V2SignalsAccountActivityType = "registration_decision"
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

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountActivityAccountDetailsData struct {
	// Default account settings.
	Defaults *V2SignalsAccountActivityAccountDetailsDataDefaults `json:"defaults,omitempty"`
}

// The account, customer, or inline account data associated with the activity.
type V2SignalsAccountActivityAccountDetails struct {
	// The v2 account ID of the account.
	Account string `json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer string `json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountActivityAccountDetailsData `json:"data,omitempty"`
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

// Details for the login decision. Present only when type is login_decision.
type V2SignalsAccountActivityLoginDecision struct {
	// The action the merchant took following the evaluation.
	Status V2SignalsAccountActivityLoginDecisionStatus `json:"status"`
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
	// Timestamp at which the account activity was created.
	Created time.Time `json:"created"`
	// Unique identifier for the account activity.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Details for the login attempt. Present only when type is login_attempt.
	LoginAttempt *V2SignalsAccountActivityLoginAttempt `json:"login_attempt,omitempty"`
	// Details for the login decision. Present only when type is login_decision.
	LoginDecision *V2SignalsAccountActivityLoginDecision `json:"login_decision,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt time.Time `json:"occurred_at"`
	// Details for the registration attempt. Present only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountActivityRegistrationAttempt `json:"registration_attempt,omitempty"`
	// Details for the registration decision. Present only when type is registration_decision.
	RegistrationDecision *V2SignalsAccountActivityRegistrationDecision `json:"registration_decision,omitempty"`
	// The type of activity.
	Type V2SignalsAccountActivityType `json:"type"`
}
