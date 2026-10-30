//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "github.com/shopspring/decimal"

// Returns a list of RewardProgramRules for a RewardProgram.
type V2MoneyManagementRewardProgramsRuleListParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram to list rules for.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
	// Maximum number of results to return. Default: 10. Valid range: 1-100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by RewardProgramRule status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
type V2MoneyManagementRewardProgramsRuleConditionCardSpendMerchantCategoryCodesParams struct {
	// Four-digit Merchant Category Codes to explicitly exclude. Takes precedence over `included`.
	Excluded []*string `form:"excluded" json:"excluded"`
	// Four-digit Merchant Category Codes to include. An empty list matches all MCCs.
	Included []*string `form:"included" json:"included"`
}

// Card-spend specific matching condition.
type V2MoneyManagementRewardProgramsRuleConditionCardSpendParams struct {
	// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
	MerchantCategoryCodes *V2MoneyManagementRewardProgramsRuleConditionCardSpendMerchantCategoryCodesParams `form:"merchant_category_codes" json:"merchant_category_codes,omitempty"`
}

// Conditions that must be true for this rule to match a card transaction. Omitting this field, or
// providing an empty condition, means the rule matches all transactions.
type V2MoneyManagementRewardProgramsRuleConditionParams struct {
	// Card-spend specific matching condition.
	CardSpend *V2MoneyManagementRewardProgramsRuleConditionCardSpendParams `form:"card_spend" json:"card_spend,omitempty"`
}

// Cashback reward configuration.
type V2MoneyManagementRewardProgramsRuleRewardCashbackParams struct {
	// Cashback rate as a percentage, represented as a decimal string. For example, `"2"` represents 2% and
	// `"1.5"` represents 1.5%. Must be greater than `"0"` and less than or equal to `"100"`.
	Rate *decimal.Decimal `form:"rate" json:"rate"`
}

// The reward to grant when this rule matches.
type V2MoneyManagementRewardProgramsRuleRewardParams struct {
	// Cashback reward configuration.
	Cashback *V2MoneyManagementRewardProgramsRuleRewardCashbackParams `form:"cashback" json:"cashback"`
}

// Creates a new RewardProgramRule under the given RewardProgram. RewardProgramRules are inactive by
// default on creation; use ActivateRewardProgramRule to make them eligible for reward evaluation.
type V2MoneyManagementRewardProgramsRuleParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram to create this rule under.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
	// Conditions that must be true for this rule to match a card transaction. Omitting this field, or
	// providing an empty condition, means the rule matches all transactions.
	Condition *V2MoneyManagementRewardProgramsRuleConditionParams `form:"condition" json:"condition,omitempty"`
	// Human-readable name for this rule. Max 250 characters.
	Name *string `form:"name" json:"name,omitempty"`
	// The reward to grant when this rule matches.
	Reward *V2MoneyManagementRewardProgramsRuleRewardParams `form:"reward" json:"reward,omitempty"`
}

// Activates a RewardProgramRule. At most one RewardProgramRule may be active per RewardProgram at a
// time. Idempotent: calling this on a RewardProgramRule that is already active returns 200 with the
// current object unchanged.
type V2MoneyManagementRewardProgramsRuleActivateParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram the rule belongs to.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
}

// Deactivates a RewardProgramRule. Idempotent: calling this on a RewardProgramRule that is already
// inactive returns 200 with the current object unchanged.
type V2MoneyManagementRewardProgramsRuleDeactivateParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram the rule belongs to.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
}

// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
type V2MoneyManagementRewardProgramsRuleCreateConditionCardSpendMerchantCategoryCodesParams struct {
	// Four-digit Merchant Category Codes to explicitly exclude. Takes precedence over `included`.
	Excluded []*string `form:"excluded" json:"excluded"`
	// Four-digit Merchant Category Codes to include. An empty list matches all MCCs.
	Included []*string `form:"included" json:"included"`
}

// Card-spend specific matching condition.
type V2MoneyManagementRewardProgramsRuleCreateConditionCardSpendParams struct {
	// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
	MerchantCategoryCodes *V2MoneyManagementRewardProgramsRuleCreateConditionCardSpendMerchantCategoryCodesParams `form:"merchant_category_codes" json:"merchant_category_codes,omitempty"`
}

// Conditions that must be true for this rule to match a card transaction. Omitting this field, or
// providing an empty condition, means the rule matches all transactions.
type V2MoneyManagementRewardProgramsRuleCreateConditionParams struct {
	// Card-spend specific matching condition.
	CardSpend *V2MoneyManagementRewardProgramsRuleCreateConditionCardSpendParams `form:"card_spend" json:"card_spend,omitempty"`
}

