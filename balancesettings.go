//
//
// File generated from our OpenAPI spec
//
//

package stripe

// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
type BalanceSettingsPaymentsPayoutsScheduleInterval string

// List of values that BalanceSettingsPaymentsPayoutsScheduleInterval can take
const (
	BalanceSettingsPaymentsPayoutsScheduleIntervalDaily   BalanceSettingsPaymentsPayoutsScheduleInterval = "daily"
	BalanceSettingsPaymentsPayoutsScheduleIntervalManual  BalanceSettingsPaymentsPayoutsScheduleInterval = "manual"
	BalanceSettingsPaymentsPayoutsScheduleIntervalMonthly BalanceSettingsPaymentsPayoutsScheduleInterval = "monthly"
	BalanceSettingsPaymentsPayoutsScheduleIntervalWeekly  BalanceSettingsPaymentsPayoutsScheduleInterval = "weekly"
)

// The days of the week when available funds are paid out, specified as an array, for example, [`monday`, `tuesday`]. Only shown if `interval` is weekly.
type BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay string

// List of values that BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay can take
const (
	BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDayFriday    BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay = "friday"
	BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDayMonday    BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay = "monday"
	BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDayThursday  BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay = "thursday"
	BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDayTuesday   BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay = "tuesday"
	BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDayWednesday BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay = "wednesday"
)

// Whether the funds in this account can be paid out.
type BalanceSettingsPaymentsPayoutsStatus string

// List of values that BalanceSettingsPaymentsPayoutsStatus can take
const (
	BalanceSettingsPaymentsPayoutsStatusDisabled BalanceSettingsPaymentsPayoutsStatus = "disabled"
	BalanceSettingsPaymentsPayoutsStatusEnabled  BalanceSettingsPaymentsPayoutsStatus = "enabled"
)

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
type BalanceSettingsParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
	Payments *BalanceSettingsPaymentsParams `form:"payments"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
type BalanceSettingsPaymentsPayoutsScheduleParams struct {
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval"`
	// The days of the month when available funds are paid out, specified as an array of numbers between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyPayoutDays []*int64 `form:"monthly_payout_days"`
	// The days of the week when available funds are paid out, specified as an array, e.g., [`monday`, `tuesday`]. Required and applicable only if `interval` is `weekly`.
	WeeklyPayoutDays []*string `form:"weekly_payout_days"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPaymentsPayoutsParams struct {
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
	Schedule *BalanceSettingsPaymentsPayoutsScheduleParams `form:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Settings related to the account's balance settlement timing.
type BalanceSettingsPaymentsSettlementTimingParams struct {
	// The number of days charge funds are held before becoming available. The default value is `minimum`, representing the lowest available value for the account. The maximum value is 31. The `delay_days` parameter remains at the last configured value if `payouts.schedule.interval` is `manual`. [Learn more about controlling delay days](https://docs.stripe.com/connect/manage-payout-schedule).
	DelayDaysOverride *int64 `form:"delay_days_override"`
}

// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
type BalanceSettingsPaymentsParams struct {
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://docs.stripe.com/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances"`
	// Settings specific to the account's payouts.
	Payouts *BalanceSettingsPaymentsPayoutsParams `form:"payouts"`
	// Settings related to the account's balance settlement timing.
	SettlementTiming *BalanceSettingsPaymentsSettlementTimingParams `form:"settlement_timing"`
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
type BalanceSettingsRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
type BalanceSettingsUpdatePaymentsPayoutsScheduleParams struct {
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval"`
	// The days of the month when available funds are paid out, specified as an array of numbers between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyPayoutDays []*int64 `form:"monthly_payout_days"`
	// The days of the week when available funds are paid out, specified as an array, e.g., [`monday`, `tuesday`]. Required and applicable only if `interval` is `weekly`.
	WeeklyPayoutDays []*string `form:"weekly_payout_days"`
}

// Settings specific to the account's payouts.
type BalanceSettingsUpdatePaymentsPayoutsParams struct {
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
	Schedule *BalanceSettingsUpdatePaymentsPayoutsScheduleParams `form:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Settings related to the account's balance settlement timing.
type BalanceSettingsUpdatePaymentsSettlementTimingParams struct {
	// The number of days charge funds are held before becoming available. The default value is `minimum`, representing the lowest available value for the account. The maximum value is 31. The `delay_days` parameter remains at the last configured value if `payouts.schedule.interval` is `manual`. [Learn more about controlling delay days](https://docs.stripe.com/connect/manage-payout-schedule).
	DelayDaysOverride *int64 `form:"delay_days_override"`
}

// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
type BalanceSettingsUpdatePaymentsParams struct {
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://docs.stripe.com/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances"`
	// Settings specific to the account's payouts.
	Payouts *BalanceSettingsUpdatePaymentsPayoutsParams `form:"payouts"`
	// Settings related to the account's balance settlement timing.
	SettlementTiming *BalanceSettingsUpdatePaymentsSettlementTimingParams `form:"settlement_timing"`
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
type BalanceSettingsUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
	Payments *BalanceSettingsUpdatePaymentsParams `form:"payments"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation for details.
type BalanceSettingsPaymentsPayoutsSchedule struct {
	// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
	Interval BalanceSettingsPaymentsPayoutsScheduleInterval `json:"interval"`
	// The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.
	MonthlyPayoutDays []int64 `json:"monthly_payout_days"`
	// The days of the week when available funds are paid out, specified as an array, for example, [`monday`, `tuesday`]. Only shown if `interval` is weekly.
	WeeklyPayoutDays []BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay `json:"weekly_payout_days"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPaymentsPayouts struct {
	// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://stripe.com/docs/connect/bank-transfers#payout-information) documentation for details.
	Schedule *BalanceSettingsPaymentsPayoutsSchedule `json:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor string `json:"statement_descriptor"`
	// Whether the funds in this account can be paid out.
	Status BalanceSettingsPaymentsPayoutsStatus `json:"status"`
}
type BalanceSettingsPaymentsSettlementTiming struct {
	// The number of days charge funds are held before becoming available.
	DelayDays int64 `json:"delay_days"`
}
type BalanceSettingsPayments struct {
	// A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See [Understanding Connect account balances](https://docs.stripe.com/connect/account-balances) for details. The default value is `false` when [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts, otherwise `true`.
	DebitNegativeBalances bool `json:"debit_negative_balances"`
	// Settings specific to the account's payouts.
	Payouts          *BalanceSettingsPaymentsPayouts          `json:"payouts"`
	SettlementTiming *BalanceSettingsPaymentsSettlementTiming `json:"settlement_timing"`
}

// Options for customizing account balances and payout settings for a Stripe platform's connected accounts.
//
// This API is only available for users enrolled in the public preview for Accounts v2 on Stripe Connect.
// If you are not in this preview, please use the [Accounts v1 API](https://docs.stripe.com/api/accounts?api-version=2025-03-31.basil)
// to manage your connected accounts' balance settings instead.
type BalanceSettings struct {
	APIResource
	// String representing the object's type. Objects of the same type share the same value.
	Object   string                   `json:"object"`
	Payments *BalanceSettingsPayments `json:"payments"`
}
