//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

// The status of the RewardProgramRule. Inactive by default on creation; use activate/deactivate to
// change.
type V2MoneyManagementRewardProgramRuleStatus string

// List of values that V2MoneyManagementRewardProgramRuleStatus can take
const (
	V2MoneyManagementRewardProgramRuleStatusActive   V2MoneyManagementRewardProgramRuleStatus = "active"
	V2MoneyManagementRewardProgramRuleStatusInactive V2MoneyManagementRewardProgramRuleStatus = "inactive"
)

// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
type V2MoneyManagementRewardProgramRuleConditionCardSpendMerchantCategoryCodes struct {
	// Four-digit Merchant Category Codes to explicitly exclude. Takes precedence over `included`.
	Excluded []string `json:"excluded"`
	// Four-digit Merchant Category Codes to include. An empty list matches all MCCs.
	Included []string `json:"included"`
}

// Card-spend specific matching condition.
type V2MoneyManagementRewardProgramRuleConditionCardSpend struct {
	// Merchant Category Code include/exclude configuration. Omitted means the rule matches all MCCs.
	MerchantCategoryCodes *V2MoneyManagementRewardProgramRuleConditionCardSpendMerchantCategoryCodes `json:"merchant_category_codes,omitempty"`
}

// Conditions that must be true for this rule to match a card transaction. Empty means the rule matches
// all transactions.
type V2MoneyManagementRewardProgramRuleCondition struct {
	// Card-spend specific matching condition.
	CardSpend *V2MoneyManagementRewardProgramRuleConditionCardSpend `json:"card_spend,omitempty"`
}

// Cashback reward configuration.
type V2MoneyManagementRewardProgramRuleRewardCashback struct {
	// Cashback rate as a percentage, represented as a decimal string. For example, `"2"` represents 2% and
	// `"1.5"` represents 1.5%. Must be greater than `"0"` and less than or equal to `"100"`.
	Rate decimal.Decimal `json:"rate"`
}

// The reward to grant when this rule matches.
type V2MoneyManagementRewardProgramRuleReward struct {
	// Cashback reward configuration.
	Cashback *V2MoneyManagementRewardProgramRuleRewardCashback `json:"cashback"`
}

// A RewardProgramRule defines the matching conditions and cashback reward outcome shared across all
// active enrollments in a RewardProgram. At most one RewardProgramRule may be active per RewardProgram
// at a time.
type V2MoneyManagementRewardProgramRule struct {
	APIResource
	// Conditions that must be true for this rule to match a card transaction. Empty means the rule matches
	// all transactions.
	Condition *V2MoneyManagementRewardProgramRuleCondition `json:"condition"`
	// Time at which the RewardProgramRule was created.
	Created time.Time `json:"created"`
	// Unique identifier for the RewardProgramRule.
	ID string `json:"id"`
	// Whether this object exists in livemode or testmode.
	Livemode bool `json:"livemode"`
	// Human-readable name for this rule. Max 250 characters.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The reward to grant when this rule matches.
	Reward *V2MoneyManagementRewardProgramRuleReward `json:"reward"`
	// The ID of the RewardProgram this rule belongs to. Immutable after creation.
	RewardProgram string `json:"reward_program"`
	// The status of the RewardProgramRule. Inactive by default on creation; use activate/deactivate to
	// change.
	Status V2MoneyManagementRewardProgramRuleStatus `json:"status"`
}
