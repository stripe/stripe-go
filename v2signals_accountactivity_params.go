//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Account profile data.
type V2SignalsAccountActivityAccountDetailsDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountActivityAccountDetailsDataDefaultsParams struct {
	// Account profile data.
	Profile *V2SignalsAccountActivityAccountDetailsDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountActivityAccountDetailsDataParams struct {
	// Default account settings.
	Defaults *V2SignalsAccountActivityAccountDetailsDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
}

// The account, customer, or inline account data associated with the activity.
type V2SignalsAccountActivityAccountDetailsParams struct {
	// The v2 account ID of the account.
	Account *string `form:"account" json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountActivityAccountDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityLoginAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityLoginAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityLoginAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the login attempt. Provide only when type is login_attempt.
type V2SignalsAccountActivityLoginAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityLoginAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Details for the login decision. Provide only when type is login_decision.
type V2SignalsAccountActivityLoginDecisionParams struct {
	// The action the merchant took following the evaluation.
	Status *string `form:"status" json:"status"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityRegistrationAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityRegistrationAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityRegistrationAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the registration attempt. Provide only when type is registration_attempt.
type V2SignalsAccountActivityRegistrationAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityRegistrationAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Details for the registration decision. Provide only when type is registration_decision.
type V2SignalsAccountActivityRegistrationDecisionParams struct {
	// The action the merchant took following the evaluation.
	Status *string `form:"status" json:"status"`
}

// Creates a new account activity to report account registration, login, or evaluation follow-up activity.
type V2SignalsAccountActivityParams struct {
	Params `form:"*"`
	// The account, customer, or inline account data associated with the activity.
	AccountDetails *V2SignalsAccountActivityAccountDetailsParams `form:"account_details" json:"account_details,omitempty"`
	// The account evaluation this activity is associated with, when applicable.
	AccountEvaluation *string `form:"account_evaluation" json:"account_evaluation,omitempty"`
	// Details for the login attempt. Provide only when type is login_attempt.
	LoginAttempt *V2SignalsAccountActivityLoginAttemptParams `form:"login_attempt" json:"login_attempt,omitempty"`
	// Details for the login decision. Provide only when type is login_decision.
	LoginDecision *V2SignalsAccountActivityLoginDecisionParams `form:"login_decision" json:"login_decision,omitempty"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt *time.Time `form:"occurred_at" json:"occurred_at,omitempty"`
	// Details for the registration attempt. Provide only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountActivityRegistrationAttemptParams `form:"registration_attempt" json:"registration_attempt,omitempty"`
	// Details for the registration decision. Provide only when type is registration_decision.
	RegistrationDecision *V2SignalsAccountActivityRegistrationDecisionParams `form:"registration_decision" json:"registration_decision,omitempty"`
	// The type of activity.
	Type *string `form:"type" json:"type,omitempty"`
}

// Account profile data.
type V2SignalsAccountActivityCreateAccountDetailsDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountActivityCreateAccountDetailsDataDefaultsParams struct {
	// Account profile data.
	Profile *V2SignalsAccountActivityCreateAccountDetailsDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountActivityCreateAccountDetailsDataParams struct {
	// Default account settings.
	Defaults *V2SignalsAccountActivityCreateAccountDetailsDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
}

// The account, customer, or inline account data associated with the activity.
type V2SignalsAccountActivityCreateAccountDetailsParams struct {
	// The v2 account ID of the account.
	Account *string `form:"account" json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountActivityCreateAccountDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityCreateLoginAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityCreateLoginAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityCreateLoginAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the login attempt. Provide only when type is login_attempt.
type V2SignalsAccountActivityCreateLoginAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityCreateLoginAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Details for the login decision. Provide only when type is login_decision.
type V2SignalsAccountActivityCreateLoginDecisionParams struct {
	// The action the merchant took following the evaluation.
	Status *string `form:"status" json:"status"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountActivityCreateRegistrationAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountActivityCreateRegistrationAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountActivityCreateRegistrationAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the registration attempt. Provide only when type is registration_attempt.
type V2SignalsAccountActivityCreateRegistrationAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountActivityCreateRegistrationAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Details for the registration decision. Provide only when type is registration_decision.
type V2SignalsAccountActivityCreateRegistrationDecisionParams struct {
	// The action the merchant took following the evaluation.
	Status *string `form:"status" json:"status"`
}

// Creates a new account activity to report account registration, login, or evaluation follow-up activity.
type V2SignalsAccountActivityCreateParams struct {
	Params `form:"*"`
	// The account, customer, or inline account data associated with the activity.
	AccountDetails *V2SignalsAccountActivityCreateAccountDetailsParams `form:"account_details" json:"account_details,omitempty"`
	// The account evaluation this activity is associated with, when applicable.
	AccountEvaluation *string `form:"account_evaluation" json:"account_evaluation,omitempty"`
	// Details for the login attempt. Provide only when type is login_attempt.
	LoginAttempt *V2SignalsAccountActivityCreateLoginAttemptParams `form:"login_attempt" json:"login_attempt,omitempty"`
	// Details for the login decision. Provide only when type is login_decision.
	LoginDecision *V2SignalsAccountActivityCreateLoginDecisionParams `form:"login_decision" json:"login_decision,omitempty"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt *time.Time `form:"occurred_at" json:"occurred_at,omitempty"`
	// Details for the registration attempt. Provide only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountActivityCreateRegistrationAttemptParams `form:"registration_attempt" json:"registration_attempt,omitempty"`
	// Details for the registration decision. Provide only when type is registration_decision.
	RegistrationDecision *V2SignalsAccountActivityCreateRegistrationDecisionParams `form:"registration_decision" json:"registration_decision,omitempty"`
	// The type of activity.
	Type *string `form:"type" json:"type"`
}

// Deletes an AccountActivity by its ID.
type V2SignalsAccountActivityDeleteParams struct {
	Params `form:"*"`
}

// Retrieves an AccountActivity by its ID.
type V2SignalsAccountActivityRetrieveParams struct {
	Params `form:"*"`
}
