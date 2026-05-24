//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of automatic transfer rule.
type BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyType string

// List of values that BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyType can take
const (
	BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyTypeTransferAll        BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyType = "transfer_all"
	BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyTypeTransferUpToAmount BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyType = "transfer_up_to_amount"
)

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
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
	Payments *BalanceSettingsPaymentsParams `form:"payments" json:"payments,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
type BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyParams struct {
	// The ID of the FinancialAccount that funds will be transferred to during automatic transfers.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
	// The maximum amount in minor units to transfer to the FinancialAccount. Required and only applicable when `type` is `transfer_up_to_amount`.
	TransferUpToAmount *int64 `form:"transfer_up_to_amount" json:"transfer_up_to_amount,omitempty"`
	// The type of automatic transfer rule.
	Type *string `form:"type" json:"type"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
type BalanceSettingsPaymentsPayoutsScheduleParams struct {
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The days of the month when available funds are paid out, specified as an array of numbers between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyPayoutDays []*int64 `form:"monthly_payout_days" json:"monthly_payout_days,omitempty"`
	// The days of the week when available funds are paid out, specified as an array, e.g., [`monday`, `tuesday`]. Required and applicable only if `interval` is `weekly`.
	WeeklyPayoutDays []*string `form:"weekly_payout_days" json:"weekly_payout_days,omitempty"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPaymentsPayoutsParams struct {
	// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
	AutomaticTransferRulesByCurrency map[string][]*BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyParams `form:"automatic_transfer_rules_by_currency" json:"automatic_transfer_rules_by_currency,omitempty"`
	// The minimum balance amount to retain per currency after automatic payouts. Only funds that exceed these amounts are paid out. Learn more about the [minimum balances for automatic payouts](https://docs.stripe.com/payouts/minimum-balances-for-automatic-payouts).
	MinimumBalanceByCurrency map[string]*int64 `form:"minimum_balance_by_currency" json:"minimum_balance_by_currency,omitempty"`
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
	Schedule *BalanceSettingsPaymentsPayoutsScheduleParams `form:"schedule" json:"schedule,omitempty"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string                                          `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	UnsetFields         []BalanceSettingsPaymentsPayoutsParamsUnsetField `form:"-" json:"-"`
}

// BalanceSettingsPaymentsPayoutsParamsUnsetField is the list of fields that can be cleared/unset on BalanceSettingsPaymentsPayoutsParams.
type BalanceSettingsPaymentsPayoutsParamsUnsetField string

const (
	BalanceSettingsPaymentsPayoutsParamsUnsetFieldAutomaticTransferRulesByCurrency BalanceSettingsPaymentsPayoutsParamsUnsetField = "automatic_transfer_rules_by_currency"
	BalanceSettingsPaymentsPayoutsParamsUnsetFieldMinimumBalanceByCurrency         BalanceSettingsPaymentsPayoutsParamsUnsetField = "minimum_balance_by_currency"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BalanceSettingsPaymentsPayoutsParams) AddUnsetField(field BalanceSettingsPaymentsPayoutsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
type BalanceSettingsPaymentsSettlementTimingStartOfDayParams struct {
	// Hour at which the customized start of day begins according to the given timezone. Must be a [supported customized start of day hour](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Hour *int64 `form:"hour" json:"hour,omitempty"`
	// Minutes at which the customized start of day begins according to the given timezone. Must be either 0 or 30.
	Minutes *int64 `form:"minutes" json:"minutes,omitempty"`
	// Timezone for the customized start of day. Must be a [supported customized start of day timezone](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// Settings related to the account's balance settlement timing.
type BalanceSettingsPaymentsSettlementTimingParams struct {
	// Change `delay_days` for this account, which determines the number of days charge funds are held before becoming available. The maximum value is 31. Passing an empty string to `delay_days_override` will return `delay_days` to the default, which is the lowest available value for the account. [Learn more about controlling delay days](https://docs.stripe.com/connect/manage-payout-schedule).
	DelayDaysOverride *int64 `form:"delay_days_override" json:"delay_days_override,omitempty"`
	// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
	StartOfDay  *BalanceSettingsPaymentsSettlementTimingStartOfDayParams  `form:"start_of_day" json:"start_of_day,omitempty"`
	UnsetFields []BalanceSettingsPaymentsSettlementTimingParamsUnsetField `form:"-" json:"-"`
}

// BalanceSettingsPaymentsSettlementTimingParamsUnsetField is the list of fields that can be cleared/unset on BalanceSettingsPaymentsSettlementTimingParams.
type BalanceSettingsPaymentsSettlementTimingParamsUnsetField string

const (
	BalanceSettingsPaymentsSettlementTimingParamsUnsetFieldDelayDaysOverride BalanceSettingsPaymentsSettlementTimingParamsUnsetField = "delay_days_override"
	BalanceSettingsPaymentsSettlementTimingParamsUnsetFieldStartOfDay        BalanceSettingsPaymentsSettlementTimingParamsUnsetField = "start_of_day"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BalanceSettingsPaymentsSettlementTimingParams) AddUnsetField(field BalanceSettingsPaymentsSettlementTimingParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
type BalanceSettingsPaymentsParams struct {
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://docs.stripe.com/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances" json:"debit_negative_balances,omitempty"`
	// Settings specific to the account's payouts.
	Payouts *BalanceSettingsPaymentsPayoutsParams `form:"payouts" json:"payouts,omitempty"`
	// Settings related to the account's balance settlement timing.
	SettlementTiming *BalanceSettingsPaymentsSettlementTimingParams `form:"settlement_timing" json:"settlement_timing,omitempty"`
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
type BalanceSettingsRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
type BalanceSettingsUpdatePaymentsPayoutsAutomaticTransferRulesByCurrencyParams struct {
	// The ID of the FinancialAccount that funds will be transferred to during automatic transfers.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
	// The maximum amount in minor units to transfer to the FinancialAccount. Required and only applicable when `type` is `transfer_up_to_amount`.
	TransferUpToAmount *int64 `form:"transfer_up_to_amount" json:"transfer_up_to_amount,omitempty"`
	// The type of automatic transfer rule.
	Type *string `form:"type" json:"type"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
type BalanceSettingsUpdatePaymentsPayoutsScheduleParams struct {
	// How frequently available funds are paid out. One of: `daily`, `manual`, `weekly`, or `monthly`. Default is `daily`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// The days of the month when available funds are paid out, specified as an array of numbers between 1--31. Payouts nominally scheduled between the 29th and 31st of the month are instead sent on the last day of a shorter month. Required and applicable only if `interval` is `monthly`.
	MonthlyPayoutDays []*int64 `form:"monthly_payout_days" json:"monthly_payout_days,omitempty"`
	// The days of the week when available funds are paid out, specified as an array, e.g., [`monday`, `tuesday`]. Required and applicable only if `interval` is `weekly`.
	WeeklyPayoutDays []*string `form:"weekly_payout_days" json:"weekly_payout_days,omitempty"`
}

// Settings specific to the account's payouts.
type BalanceSettingsUpdatePaymentsPayoutsParams struct {
	// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
	AutomaticTransferRulesByCurrency map[string][]*BalanceSettingsUpdatePaymentsPayoutsAutomaticTransferRulesByCurrencyParams `form:"automatic_transfer_rules_by_currency" json:"automatic_transfer_rules_by_currency,omitempty"`
	// The minimum balance amount to retain per currency after automatic payouts. Only funds that exceed these amounts are paid out. Learn more about the [minimum balances for automatic payouts](https://docs.stripe.com/payouts/minimum-balances-for-automatic-payouts).
	MinimumBalanceByCurrency map[string]*int64 `form:"minimum_balance_by_currency" json:"minimum_balance_by_currency,omitempty"`
	// Details on when funds from charges are available, and when they are paid out to an external account. For details, see our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation.
	Schedule *BalanceSettingsUpdatePaymentsPayoutsScheduleParams `form:"schedule" json:"schedule,omitempty"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor *string                                                `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
	UnsetFields         []BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField `form:"-" json:"-"`
}

// BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField is the list of fields that can be cleared/unset on BalanceSettingsUpdatePaymentsPayoutsParams.
type BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField string

const (
	BalanceSettingsUpdatePaymentsPayoutsParamsUnsetFieldAutomaticTransferRulesByCurrency BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField = "automatic_transfer_rules_by_currency"
	BalanceSettingsUpdatePaymentsPayoutsParamsUnsetFieldMinimumBalanceByCurrency         BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField = "minimum_balance_by_currency"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BalanceSettingsUpdatePaymentsPayoutsParams) AddUnsetField(field BalanceSettingsUpdatePaymentsPayoutsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
type BalanceSettingsUpdatePaymentsSettlementTimingStartOfDayParams struct {
	// Hour at which the customized start of day begins according to the given timezone. Must be a [supported customized start of day hour](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Hour *int64 `form:"hour" json:"hour,omitempty"`
	// Minutes at which the customized start of day begins according to the given timezone. Must be either 0 or 30.
	Minutes *int64 `form:"minutes" json:"minutes,omitempty"`
	// Timezone for the customized start of day. Must be a [supported customized start of day timezone](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Timezone *string `form:"timezone" json:"timezone,omitempty"`
}

// Settings related to the account's balance settlement timing.
type BalanceSettingsUpdatePaymentsSettlementTimingParams struct {
	// Change `delay_days` for this account, which determines the number of days charge funds are held before becoming available. The maximum value is 31. Passing an empty string to `delay_days_override` will return `delay_days` to the default, which is the lowest available value for the account. [Learn more about controlling delay days](https://docs.stripe.com/connect/manage-payout-schedule).
	DelayDaysOverride *int64 `form:"delay_days_override" json:"delay_days_override,omitempty"`
	// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
	StartOfDay  *BalanceSettingsUpdatePaymentsSettlementTimingStartOfDayParams  `form:"start_of_day" json:"start_of_day,omitempty"`
	UnsetFields []BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField `form:"-" json:"-"`
}

// BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField is the list of fields that can be cleared/unset on BalanceSettingsUpdatePaymentsSettlementTimingParams.
type BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField string

const (
	BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetFieldDelayDaysOverride BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField = "delay_days_override"
	BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetFieldStartOfDay        BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField = "start_of_day"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BalanceSettingsUpdatePaymentsSettlementTimingParams) AddUnsetField(field BalanceSettingsUpdatePaymentsSettlementTimingParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
type BalanceSettingsUpdatePaymentsParams struct {
	// A Boolean indicating whether Stripe should try to reclaim negative balances from an attached bank account. For details, see [Understanding Connect Account Balances](https://docs.stripe.com/connect/account-balances).
	DebitNegativeBalances *bool `form:"debit_negative_balances" json:"debit_negative_balances,omitempty"`
	// Settings specific to the account's payouts.
	Payouts *BalanceSettingsUpdatePaymentsPayoutsParams `form:"payouts" json:"payouts,omitempty"`
	// Settings related to the account's balance settlement timing.
	SettlementTiming *BalanceSettingsUpdatePaymentsSettlementTimingParams `form:"settlement_timing" json:"settlement_timing,omitempty"`
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
type BalanceSettingsUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Settings that apply to the [Payments Balance](https://docs.stripe.com/api/balance).
	Payments *BalanceSettingsUpdatePaymentsParams `form:"payments" json:"payments,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BalanceSettingsUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
type BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrency struct {
	// The ID of the FinancialAccount that funds will be transferred to during automatic transfers.
	PayoutMethod Currency `json:"payout_method"`
	// The maximum amount in minor units to transfer to the FinancialAccount. Only applicable when `type` is `transfer_up_to_amount`.
	TransferUpToAmount int64 `json:"transfer_up_to_amount"`
	// The type of automatic transfer rule.
	Type BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyType `json:"type"`
}

// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation for details.
type BalanceSettingsPaymentsPayoutsSchedule struct {
	// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
	Interval BalanceSettingsPaymentsPayoutsScheduleInterval `json:"interval"`
	// The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.
	MonthlyPayoutDays []int64 `json:"monthly_payout_days,omitempty"`
	// The days of the week when available funds are paid out, specified as an array, for example, [`monday`, `tuesday`]. Only shown if `interval` is weekly.
	WeeklyPayoutDays []BalanceSettingsPaymentsPayoutsScheduleWeeklyPayoutDay `json:"weekly_payout_days,omitempty"`
}

// Settings specific to the account's payouts.
type BalanceSettingsPaymentsPayouts struct {
	// Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.
	AutomaticTransferRulesByCurrency map[Currency][]*BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrency `json:"automatic_transfer_rules_by_currency"`
	// The minimum balance amount to retain per currency after automatic payouts. Only funds that exceed these amounts are paid out. Learn more about the [minimum balances for automatic payouts](https://docs.stripe.com/payouts/minimum-balances-for-automatic-payouts).
	MinimumBalanceByCurrency map[Currency]int64 `json:"minimum_balance_by_currency"`
	// Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation for details.
	Schedule *BalanceSettingsPaymentsPayoutsSchedule `json:"schedule"`
	// The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.
	StatementDescriptor string `json:"statement_descriptor"`
	// Whether the funds in this account can be paid out.
	Status BalanceSettingsPaymentsPayoutsStatus `json:"status"`
}

// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
type BalanceSettingsPaymentsSettlementTimingStartOfDay struct {
	// Hour at which the customized start of day begins according to the given timezone. Must be a [supported customized start of day hour](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Hour int64 `json:"hour"`
	// Minutes at which the customized start of day begins according to the given timezone. Must be either 0 or 30.
	Minutes int64 `json:"minutes"`
	// Timezone for the customized start of day. Must be a [supported customized start of day timezone](https://docs.stripe.com/connect/customized-start-of-day#available-timezones-and-cutoffs).
	Timezone string `json:"timezone"`
}
type BalanceSettingsPaymentsSettlementTiming struct {
	// The number of days charge funds are held before becoming available.
	DelayDays int64 `json:"delay_days"`
	// The number of days charge funds are held before becoming available. If present, overrides the default, or minimum available, for the account.
	DelayDaysOverride int64 `json:"delay_days_override,omitempty"`
	// Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](https://docs.stripe.com/connect/customized-start-of-day) documentation.
	StartOfDay *BalanceSettingsPaymentsSettlementTimingStartOfDay `json:"start_of_day"`
}
type BalanceSettingsPayments struct {
	// A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See [Understanding Connect account balances](https://docs.stripe.com/connect/account-balances) for details. The default value is `false` when [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts, otherwise `true`.
	DebitNegativeBalances bool `json:"debit_negative_balances"`
	// Settings specific to the account's payouts.
	Payouts          *BalanceSettingsPaymentsPayouts          `json:"payouts"`
	SettlementTiming *BalanceSettingsPaymentsSettlementTiming `json:"settlement_timing"`
}

// Options for customizing account balances and payout settings for a Stripe platform's connected accounts.
type BalanceSettings struct {
	APIResource
	// String representing the object's type. Objects of the same type share the same value.
	Object   string                   `json:"object"`
	Payments *BalanceSettingsPayments `json:"payments"`
}
