//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of revenue share that caused the EarnedCredit.
type V2MoneyManagementEarnedCreditRevenueShareType string

// List of values that V2MoneyManagementEarnedCreditRevenueShareType can take
const (
	V2MoneyManagementEarnedCreditRevenueShareTypeAdministrativeFacilitationFee V2MoneyManagementEarnedCreditRevenueShareType = "administrative_facilitation_fee"
	V2MoneyManagementEarnedCreditRevenueShareTypeSavingsReferral               V2MoneyManagementEarnedCreditRevenueShareType = "savings_referral"
)

// The program from which the reward was earned.
type V2MoneyManagementEarnedCreditRewardEarnedFrom string

// List of values that V2MoneyManagementEarnedCreditRewardEarnedFrom can take
const (
	V2MoneyManagementEarnedCreditRewardEarnedFromPlatformCashRewards V2MoneyManagementEarnedCreditRewardEarnedFrom = "platform_cash_rewards"
)

// The status of the EarnedCredit.
type V2MoneyManagementEarnedCreditStatus string

// List of values that V2MoneyManagementEarnedCreditStatus can take
const (
	V2MoneyManagementEarnedCreditStatusSucceeded V2MoneyManagementEarnedCreditStatus = "succeeded"
)

// The type of flow that caused the EarnedCredit.
type V2MoneyManagementEarnedCreditType string

// List of values that V2MoneyManagementEarnedCreditType can take
const (
	V2MoneyManagementEarnedCreditTypeInterest     V2MoneyManagementEarnedCreditType = "interest"
	V2MoneyManagementEarnedCreditTypeRevenueShare V2MoneyManagementEarnedCreditType = "revenue_share"
	V2MoneyManagementEarnedCreditTypeReward       V2MoneyManagementEarnedCreditType = "reward"
)

// The period during which the credit was earned.
type V2MoneyManagementEarnedCreditPeriod struct {
	// The end date of the period during which the credit was earned, inclusive.
	EndDate string `json:"end_date"`
	// The start date of the period during which the credit was earned, inclusive.
	StartDate string `json:"start_date"`
}

// Details about the revenue share that caused the EarnedCredit.
type V2MoneyManagementEarnedCreditRevenueShare struct {
	// The type of revenue share that caused the EarnedCredit.
	Type V2MoneyManagementEarnedCreditRevenueShareType `json:"type"`
}

// Details about the reward that caused the EarnedCredit.
type V2MoneyManagementEarnedCreditReward struct {
	// The program from which the reward was earned.
	EarnedFrom V2MoneyManagementEarnedCreditRewardEarnedFrom `json:"earned_from"`
	// The Account that funded the reward.
	FromAccount string `json:"from_account"`
	// The OutboundPayment that delivered the reward.
	OutboundPayment string `json:"outbound_payment"`
}

// Timestamps for EarnedCredit status transitions.
type V2MoneyManagementEarnedCreditStatusTransitions struct {
	// The time at which the EarnedCredit succeeded.
	SucceededAt time.Time `json:"succeeded_at,omitempty"`
}

// The EarnedCredit object.
type V2MoneyManagementEarnedCredit struct {
	APIResource
	// The amount and currency of the EarnedCredit.
	Amount Amount `json:"amount"`
	// Time at which the EarnedCredit was created.
	Created time.Time `json:"created"`
	// Description of the EarnedCredit.
	Description string `json:"description"`
	// The FinancialAccount that earned the credit.
	FinancialAccount string `json:"financial_account"`
	// Unique identifier for the EarnedCredit.
	ID string `json:"id"`
	// Has the value true if the object exists in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The period during which the credit was earned.
	Period *V2MoneyManagementEarnedCreditPeriod `json:"period,omitempty"`
	// Details about the revenue share that caused the EarnedCredit.
	RevenueShare *V2MoneyManagementEarnedCreditRevenueShare `json:"revenue_share,omitempty"`
	// Details about the reward that caused the EarnedCredit.
	Reward *V2MoneyManagementEarnedCreditReward `json:"reward,omitempty"`
	// The status of the EarnedCredit.
	Status V2MoneyManagementEarnedCreditStatus `json:"status"`
	// Timestamps for EarnedCredit status transitions.
	StatusTransitions *V2MoneyManagementEarnedCreditStatusTransitions `json:"status_transitions"`
	// The type of flow that caused the EarnedCredit.
	Type V2MoneyManagementEarnedCreditType `json:"type"`
}
