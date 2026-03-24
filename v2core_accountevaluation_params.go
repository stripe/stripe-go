//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Account profile data.
type V2CoreAccountEvaluationAccountDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2CoreAccountEvaluationAccountDataDefaultsParams struct {
	// Account profile data.
	Profile *V2CoreAccountEvaluationAccountDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Business details for identity data.
type V2CoreAccountEvaluationAccountDataIdentityBusinessDetailsParams struct {
	// Registered business name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Identity data.
type V2CoreAccountEvaluationAccountDataIdentityParams struct {
	// Business details for identity data.
	BusinessDetails *V2CoreAccountEvaluationAccountDataIdentityBusinessDetailsParams `form:"business_details" json:"business_details"`
}

// Account data for entity-less evaluation. Exactly one of account or account_data must be provided.
type V2CoreAccountEvaluationAccountDataParams struct {
	// Default account settings.
	Defaults *V2CoreAccountEvaluationAccountDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// Identity data.
	Identity *V2CoreAccountEvaluationAccountDataIdentityParams `form:"identity" json:"identity,omitempty"`
}

// Creates a new account evaluation to trigger signal evaluations on an account or account data.
type V2CoreAccountEvaluationParams struct {
	Params `form:"*"`
	// The account ID to evaluate. Exactly one of account or account_data must be provided.
	Account *string `form:"account" json:"account,omitempty"`
	// Account data for entity-less evaluation. Exactly one of account or account_data must be provided.
	AccountData *V2CoreAccountEvaluationAccountDataParams `form:"account_data" json:"account_data,omitempty"`
	// List of signals to evaluate.
	Signals []*string `form:"signals" json:"signals"`
}

// Account profile data.
type V2CoreAccountEvaluationCreateAccountDataDefaultsProfileParams struct {
	// The business URL.
	BusinessURL *string `form:"business_url" json:"business_url"`
	// Doing business as (DBA) name.
	DoingBusinessAs *string `form:"doing_business_as" json:"doing_business_as,omitempty"`
	// Description of the account's product or service.
	ProductDescription *string `form:"product_description" json:"product_description,omitempty"`
}

// Default account settings.
type V2CoreAccountEvaluationCreateAccountDataDefaultsParams struct {
	// Account profile data.
	Profile *V2CoreAccountEvaluationCreateAccountDataDefaultsProfileParams `form:"profile" json:"profile"`
}

// Business details for identity data.
type V2CoreAccountEvaluationCreateAccountDataIdentityBusinessDetailsParams struct {
	// Registered business name.
	RegisteredName *string `form:"registered_name" json:"registered_name,omitempty"`
}

// Identity data.
type V2CoreAccountEvaluationCreateAccountDataIdentityParams struct {
	// Business details for identity data.
	BusinessDetails *V2CoreAccountEvaluationCreateAccountDataIdentityBusinessDetailsParams `form:"business_details" json:"business_details"`
}

// Account data for entity-less evaluation. Exactly one of account or account_data must be provided.
type V2CoreAccountEvaluationCreateAccountDataParams struct {
	// Default account settings.
	Defaults *V2CoreAccountEvaluationCreateAccountDataDefaultsParams `form:"defaults" json:"defaults,omitempty"`
	// Identity data.
	Identity *V2CoreAccountEvaluationCreateAccountDataIdentityParams `form:"identity" json:"identity,omitempty"`
}

// Creates a new account evaluation to trigger signal evaluations on an account or account data.
type V2CoreAccountEvaluationCreateParams struct {
	Params `form:"*"`
	// The account ID to evaluate. Exactly one of account or account_data must be provided.
	Account *string `form:"account" json:"account,omitempty"`
	// Account data for entity-less evaluation. Exactly one of account or account_data must be provided.
	AccountData *V2CoreAccountEvaluationCreateAccountDataParams `form:"account_data" json:"account_data,omitempty"`
	// List of signals to evaluate.
	Signals []*string `form:"signals" json:"signals"`
}
