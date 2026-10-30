//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of RewardPrograms.
type V2MoneyManagementRewardProgramListParams struct {
	Params `form:"*"`
	// Maximum number of results to return. Default: 10. Valid range: 1-100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by RewardProgram status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Creates a new RewardProgram. RewardPrograms are inactive by default on creation; use
// ActivateRewardProgram to make them eligible for reward evaluation.
type V2MoneyManagementRewardProgramParams struct {
	Params `form:"*"`
	// Determines who performs reward money movement after rule evaluation. Only `application` is currently
	// supported.
	FulfilledBy *string `form:"fulfilled_by" json:"fulfilled_by,omitempty"`
	// Display name for the RewardProgram.
	Name *string `form:"name" json:"name,omitempty"`
}

// Activates a RewardProgram, making it eligible for reward evaluation. Idempotent: calling this on a
// RewardProgram that is already active returns 200 with the current object unchanged.
type V2MoneyManagementRewardProgramActivateParams struct {
	Params `form:"*"`
}

// Deactivates a RewardProgram. Idempotent: calling this on a RewardProgram that is already inactive
// returns 200 with the current object unchanged.
type V2MoneyManagementRewardProgramDeactivateParams struct {
	Params `form:"*"`
}

// Creates a new RewardProgram. RewardPrograms are inactive by default on creation; use
// ActivateRewardProgram to make them eligible for reward evaluation.
type V2MoneyManagementRewardProgramCreateParams struct {
	Params `form:"*"`
	// Determines who performs reward money movement after rule evaluation. Only `application` is currently
	// supported.
	FulfilledBy *string `form:"fulfilled_by" json:"fulfilled_by"`
	// Display name for the RewardProgram.
	Name *string `form:"name" json:"name"`
}

// Retrieves the details of an existing RewardProgram.
type V2MoneyManagementRewardProgramRetrieveParams struct {
	Params `form:"*"`
}

// Updates the mutable fields of an existing RewardProgram. Use ActivateRewardProgram /
// DeactivateRewardProgram to change the RewardProgram's status.
type V2MoneyManagementRewardProgramUpdateParams struct {
	Params `form:"*"`
	// Display name for the RewardProgram.
	Name *string `form:"name" json:"name,omitempty"`
}
