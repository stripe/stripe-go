//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Account selector for the historical outcomes requested.
type V2RiskEvaluationOutcomeListAccountDetailsParams struct {
	// The v2 account ID to list evaluation outcomes for.
	Account *string `form:"account" json:"account,omitempty"`
}

// Lists all matching historical evaluation outcomes, newest first.
type V2RiskEvaluationOutcomeListParams struct {
	Params `form:"*"`
	// Account selector for the historical outcomes requested.
	AccountDetails *V2RiskEvaluationOutcomeListAccountDetailsParams `form:"account_details" json:"account_details"`
	// Maximum number of results per page. Defaults to 10; use 1 for the latest matching outcome.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Evaluation origin to filter by. Omit to match all origins.
	Origin *string `form:"origin" json:"origin,omitempty"`
	// Evaluation type to filter by. Omit to match all types.
	Type *string `form:"type" json:"type,omitempty"`
}

// Retrieves an evaluation outcome by its ID.
type V2RiskEvaluationOutcomeParams struct {
	Params `form:"*"`
}

// Retrieves an evaluation outcome by its ID.
type V2RiskEvaluationOutcomeRetrieveParams struct {
	Params `form:"*"`
}
