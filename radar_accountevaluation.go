//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves an AccountEvaluation object.
type RadarAccountEvaluationParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Event payload for login_failed.
	LoginFailed *RadarAccountEvaluationLoginFailedParams `form:"login_failed"`
	// Event payload for login_initiated.
	LoginInitiated *RadarAccountEvaluationLoginInitiatedParams `form:"login_initiated"`
	// Event payload for registration_failed.
	RegistrationFailed *RadarAccountEvaluationRegistrationFailedParams `form:"registration_failed"`
	// Event payload for registration_initiated.
	RegistrationInitiated *RadarAccountEvaluationRegistrationInitiatedParams `form:"registration_initiated"`
	// The type of evaluation requested.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *RadarAccountEvaluationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Client device metadata details (e.g., radar_session).
type RadarAccountEvaluationLoginInitiatedClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session associated with the account evaluation.
	RadarSession *string `form:"radar_session"`
}

// Event payload for login_initiated.
type RadarAccountEvaluationLoginInitiatedParams struct {
	// Client device metadata details (e.g., radar_session).
	ClientDeviceMetadataDetails *RadarAccountEvaluationLoginInitiatedClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Stripe customer ID
	Customer *string `form:"customer"`
}

// Client device metadata details (e.g., radar_session).
type RadarAccountEvaluationRegistrationInitiatedClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session associated with the account evaluation.
	RadarSession *string `form:"radar_session"`
}

// Customer data
type RadarAccountEvaluationRegistrationInitiatedCustomerDataParams struct {
	// Customer email
	Email *string `form:"email"`
	// Customer name
	Name *string `form:"name"`
	// Customer phone
	Phone *string `form:"phone"`
}

// Event payload for registration_initiated.
type RadarAccountEvaluationRegistrationInitiatedParams struct {
	// Client device metadata details (e.g., radar_session).
	ClientDeviceMetadataDetails *RadarAccountEvaluationRegistrationInitiatedClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Stripe customer ID
	Customer *string `form:"customer"`
	// Customer data
	CustomerData *RadarAccountEvaluationRegistrationInitiatedCustomerDataParams `form:"customer_data"`
}

// Event payload for login_failed.
type RadarAccountEvaluationLoginFailedParams struct {
	// The reason why this login failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_failed.
type RadarAccountEvaluationRegistrationFailedParams struct {
	// The reason why this registration failed.
	Reason *string `form:"reason"`
}

// Retrieves an AccountEvaluation object.
type RadarAccountEvaluationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *RadarAccountEvaluationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Client device metadata details (e.g., radar_session).
type RadarAccountEvaluationCreateLoginInitiatedClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session associated with the account evaluation.
	RadarSession *string `form:"radar_session"`
}

// Event payload for login_initiated.
type RadarAccountEvaluationCreateLoginInitiatedParams struct {
	// Client device metadata details (e.g., radar_session).
	ClientDeviceMetadataDetails *RadarAccountEvaluationCreateLoginInitiatedClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Stripe customer ID
	Customer *string `form:"customer"`
}

// Client device metadata details (e.g., radar_session).
type RadarAccountEvaluationCreateRegistrationInitiatedClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session associated with the account evaluation.
	RadarSession *string `form:"radar_session"`
}

// Customer data
type RadarAccountEvaluationCreateRegistrationInitiatedCustomerDataParams struct {
	// Customer email
	Email *string `form:"email"`
	// Customer name
	Name *string `form:"name"`
	// Customer phone
	Phone *string `form:"phone"`
}

// Event payload for registration_initiated.
type RadarAccountEvaluationCreateRegistrationInitiatedParams struct {
	// Client device metadata details (e.g., radar_session).
	ClientDeviceMetadataDetails *RadarAccountEvaluationCreateRegistrationInitiatedClientDeviceMetadataDetailsParams `form:"client_device_metadata_details"`
	// Stripe customer ID
	Customer *string `form:"customer"`
	// Customer data
	CustomerData *RadarAccountEvaluationCreateRegistrationInitiatedCustomerDataParams `form:"customer_data"`
}

// Creates a new AccountEvaluation object.
type RadarAccountEvaluationCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Event payload for login_initiated.
	LoginInitiated *RadarAccountEvaluationCreateLoginInitiatedParams `form:"login_initiated"`
	// Event payload for registration_initiated.
	RegistrationInitiated *RadarAccountEvaluationCreateRegistrationInitiatedParams `form:"registration_initiated"`
	// The type of evaluation requested.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *RadarAccountEvaluationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Event payload for login_failed.
type RadarAccountEvaluationUpdateLoginFailedParams struct {
	// The reason why this login failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_failed.
type RadarAccountEvaluationUpdateRegistrationFailedParams struct {
	// The reason why this registration failed.
	Reason *string `form:"reason"`
}

// Reports an event on an AccountEvaluation object.
type RadarAccountEvaluationUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Event payload for login_failed.
	LoginFailed *RadarAccountEvaluationUpdateLoginFailedParams `form:"login_failed"`
	// Event payload for registration_failed.
	RegistrationFailed *RadarAccountEvaluationUpdateRegistrationFailedParams `form:"registration_failed"`
	// The type of event to report.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *RadarAccountEvaluationUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Data about a failed login event.
type RadarAccountEvaluationEventLoginFailed struct {
	// The reason why this login failed.
	Reason string `json:"reason"`
}

// Data about a failed registration event.
type RadarAccountEvaluationEventRegistrationFailed struct {
	// The reason why this registration failed.
	Reason string `json:"reason"`
}

// The list of events that were reported for this Account Evaluation object via the report API.
type RadarAccountEvaluationEvent struct {
	// Data about a failed login event.
	LoginFailed *RadarAccountEvaluationEventLoginFailed `json:"login_failed"`
	// Time at which the event occurred. Measured in seconds since the Unix epoch.
	OccurredAt int64 `json:"occurred_at"`
	// Data about a failed registration event.
	RegistrationFailed *RadarAccountEvaluationEventRegistrationFailed `json:"registration_failed"`
	// The type of event that occurred.
	Type string `json:"type"`
}
type RadarAccountEvaluationSignalsAccountSharing struct {
	// Score for this signal (float between 0.0-100.0).
	Score float64 `json:"score"`
}
type RadarAccountEvaluationSignalsMultiAccounting struct {
	// Score for this signal (float between 0.0-100.0).
	Score float64 `json:"score"`
}

// A hash of signal objects providing Radar's evaluation for the lifecycle event.
type RadarAccountEvaluationSignals struct {
	AccountSharing  *RadarAccountEvaluationSignalsAccountSharing  `json:"account_sharing"`
	MultiAccounting *RadarAccountEvaluationSignalsMultiAccounting `json:"multi_accounting"`
}

// Account Evaluation resource returned by the Radar Account Evaluations API.
type RadarAccountEvaluation struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// The ID of the Stripe customer the account evaluation is associated with.
	Customer string `json:"customer"`
	// The list of events that were reported for this Account Evaluation object via the report API.
	Events []*RadarAccountEvaluationEvent `json:"events"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A hash of signal objects providing Radar's evaluation for the lifecycle event.
	Signals *RadarAccountEvaluationSignals `json:"signals"`
	// The type of evaluation returned, based on the user's request.
	Type string `json:"type"`
}
