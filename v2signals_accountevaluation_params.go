//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the login attempt. Provide only when type is login_attempt.
type V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the registration attempt. Provide only when type is registration_attempt.
type V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Inline activity data used to create a new account activity for the evaluation.
type V2SignalsAccountEvaluationAccountActivityDetailsDataParams struct {
	// Details for the login attempt. Provide only when type is login_attempt.
	LoginAttempt *V2SignalsAccountEvaluationAccountActivityDetailsDataLoginAttemptParams `form:"login_attempt" json:"login_attempt,omitempty"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt *time.Time `form:"occurred_at" json:"occurred_at,omitempty"`
	// Details for the registration attempt. Provide only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountEvaluationAccountActivityDetailsDataRegistrationAttemptParams `form:"registration_attempt" json:"registration_attempt,omitempty"`
	// The type of activity. Must be registration_attempt or login_attempt.
	Type *string `form:"type" json:"type"`
}

// Account activity to record alongside this evaluation.
type V2SignalsAccountEvaluationAccountActivityDetailsParams struct {
	// The ID of an existing account activity to associate with the evaluation.
	AccountActivity *string `form:"account_activity" json:"account_activity,omitempty"`
	// Inline activity data used to create a new account activity for the evaluation.
	Data *V2SignalsAccountEvaluationAccountActivityDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Account profile data.
type V2SignalsAccountEvaluationAccountDetailsDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountEvaluationAccountDetailsDataDefaultsParams struct {
	// Account profile data.
	Profile *V2SignalsAccountEvaluationAccountDetailsDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Business details for identity data.
type V2SignalsAccountEvaluationAccountDetailsDataIdentityBusinessDetailsParams struct {
	// Registered business name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Identity data.
type V2SignalsAccountEvaluationAccountDetailsDataIdentityParams struct {
	// Business details for identity data.
	BusinessDetails *V2SignalsAccountEvaluationAccountDetailsDataIdentityBusinessDetailsParams `form:"business_details" json:"business_details"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountEvaluationAccountDetailsDataParams struct {
	// Default account settings.
	Defaults *V2SignalsAccountEvaluationAccountDetailsDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// Identity data.
	Identity *V2SignalsAccountEvaluationAccountDetailsDataIdentityParams `form:"identity" json:"identity,omitempty"`
}

// The account, customer, or inline account data to evaluate.
type V2SignalsAccountEvaluationAccountDetailsParams struct {
	// The v2 account ID of the account.
	Account *string `form:"account" json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountEvaluationAccountDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Creates a new account evaluation to request signal evaluations on an account, customer, or inline account data.
type V2SignalsAccountEvaluationParams struct {
	Params `form:"*"`
	// Account activity to record alongside this evaluation.
	AccountActivityDetails *V2SignalsAccountEvaluationAccountActivityDetailsParams `form:"account_activity_details" json:"account_activity_details,omitempty"`
	// The account, customer, or inline account data to evaluate.
	AccountDetails *V2SignalsAccountEvaluationAccountDetailsParams `form:"account_details" json:"account_details,omitempty"`
	// List of signals to evaluate.
	RequestedSignals []*string `form:"requested_signals" json:"requested_signals,omitempty"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the login attempt. Provide only when type is login_attempt.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Raw client details for the activity, when a Radar session is not available.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptClientDetailsDataParams struct {
	// The IP address associated with the activity.
	IP *string `form:"ip" json:"ip"`
	// The referrer associated with the activity.
	Referrer *string `form:"referrer" json:"referrer,omitempty"`
	// The user agent associated with the activity.
	UserAgent *string `form:"user_agent" json:"user_agent,omitempty"`
}

// Client details captured for the attempt.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptClientDetailsParams struct {
	// Raw client details for the activity, when a Radar session is not available.
	Data *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptClientDetailsDataParams `form:"data" json:"data,omitempty"`
	// The Radar session ID capturing client details for the activity.
	RadarSession *string `form:"radar_session" json:"radar_session,omitempty"`
}

// Details for the registration attempt. Provide only when type is registration_attempt.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptParams struct {
	// Client details captured for the attempt.
	ClientDetails *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptClientDetailsParams `form:"client_details" json:"client_details"`
}

