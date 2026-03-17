//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Client details context.
type RadarCustomerEvaluationEvaluationContextClientDetailsParams struct {
	// ID for the Radar Session associated with the customer evaluation.
	RadarSession *string `form:"radar_session"`
}

// Customer data
type RadarCustomerEvaluationEvaluationContextCustomerDetailsCustomerDataParams struct {
	// Customer email
	Email *string `form:"email"`
	// Customer name
	Name *string `form:"name"`
	// Customer phone
	Phone *string `form:"phone"`
}

// Customer details context.
type RadarCustomerEvaluationEvaluationContextCustomerDetailsParams struct {
	// Stripe customer ID
	Customer *string `form:"customer"`
	// Customer data
	CustomerData *RadarCustomerEvaluationEvaluationContextCustomerDetailsCustomerDataParams `form:"customer_data"`
}

// Array of context entries for the evaluation.
type RadarCustomerEvaluationEvaluationContextParams struct {
	// Client details context.
	ClientDetails *RadarCustomerEvaluationEvaluationContextClientDetailsParams `form:"client_details"`
	// Customer details context.
	CustomerDetails *RadarCustomerEvaluationEvaluationContextCustomerDetailsParams `form:"customer_details"`
	// The type of context entry.
	Type *string `form:"type"`
}

