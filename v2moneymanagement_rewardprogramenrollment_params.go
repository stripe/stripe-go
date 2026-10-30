//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of RewardProgramEnrollments. Omitting `reward_program` returns platform-wide
// enrollments across all RewardPrograms.
type V2MoneyManagementRewardProgramEnrollmentListParams struct {
	Params `form:"*"`
	// Maximum number of results to return. Default: 10. Valid range: 1-100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by RewardProgram. Omitting this field returns platform-wide enrollments across all
	// RewardPrograms.
	RewardProgram *string `form:"reward_program" json:"reward_program,omitempty"`
	// Filter by RewardProgramEnrollment status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Issuing-specific per-enrollment configuration.
type V2MoneyManagementRewardProgramEnrollmentCardSpendParams struct {
	// The Issuer Program for this connected account. Immutable after creation.
	IssuerProgram *string `form:"issuer_program" json:"issuer_program"`
}

// Creates a new RewardProgramEnrollment, binding a connected account / Issuer Program to a
// RewardProgram. RewardProgramEnrollments are active by default on creation.
type V2MoneyManagementRewardProgramEnrollmentParams struct {
	Params `form:"*"`
	// The connected account to enroll into this RewardProgram.
	Account *string `form:"account" json:"account,omitempty"`
	// Issuing-specific per-enrollment configuration.
	CardSpend *V2MoneyManagementRewardProgramEnrollmentCardSpendParams `form:"card_spend" json:"card_spend,omitempty"`
	// The RewardProgram to enroll into.
	RewardProgram *string `form:"reward_program" json:"reward_program,omitempty"`
}

// Activates a RewardProgramEnrollment. At most one RewardProgramEnrollment may be active per Issuer
// Program globally, across all RewardPrograms. Idempotent: calling this on a RewardProgramEnrollment
// that is already active returns 200 with the current object unchanged.
type V2MoneyManagementRewardProgramEnrollmentActivateParams struct {
	Params `form:"*"`
}

// Deactivates a RewardProgramEnrollment. Idempotent: calling this on a RewardProgramEnrollment that is
// already inactive returns 200 with the current object unchanged.
type V2MoneyManagementRewardProgramEnrollmentDeactivateParams struct {
	Params `form:"*"`
}

// Issuing-specific per-enrollment configuration.
type V2MoneyManagementRewardProgramEnrollmentCreateCardSpendParams struct {
	// The Issuer Program for this connected account. Immutable after creation.
	IssuerProgram *string `form:"issuer_program" json:"issuer_program"`
}

// Creates a new RewardProgramEnrollment, binding a connected account / Issuer Program to a
// RewardProgram. RewardProgramEnrollments are active by default on creation.
type V2MoneyManagementRewardProgramEnrollmentCreateParams struct {
	Params `form:"*"`
	// The connected account to enroll into this RewardProgram.
	Account *string `form:"account" json:"account"`
	// Issuing-specific per-enrollment configuration.
	CardSpend *V2MoneyManagementRewardProgramEnrollmentCreateCardSpendParams `form:"card_spend" json:"card_spend"`
	// The RewardProgram to enroll into.
	RewardProgram *string `form:"reward_program" json:"reward_program"`
}

// Retrieves the details of an existing RewardProgramEnrollment.
type V2MoneyManagementRewardProgramEnrollmentRetrieveParams struct {
	Params `form:"*"`
}
