//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The frequency at which a cadence bills.
type V2BillingCadenceBillingCycleType string

// List of values that V2BillingCadenceBillingCycleType can take
const (
	V2BillingCadenceBillingCycleTypeDay   V2BillingCadenceBillingCycleType = "day"
	V2BillingCadenceBillingCycleTypeMonth V2BillingCadenceBillingCycleType = "month"
	V2BillingCadenceBillingCycleTypeWeek  V2BillingCadenceBillingCycleType = "week"
	V2BillingCadenceBillingCycleTypeYear  V2BillingCadenceBillingCycleType = "year"
)

// The type of the discount.
type V2BillingCadenceInvoiceDiscountRuleType string

// List of values that V2BillingCadenceInvoiceDiscountRuleType can take
const (
	V2BillingCadenceInvoiceDiscountRuleTypePercentOff V2BillingCadenceInvoiceDiscountRuleType = "percent_off"
)

// Max applications type of this discount, ex: indefinite.
type V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplicationsType string

// List of values that V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplicationsType can take
const (
	V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplicationsTypeIndefinite V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplicationsType = "indefinite"
)

// A string identifying the type of the payer. Currently the only supported value is `customer`.
type V2BillingCadencePayerType string

// List of values that V2BillingCadencePayerType can take
const (
	V2BillingCadencePayerTypeCustomer V2BillingCadencePayerType = "customer"
)

// The current status of the cadence.
type V2BillingCadenceStatus string

// List of values that V2BillingCadenceStatus can take
const (
	V2BillingCadenceStatusActive   V2BillingCadenceStatus = "active"
	V2BillingCadenceStatusCanceled V2BillingCadenceStatus = "canceled"
)

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleDayTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingCadenceBillingCycleDay struct {
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleDayTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleMonthTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingCadenceBillingCycleMonth struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleMonthTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleWeekTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingCadenceBillingCycleWeek struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601 week day numbering.
	DayOfWeek int64 `json:"day_of_week"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleWeekTime `json:"time"`
}

// The time at which the billing cycle ends.
type V2BillingCadenceBillingCycleYearTime struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour int64 `json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute int64 `json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second int64 `json:"second"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingCadenceBillingCycleYear struct {
	// The day to anchor the billing on for a type="month" billing cycle from 1-31.
	// If this number is greater than the number of days in the month being billed,
	// this will anchor to the last day of the month.
	DayOfMonth int64 `json:"day_of_month"`
	// The month to bill on from 1-12. If not provided, this will default to the month the cadence was created.
	MonthOfYear int64 `json:"month_of_year"`
	// The time at which the billing cycle ends.
	Time *V2BillingCadenceBillingCycleYearTime `json:"time"`
}

// The billing cycle is the object that defines future billing cycle dates.
type V2BillingCadenceBillingCycle struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingCadenceBillingCycleDay `json:"day"`
	// The number of intervals (specified in the interval attribute) between cadence billings. For example, type=month and interval_count=3 bills every 3 months.
	IntervalCount int64 `json:"interval_count"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingCadenceBillingCycleMonth `json:"month"`
	// The frequency at which a cadence bills.
	Type V2BillingCadenceBillingCycleType `json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingCadenceBillingCycleWeek `json:"week"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingCadenceBillingCycleYear `json:"year"`
}

// The maximum applications configuration for this discount.
type V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplications struct {
	// Max applications type of this discount, ex: indefinite.
	Type V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplicationsType `json:"type"`
}

// Details if the discount is a percentage off.
type V2BillingCadenceInvoiceDiscountRulePercentOff struct {
	// The maximum applications configuration for this discount.
	MaximumApplications *V2BillingCadenceInvoiceDiscountRulePercentOffMaximumApplications `json:"maximum_applications"`
	// Percent that will be taken off of the amount. For example, percent_off of 50.0 will make $100 amount $50 instead.
	PercentOff string `json:"percent_off"`
}

// The discount rules applied to all invoices for the cadence.
type V2BillingCadenceInvoiceDiscountRule struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// Details if the discount is a percentage off.
	PercentOff *V2BillingCadenceInvoiceDiscountRulePercentOff `json:"percent_off"`
	// The type of the discount.
	Type V2BillingCadenceInvoiceDiscountRuleType `json:"type"`
}

// The payer determines the entity financially responsible for the bill.
type V2BillingCadencePayer struct {
	// The ID of the Billing Profile object which determines how a bill will be paid.
	BillingProfile string `json:"billing_profile"`
	// The ID of the Customer object.
	Customer string `json:"customer"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type V2BillingCadencePayerType `json:"type"`
}

// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
type V2BillingCadenceSettingsBill struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version"`
}

// Settings that configure and manage the behavior of collecting payments.
type V2BillingCadenceSettingsCollection struct {
	// The ID of the referenced settings object.
	ID string `json:"id"`
	// Returns the Settings Version when the cadence is pinned to a specific version.
	Version string `json:"version"`
}

// The settings associated with the cadence.
type V2BillingCadenceSettings struct {
	// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices.
	Bill *V2BillingCadenceSettingsBill `json:"bill"`
	// Settings that configure and manage the behavior of collecting payments.
	Collection *V2BillingCadenceSettingsCollection `json:"collection"`
}
type V2BillingCadence struct {
	APIResource
	// The billing cycle is the object that defines future billing cycle dates.
	BillingCycle *V2BillingCadenceBillingCycle `json:"billing_cycle"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The discount rules applied to all invoices for the cadence.
	InvoiceDiscountRules []*V2BillingCadenceInvoiceDiscountRule `json:"invoice_discount_rules"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The date that the billing cadence will next bill. Null if the cadence is not active.
	NextBillingDate time.Time `json:"next_billing_date"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The payer determines the entity financially responsible for the bill.
	Payer *V2BillingCadencePayer `json:"payer"`
	// The settings associated with the cadence.
	Settings *V2BillingCadenceSettings `json:"settings"`
	// The current status of the cadence.
	Status V2BillingCadenceStatus `json:"status"`
	// The ID of the Test Clock.
	TestClock string `json:"test_clock"`
}
