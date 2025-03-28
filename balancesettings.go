//
//
// File generated from our OpenAPI spec
//
//

package stripe

// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
type BalanceSettingsPayoutsScheduleInterval string

// List of values that BalanceSettingsPayoutsScheduleInterval can take
const (
	BalanceSettingsPayoutsScheduleIntervalDaily   BalanceSettingsPayoutsScheduleInterval = "daily"
	BalanceSettingsPayoutsScheduleIntervalManual  BalanceSettingsPayoutsScheduleInterval = "manual"
	BalanceSettingsPayoutsScheduleIntervalMonthly BalanceSettingsPayoutsScheduleInterval = "monthly"
	BalanceSettingsPayoutsScheduleIntervalWeekly  BalanceSettingsPayoutsScheduleInterval = "weekly"
)

// The day of the week funds will be paid out, of the style 'monday', 'tuesday', etc. Only shown if `interval` is weekly.
type BalanceSettingsPayoutsScheduleWeeklyAnchor string

// List of values that BalanceSettingsPayoutsScheduleWeeklyAnchor can take
const (
	BalanceSettingsPayoutsScheduleWeeklyAnchorFriday    BalanceSettingsPayoutsScheduleWeeklyAnchor = "friday"
	BalanceSettingsPayoutsScheduleWeeklyAnchorMonday    BalanceSettingsPayoutsScheduleWeeklyAnchor = "monday"
	BalanceSettingsPayoutsScheduleWeeklyAnchorThursday  BalanceSettingsPayoutsScheduleWeeklyAnchor = "thursday"
	BalanceSettingsPayoutsScheduleWeeklyAnchorTuesday   BalanceSettingsPayoutsScheduleWeeklyAnchor = "tuesday"
	BalanceSettingsPayoutsScheduleWeeklyAnchorWednesday BalanceSettingsPayoutsScheduleWeeklyAnchor = "wednesday"
)

// Whether the funds in this account can be paid out.
type BalanceSettingsPayoutsStatus string

// List of values that BalanceSettingsPayoutsStatus can take
const (
	BalanceSettingsPayoutsStatusDisabled BalanceSettingsPayoutsStatus = "disabled"
	BalanceSettingsPayoutsStatusEnabled  BalanceSettingsPayoutsStatus = "enabled"
)

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
type BalanceSettingsParams struct {
	Params `form:"*"`
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://stripe.com/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Settings specific to the account's payouts.
	Payouts *BalanceSettingsPayoutsParams `form:"payouts"`
	// Settings related to the account's balance settlement timing.
	SettlementTiming *BalanceSettingsSettlementTimingParams `form:"settlement_timing"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/connect/bank-transfers#payout-information) documentation.
type BalanceSettingsPayoutsScheduleParams struct {
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval"`
	// The day of the month when available funds are paid out, specified as a number between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyAnchor *int64 `form:"monthly_anchor"`
	// The day of the week when available funds are paid out (required and applicable only if `interval` is `weekly`.)
	WeeklyAnchor *string `form:"weekly_anchor"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPayoutsParams struct {
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://stripe.com/connect/bank-transfers#payout-information) documentation.
	Schedule *BalanceSettingsPayoutsScheduleParams `form:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Settings related to the account's balance settlement timing.
type BalanceSettingsSettlementTimingParams struct {
	// The number of days charge funds are held before becoming available. May also be set to `minimum`, representing the lowest available value for the account country. Default is `minimum`. The `delay_days` parameter remains at the last configured value if `payouts.schedule.interval` is `manual`. [Learn more about controlling payout delay days](https://stripe.com/connect/manage-payout-schedule).
	DelayDays *int64 `form:"delay_days"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation for details.
type BalanceSettingsPayoutsSchedule struct {
	// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
	Interval BalanceSettingsPayoutsScheduleInterval `json:"interval"`
	// The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.
	MonthlyAnchor int64 `json:"monthly_anchor"`
	// The day of the week funds will be paid out, of the style 'monday', 'tuesday', etc. Only shown if `interval` is weekly.
	WeeklyAnchor BalanceSettingsPayoutsScheduleWeeklyAnchor `json:"weekly_anchor"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPayouts struct {
	// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation for details.
	Schedule *BalanceSettingsPayoutsSchedule `json:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor string `json:"statement_descriptor"`
	// Whether the funds in this account can be paid out.
	Status BalanceSettingsPayoutsStatus `json:"status"`
}
type BalanceSettingsSettlementTiming struct {
	// The number of days charge funds are held before becoming available.
	DelayDays int64 `json:"delay_days"`
}

// "Options for customizing account balances within Stripe."
type BalanceSettings struct {
	APIResource
	// A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See [Understanding Connect account balances](https://stripe.com/connect/account-balances) for details. The default value is `false` when [controller.requirement_collection](https://stripe.com/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts, otherwise `true`.
	DebitNegativeBalances bool `json:"debit_negative_balances"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Settings specific to the account's payouts.
	Payouts          *BalanceSettingsPayouts          `json:"payouts"`
	SettlementTiming *BalanceSettingsSettlementTiming `json:"settlement_timing"`
}
