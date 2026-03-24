//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List of signals that were triggered for evaluation.
type V2CoreAccountEvaluationEvaluationsTriggered string

// List of values that V2CoreAccountEvaluationEvaluationsTriggered can take
const (
	V2CoreAccountEvaluationEvaluationsTriggeredFraudulentWebsite V2CoreAccountEvaluationEvaluationsTriggered = "fraudulent_website"
)

// Account profile data.
type V2CoreAccountEvaluationAccountDataDefaultsProfile struct {
	// The business URL.
	BusinessURL string `json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs string `json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription string `json:"product_description,omitempty"`
}

// Default account settings.
type V2CoreAccountEvaluationAccountDataDefaults struct {
	// Account profile data.
	Profile *V2CoreAccountEvaluationAccountDataDefaultsProfile `json:"profile"`
}

// Business details for identity data.
type V2CoreAccountEvaluationAccountDataIdentityBusinessDetails struct {
	// Registered business name.
	RegisteredName string `json:"registered_name,omitempty"`
}

// Identity data.
type V2CoreAccountEvaluationAccountDataIdentity struct {
	// Business details for identity data.
	BusinessDetails *V2CoreAccountEvaluationAccountDataIdentityBusinessDetails `json:"business_details"`
}

// Account data if this evaluation is for an account without an existing Stripe entity.
type V2CoreAccountEvaluationAccountData struct {
	// Default account settings.
	Defaults *V2CoreAccountEvaluationAccountDataDefaults `json:"defaults,omitempty"`
	// Identity data.
	Identity *V2CoreAccountEvaluationAccountDataIdentity `json:"identity,omitempty"`
}

// Account Evaluation resource.
type V2CoreAccountEvaluation struct {
	APIResource
	// The account ID if this evaluation is for an existing account.
	Account string `json:"account,omitempty"`
	// Account data if this evaluation is for an account without an existing Stripe entity.
	AccountData *V2CoreAccountEvaluationAccountData `json:"account_data,omitempty"`
	// Timestamp at which the evaluation was created.
	Created time.Time `json:"created"`
	// List of signals that were triggered for evaluation.
	EvaluationsTriggered []V2CoreAccountEvaluationEvaluationsTriggered `json:"evaluations_triggered"`
	// Unique identifier for the account evaluation.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