// Cashback reward configuration.
type V2MoneyManagementRewardProgramsRuleCreateRewardCashbackParams struct {
	// Cashback rate as a percentage, represented as a decimal string. For example, `"2"` represents 2% and
	// `"1.5"` represents 1.5%. Must be greater than `"0"` and less than or equal to `"100"`.
	Rate *decimal.Decimal `form:"rate" json:"rate"`
}

// The reward to grant when this rule matches.
type V2MoneyManagementRewardProgramsRuleCreateRewardParams struct {
	// Cashback reward configuration.
	Cashback *V2MoneyManagementRewardProgramsRuleCreateRewardCashbackParams `form:"cashback" json:"cashback"`
}

// Creates a new RewardProgramRule under the given RewardProgram. RewardProgramRules are inactive by
// default on creation; use ActivateRewardProgramRule to make them eligible for reward evaluation.
type V2MoneyManagementRewardProgramsRuleCreateParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram to create this rule under.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
	// Conditions that must be true for this rule to match a card transaction. Omitting this field, or
	// providing an empty condition, means the rule matches all transactions.
	Condition *V2MoneyManagementRewardProgramsRuleCreateConditionParams `form:"condition" json:"condition,omitempty"`
	// Human-readable name for this rule. Max 250 characters.
	Name *string `form:"name" json:"name"`
	// The reward to grant when this rule matches.
	Reward *V2MoneyManagementRewardProgramsRuleCreateRewardParams `form:"reward" json:"reward"`
}

// Retrieves the details of an existing RewardProgramRule.
type V2MoneyManagementRewardProgramsRuleRetrieveParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram the rule belongs to.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
}

// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
type V2MoneyManagementRewardProgramsRuleUpdateConditionCardSpendMerchantCategoryCodesParams struct {
	// Four-digit Merchant Category Codes to explicitly exclude. Takes precedence over `included`.
	Excluded []*string `form:"excluded" json:"excluded"`
	// Four-digit Merchant Category Codes to include. An empty list matches all MCCs.
	Included []*string `form:"included" json:"included"`
}

// Card-spend specific matching condition.
type V2MoneyManagementRewardProgramsRuleUpdateConditionCardSpendParams struct {
	// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
	MerchantCategoryCodes *V2MoneyManagementRewardProgramsRuleUpdateConditionCardSpendMerchantCategoryCodesParams `form:"merchant_category_codes" json:"merchant_category_codes,omitempty"`
}

// Conditions that must be true for this rule to match a card transaction.
type V2MoneyManagementRewardProgramsRuleUpdateConditionParams struct {
	// Card-spend specific matching condition.
	CardSpend *V2MoneyManagementRewardProgramsRuleUpdateConditionCardSpendParams `form:"card_spend" json:"card_spend,omitempty"`
}

// Cashback reward configuration.
type V2MoneyManagementRewardProgramsRuleUpdateRewardCashbackParams struct {
	// Cashback rate as a percentage, represented as a decimal string. For example, `"2"` represents 2% and
	// `"1.5"` represents 1.5%. Must be greater than `"0"` and less than or equal to `"100"`.
	Rate *decimal.Decimal `form:"rate" json:"rate"`
}

// The reward to grant when this rule matches.
type V2MoneyManagementRewardProgramsRuleUpdateRewardParams struct {
	// Cashback reward configuration.
	Cashback *V2MoneyManagementRewardProgramsRuleUpdateRewardCashbackParams `form:"cashback" json:"cashback"`
}

// Updates the mutable fields of an existing RewardProgramRule. Use ActivateRewardProgramRule /
// DeactivateRewardProgramRule to change the RewardProgramRule's status.
type V2MoneyManagementRewardProgramsRuleUpdateParams struct {
	Params `form:"*"`
	// The ID of the RewardProgram the rule belongs to.
	RewardProgramID *string `form:"-" json:"-"` // Included in URL
	// Conditions that must be true for this rule to match a card transaction.
	Condition *V2MoneyManagementRewardProgramsRuleUpdateConditionParams `form:"condition" json:"condition,omitempty"`
	// Human-readable name for this rule. Max 250 characters.
	Name *string `form:"name" json:"name,omitempty"`
	// The reward to grant when this rule matches.
	Reward *V2MoneyManagementRewardProgramsRuleUpdateRewardParams `form:"reward" json:"reward,omitempty"`
}
