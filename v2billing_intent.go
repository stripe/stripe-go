//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Current status of the Billing Intent.
type V2BillingIntentStatus string

// List of values that V2BillingIntentStatus can take
const (
	V2BillingIntentStatusCanceled  V2BillingIntentStatus = "canceled"
	V2BillingIntentStatusCommitted V2BillingIntentStatus = "committed"
	V2BillingIntentStatusDraft     V2BillingIntentStatus = "draft"
	V2BillingIntentStatusReserved  V2BillingIntentStatus = "reserved"
)

// The frequency at which a cadence bills.
type V2BillingIntentCadenceDataBillingCycleType string

// List of values that V2BillingIntentCadenceDataBillingCycleType can take
const (
	V2BillingIntentCadenceDataBillingCycleTypeDay   V2BillingIntentCadenceDataBillingCycleType = "day"
	V2BillingIntentCadenceDataBillingCycleTypeMonth V2BillingIntentCadenceDataBillingCycleType = "month"
	V2BillingIntentCadenceDataBillingCycleTypeWeek  V2BillingIntentCadenceDataBillingCycleType = "week"
	V2BillingIntentCadenceDataBillingCycleTypeYear  V2BillingIntentCadenceDataBillingCycleType = "year"
)

// Breakdown of the amount for this Billing Intent.
type V2BillingIntentAmountDetails struct {
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// Amount of discount applied.
	Discount string `json:"discount"`
	// Amount of shipping charges.
	Shipping string `json:"shipping"`
	// Subtotal amount before tax and discounts.
	Subtotal string `json:"subtotal"`
	// Amount of tax.
	Tax string `json:"tax"`
	// Total amount for the Billing Intent.
	Total string `json:"total"`
}

// Timestamps for status transitions of the Billing Intent.
type V2BillingIntentStatusTransitions struct {
	// Time at which the Billing Intent was canceled.
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	// Time at which the Billing Intent was committed.
	CommittedAt time.Time `json:"committed_at,omitempty"`
	// Time at which the Billing Intent was drafted.
	DraftedAt time.Time `json:"drafted_at,omitempty"`
	// Time at which the Billing Intent was reserved.
	ReservedAt time.Time `json:"reserved_at,omitempty"`
}

// The time at which the billing cycle ends.
type V2BillingIntentCadenceDataBillingCycleDayTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingIntentCadenceDataBillingCycleDay struct {
	// The time at which the billing cycle ends.
	Time *V2BillingIntentCadenceDataBillingCycleDayTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingIntentCadenceDataBillingCycleMonthTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingIntentCadenceDataBillingCycleMonth struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to anchor the billing on for a type="month" billing cycle from
	// 1-12. Occurrences are calculated from the month anchor.
	MonthOfYear int64 `json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	Time *V2BillingIntentCadenceDataBillingCycleMonthTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingIntentCadenceDataBillingCycleWeekTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingIntentCadenceDataBillingCycleWeek struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601 week day numbering.
	DayOfWeek int64 `json:"day_of_week"`
	// The time at which the billing cycle ends.
	Time *V2BillingIntentCadenceDataBillingCycleWeekTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingIntentCadenceDataBillingCycleYearTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second,omitempty"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingIntentCadenceDataBillingCycleYear struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to bill on from 1-12. If not provided, this will default to the month the cadence was created.
	MonthOfYear int64 `json:"month_of_year"`
	// The time at which the billing cycle ends.
	Time *V2BillingIntentCadenceDataBillingCycleYearTime `json:"time"`
}

// The billing cycle configuration for this Cadence.
type V2BillingIntentCadenceDataBillingCycle struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingIntentCadenceDataBillingCycleDay `json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between cadence billings. For example, type=month and interval_count=3 bills every 3 months.
	IntervalCount int64 `json:"interval_count"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingIntentCadenceDataBillingCycleMonth `json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type V2BillingIntentCadenceDataBillingCycleType `json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingIntentCadenceDataBillingCycleWeek `json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingIntentCadenceDataBillingCycleYear `json:"year,omitempty"`
}

// Data for creating a new profile.
type V2BillingIntentCadenceDataPayerBillingProfileData struct {
	// The customer to associate with the profile.
	Customer string `json:"customer"`
	// The default payment method to use when billing this profile.
	// If none is provided, the customer `default_payment_method` will be used.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
}

// Information about the payer for this Cadence.
type V2BillingIntentCadenceDataPayer struct {
	// The ID of the Billing Profile object which determines how a bill will be paid.
	BillingProfile string `json:"billing_profile,omitempty"`
	// Data for creating a new profile.
	BillingProfileData *V2BillingIntentCadenceDataPayerBillingProfileData `json:"billing_profile_data,omitempty"`
}

// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
type V2BillingIntentCadenceDataSettingsBill struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
type V2BillingIntentCadenceDataSettingsCollection struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version,omitempty"`
}

// Settings for creating the Cadence.
type V2BillingIntentCadenceDataSettings struct {
	// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
	Bill *V2BillingIntentCadenceDataSettingsBill `json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	Collection *V2BillingIntentCadenceDataSettingsCollection `json:"collection,omitempty"`
}

// Data for creating a new Cadence.
type V2BillingIntentCadenceData struct {
	// The billing cycle configuration for this Cadence.
	BillingCycle *V2BillingIntentCadenceDataBillingCycle `json:"billing_cycle"`
	// Information about the payer for this Cadence.
	Payer *V2BillingIntentCadenceDataPayer `json:"payer"`
	// Settings for creating the Cadence.
	Settings *V2BillingIntentCadenceDataSettings `json:"settings,omitempty"`
}
type V2BillingIntent struct {
	APIResource
	// Breakdown of the amount for this Billing Intent.
	AmountDetails *V2BillingIntentAmountDetails `json:"amount_details"`
	// ID of an existing Cadence to use.
	Cadence string `json:"cadence,omitempty"`
	// Data for creating a new Cadence.
	CadenceData *V2BillingIntentCadenceData `json:"cadence_data,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Current status of the Billing Intent.
	Status V2BillingIntentStatus `json:"status"`
	// Timestamps for status transitions of the Billing Intent.
	StatusTransitions *V2BillingIntentStatusTransitions `json:"status_transitions"`
}