// Creates a new CustomerEvaluation object.
type RadarCustomerEvaluationParams struct {
	Params `form:"*"`
	// Array of context entries for the evaluation.
	EvaluationContext []*RadarCustomerEvaluationEvaluationContextParams `form:"evaluation_context"`
	// The type of evaluation event.
	EventType *string `form:"event_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Event payload for login_failed.
	LoginFailed *RadarCustomerEvaluationLoginFailedParams `form:"login_failed"`
	// Event payload for registration_failed.
	RegistrationFailed *RadarCustomerEvaluationRegistrationFailedParams `form:"registration_failed"`
	// Event payload for registration_success.
	RegistrationSuccess *RadarCustomerEvaluationRegistrationSuccessParams `form:"registration_success"`
	// The type of event to report.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *RadarCustomerEvaluationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Event payload for login_failed.
type RadarCustomerEvaluationLoginFailedParams struct {
	// The reason why this login failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_failed.
type RadarCustomerEvaluationRegistrationFailedParams struct {
	// The reason why this registration failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_success.
type RadarCustomerEvaluationRegistrationSuccessParams struct {
	// Stripe customer ID to attach to an entity-less registration evaluation.
	Customer *string `form:"customer"`
}

// Client details context.
type RadarCustomerEvaluationCreateEvaluationContextClientDetailsParams struct {
	// ID for the Radar Session associated with the customer evaluation.
	RadarSession *string `form:"radar_session"`
}

// Customer data
type RadarCustomerEvaluationCreateEvaluationContextCustomerDetailsCustomerDataParams struct {
	// Customer email
	Email *string `form:"email"`
	// Customer name
	Name *string `form:"name"`
	// Customer phone
	Phone *string `form:"phone"`
}

// Customer details context.
type RadarCustomerEvaluationCreateEvaluationContextCustomerDetailsParams struct {
	// Stripe customer ID
	Customer *string `form:"customer"`
	// Customer data
	CustomerData *RadarCustomerEvaluationCreateEvaluationContextCustomerDetailsCustomerDataParams `form:"customer_data"`
}

// Array of context entries for the evaluation.
type RadarCustomerEvaluationCreateEvaluationContextParams struct {
	// Client details context.
	ClientDetails *RadarCustomerEvaluationCreateEvaluationContextClientDetailsParams `form:"client_details"`
	// Customer details context.
	CustomerDetails *RadarCustomerEvaluationCreateEvaluationContextCustomerDetailsParams `form:"customer_details"`
	// The type of context entry.
	Type *string `form:"type"`
}

// Creates a new CustomerEvaluation object.
type RadarCustomerEvaluationCreateParams struct {
	Params `form:"*"`
	// Array of context entries for the evaluation.
	EvaluationContext []*RadarCustomerEvaluationCreateEvaluationContextParams `form:"evaluation_context"`
	// The type of evaluation event.
	EventType *string `form:"event_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *RadarCustomerEvaluationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Event payload for login_failed.
type RadarCustomerEvaluationUpdateLoginFailedParams struct {
	// The reason why this login failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_failed.
type RadarCustomerEvaluationUpdateRegistrationFailedParams struct {
	// The reason why this registration failed.
	Reason *string `form:"reason"`
}

// Event payload for registration_success.
type RadarCustomerEvaluationUpdateRegistrationSuccessParams struct {
	// Stripe customer ID to attach to an entity-less registration evaluation.
	Customer *string `form:"customer"`
}

// Reports an event on a CustomerEvaluation object.
type RadarCustomerEvaluationUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Event payload for login_failed.
	LoginFailed *RadarCustomerEvaluationUpdateLoginFailedParams `form:"login_failed"`
	// Event payload for registration_failed.
	RegistrationFailed *RadarCustomerEvaluationUpdateRegistrationFailedParams `form:"registration_failed"`
	// Event payload for registration_success.
	RegistrationSuccess *RadarCustomerEvaluationUpdateRegistrationSuccessParams `form:"registration_success"`
	// The type of event to report.
	Type *string `form:"type"`
}

// AddExpand appends a new field to expand.
func (p *RadarCustomerEvaluationUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Data about a failed login event.
type RadarCustomerEvaluationEventLoginFailed struct {
	// The reason why this login failed.
	Reason string `json:"reason"`
}

// Data about a failed registration event.
type RadarCustomerEvaluationEventRegistrationFailed struct {
	// The reason why this registration failed.
	Reason string `json:"reason"`
}

// A list of events that have been reported on this customer evaluation.
type RadarCustomerEvaluationEvent struct {
	// Data about a failed login event.
	LoginFailed *RadarCustomerEvaluationEventLoginFailed `json:"login_failed"`
	// Time at which the event occurred. Measured in seconds since the Unix epoch.
	OccurredAt int64 `json:"occurred_at"`
	// Data about a failed registration event.
	RegistrationFailed *RadarCustomerEvaluationEventRegistrationFailed `json:"registration_failed"`
	// The type of event that occurred.
	Type string `json:"type"`
}
type RadarCustomerEvaluationSignalsAccountSharing struct {
	// Time at which the signal was evaluated. Measured in seconds since the Unix epoch.
	EvaluatedAt int64 `json:"evaluated_at"`
	// The risk level for this signal.
	RiskLevel string `json:"risk_level"`
	// Score for this signal (float between 0.0-100.0).
	Score float64 `json:"score"`
}
type RadarCustomerEvaluationSignalsMultiAccounting struct {
	// Time at which the signal was evaluated. Measured in seconds since the Unix epoch.
	EvaluatedAt int64 `json:"evaluated_at"`
	// The risk level for this signal.
	RiskLevel string `json:"risk_level"`
	// Score for this signal (float between 0.0-100.0).
	Score float64 `json:"score"`
}

// A hash of signal objects providing Radar's evaluation for the lifecycle event.
type RadarCustomerEvaluationSignals struct {
	AccountSharing  *RadarCustomerEvaluationSignalsAccountSharing  `json:"account_sharing"`
	MultiAccounting *RadarCustomerEvaluationSignalsMultiAccounting `json:"multi_accounting"`
}

// Customer Evaluation resource returned by the Radar Customer Evaluations API.
type RadarCustomerEvaluation struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// The ID of the Stripe customer the customer evaluation is associated with.
	Customer string `json:"customer"`
	// A list of events that have been reported on this customer evaluation.
	Events []*RadarCustomerEvaluationEvent `json:"events"`
	// The type of evaluation event.
	EventType string `json:"event_type"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A hash of signal objects providing Radar's evaluation for the lifecycle event.
	Signals *RadarCustomerEvaluationSignals `json:"signals"`
}