// Inline activity data used to create a new account activity for the evaluation.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsDataParams struct {
	// Details for the login attempt. Provide only when type is login_attempt.
	LoginAttempt *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataLoginAttemptParams `form:"login_attempt" json:"login_attempt,omitempty"`
	// Timestamp at which the activity occurred. Defaults to the created time if not provided.
	OccurredAt *time.Time `form:"occurred_at" json:"occurred_at,omitempty"`
	// Details for the registration attempt. Provide only when type is registration_attempt.
	RegistrationAttempt *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataRegistrationAttemptParams `form:"registration_attempt" json:"registration_attempt,omitempty"`
	// The type of activity. Must be registration_attempt or login_attempt.
	Type *string `form:"type" json:"type"`
}

// Account activity to record alongside this evaluation.
type V2SignalsAccountEvaluationCreateAccountActivityDetailsParams struct {
	// The ID of an existing account activity to associate with the evaluation.
	AccountActivity *string `form:"account_activity" json:"account_activity,omitempty"`
	// Inline activity data used to create a new account activity for the evaluation.
	Data *V2SignalsAccountEvaluationCreateAccountActivityDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Account profile data.
type V2SignalsAccountEvaluationCreateAccountDetailsDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2SignalsAccountEvaluationCreateAccountDetailsDataDefaultsParams struct {
	// Account profile data.
	Profile *V2SignalsAccountEvaluationCreateAccountDetailsDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Business details for identity data.
type V2SignalsAccountEvaluationCreateAccountDetailsDataIdentityBusinessDetailsParams struct {
	// Registered business name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Identity data.
type V2SignalsAccountEvaluationCreateAccountDetailsDataIdentityParams struct {
	// Business details for identity data.
	BusinessDetails *V2SignalsAccountEvaluationCreateAccountDetailsDataIdentityBusinessDetailsParams `form:"business_details" json:"business_details"`
}

// Inline account data to evaluate without creating a v2 account.
type V2SignalsAccountEvaluationCreateAccountDetailsDataParams struct {
	// Default account settings.
	Defaults *V2SignalsAccountEvaluationCreateAccountDetailsDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// Identity data.
	Identity *V2SignalsAccountEvaluationCreateAccountDetailsDataIdentityParams `form:"identity" json:"identity,omitempty"`
}

// The account, customer, or inline account data to evaluate.
type V2SignalsAccountEvaluationCreateAccountDetailsParams struct {
	// The v2 account ID of the account.
	Account *string `form:"account" json:"account,omitempty"`
	// The v1 customer ID of the account, for users not yet migrated to v2/accounts.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Inline account data to evaluate without creating a v2 account.
	Data *V2SignalsAccountEvaluationCreateAccountDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Creates a new account evaluation to request signal evaluations on an account, customer, or inline account data.
type V2SignalsAccountEvaluationCreateParams struct {
	Params `form:"*"`
	// Account activity to record alongside this evaluation.
	AccountActivityDetails *V2SignalsAccountEvaluationCreateAccountActivityDetailsParams `form:"account_activity_details" json:"account_activity_details,omitempty"`
	// The account, customer, or inline account data to evaluate.
	AccountDetails *V2SignalsAccountEvaluationCreateAccountDetailsParams `form:"account_details" json:"account_details"`
	// List of signals to evaluate.
	RequestedSignals []*string `form:"requested_signals" json:"requested_signals"`
}

// Retrieves an AccountEvaluation by its ID.
type V2SignalsAccountEvaluationRetrieveParams struct {
	Params `form:"*"`
}
