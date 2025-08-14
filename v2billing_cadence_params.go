//
//
// File generated from our OpenAPI spec
//
//

package stripe

// If provided, only Cadences that specifically reference the payer will be returned. Mutually exclusive with `test_clock`.
type V2BillingCadenceListPayerParams struct {
	// The ID of the Customer object. If provided, only Cadences that specifically reference the provided customer ID will be returned.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type *string `form:"type" json:"type"`
}

// List all the billing Cadences.
type V2BillingCadenceListParams struct {
	Params `form:"*"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// If provided, only Cadences that specifically reference the payer will be returned. Mutually exclusive with `test_clock`.
	Payer *V2BillingCadenceListPayerParams `form:"payer" json:"payer,omitempty"`
	// If provided, only Cadences that specifically reference the provided test clock ID (via the
	// customer's test clock) will be returned.
	// Mutually exclusive with `payer`.
	TestClock *string `form:"test_clock" json:"test_clock,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceBillingCycleDayTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingCadenceBillingCycleDayParams struct {
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceBillingCycleDayTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceBillingCycleMonthTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingCadenceBillingCycleMonthParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this will anchor to the last day of the month. If not provided,
	// this will default to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceBillingCycleMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceBillingCycleWeekTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingCadenceBillingCycleWeekParams struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601
	// week day numbering. If not provided, this will default to the day the
	// cadence was created.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceBillingCycleWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceBillingCycleYearTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingCadenceBillingCycleYearParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this will anchor to the last day of the month. If not provided,
	// this will default to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month,omitempty"`
	// The month to bill on from 1-12. If not provided, this will default to the
	// month the cadence was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceBillingCycleYearTimeParams `form:"time" json:"time,omitempty"`
}

// The billing cycle is the object that defines future billing cycle dates.
type V2BillingCadenceBillingCycleParams struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingCadenceBillingCycleDayParams `form:"day" json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between
	// cadence billings. For example, type=month and interval_count=3 bills every
	// 3 months. If this is not provided, it will default to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingCadenceBillingCycleMonthParams `form:"month" json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type *string `form:"type" json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingCadenceBillingCycleWeekParams `form:"week" json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingCadenceBillingCycleYearParams `form:"year" json:"year,omitempty"`
}

// The payer determines the entity financially responsible for the bill.
type V2BillingCadencePayerParams struct {
	// The ID of the Billing Profile object which determines how a bill will be paid. If provided, the created cadence will be
	// associated with the provided Billing Profile. If not provided, a new Billing Profile will be created and associated with the cadence.
	BillingProfile *string `form:"billing_profile" json:"billing_profile,omitempty"`
	// The ID of the Customer object.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type *string `form:"type" json:"type,omitempty"`
}

// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
// If no setting is provided here, the settings from the customer referenced on the payer will be used.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingCadenceSettingsBillParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of Settings to use.
	// If not provided, this will always default to the `live_version` specified on the setting, any time the settings are used.
	// Using a specific version here will prevent the settings from updating, and is discouraged for cadences.
	// To clear a pinned version, set the version to null.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingCadenceSettingsCollectionParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of Settings to use.
	// If not provided, this will always default to the `live_version` specified on the setting, any time the settings are used.
	// Using a specific version here will prevent the settings from updating, and is discouraged for cadences.
	// To clear a pinned version, set the version to null.
	Version *string `form:"version" json:"version,omitempty"`
}

// The settings associated with the cadence.
type V2BillingCadenceSettingsParams struct {
	// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
	// If no setting is provided here, the settings from the customer referenced on the payer will be used.
	// If no customer settings are present, the merchant default settings will be used.
	Bill *V2BillingCadenceSettingsBillParams `form:"bill" json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
	// If no customer settings are present, the merchant default settings will be used.
	Collection *V2BillingCadenceSettingsCollectionParams `form:"collection" json:"collection,omitempty"`
}

// Create a billing Cadence object.
type V2BillingCadenceParams struct {
	Params `form:"*"`
	// The billing cycle is the object that defines future billing cycle dates.
	BillingCycle *V2BillingCadenceBillingCycleParams `form:"billing_cycle" json:"billing_cycle,omitempty"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payer determines the entity financially responsible for the bill.
	Payer *V2BillingCadencePayerParams `form:"payer" json:"payer,omitempty"`
	// The settings associated with the cadence.
	Settings *V2BillingCadenceSettingsParams `form:"settings" json:"settings,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCadenceParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Cancel the billing cadence.
type V2BillingCadenceCancelParams struct {
	Params `form:"*"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceCreateBillingCycleDayTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingCadenceCreateBillingCycleDayParams struct {
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceCreateBillingCycleDayTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceCreateBillingCycleMonthTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingCadenceCreateBillingCycleMonthParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this will anchor to the last day of the month. If not provided,
	// this will default to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceCreateBillingCycleMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceCreateBillingCycleWeekTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingCadenceCreateBillingCycleWeekParams struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601
	// week day numbering. If not provided, this will default to the day the
	// cadence was created.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceCreateBillingCycleWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it will default to
// the time at which the cadence was created in UTC timezone.
type V2BillingCadenceCreateBillingCycleYearTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	// Will default to the minute the cadence was created in UTC timezone.
	Minute *int64 `form:"minute" json:"minute,omitempty"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingCadenceCreateBillingCycleYearParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this will anchor to the last day of the month. If not provided,
	// this will default to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month,omitempty"`
	// The month to bill on from 1-12. If not provided, this will default to the
	// month the cadence was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it will default to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingCadenceCreateBillingCycleYearTimeParams `form:"time" json:"time,omitempty"`
}

// The billing cycle is the object that defines future billing cycle dates.
type V2BillingCadenceCreateBillingCycleParams struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingCadenceCreateBillingCycleDayParams `form:"day" json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between
	// cadence billings. For example, type=month and interval_count=3 bills every
	// 3 months. If this is not provided, it will default to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingCadenceCreateBillingCycleMonthParams `form:"month" json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type *string `form:"type" json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingCadenceCreateBillingCycleWeekParams `form:"week" json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingCadenceCreateBillingCycleYearParams `form:"year" json:"year,omitempty"`
}

// The payer determines the entity financially responsible for the bill.
type V2BillingCadenceCreatePayerParams struct {
	// The ID of the Billing Profile object which determines how a bill will be paid. If provided, the created cadence will be
	// associated with the provided Billing Profile. If not provided, a new Billing Profile will be created and associated with the cadence.
	BillingProfile *string `form:"billing_profile" json:"billing_profile,omitempty"`
	// The ID of the Customer object.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type *string `form:"type" json:"type,omitempty"`
}

// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
// If no setting is provided here, the settings from the customer referenced on the payer will be used.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingCadenceCreateSettingsBillParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this will always default to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingCadenceCreateSettingsCollectionParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this will always default to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// The settings associated with the cadence.
type V2BillingCadenceCreateSettingsParams struct {
	// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
	// If no setting is provided here, the settings from the customer referenced on the payer will be used.
	// If no customer settings are present, the merchant default settings will be used.
	Bill *V2BillingCadenceCreateSettingsBillParams `form:"bill" json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
	// If no customer settings are present, the merchant default settings will be used.
	Collection *V2BillingCadenceCreateSettingsCollectionParams `form:"collection" json:"collection,omitempty"`
}

// Create a billing Cadence object.
type V2BillingCadenceCreateParams struct {
	Params `form:"*"`
	// The billing cycle is the object that defines future billing cycle dates.
	BillingCycle *V2BillingCadenceCreateBillingCycleParams `form:"billing_cycle" json:"billing_cycle"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful
	// for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payer determines the entity financially responsible for the bill.
	Payer *V2BillingCadenceCreatePayerParams `form:"payer" json:"payer"`
	// The settings associated with the cadence.
	Settings *V2BillingCadenceCreateSettingsParams `form:"settings" json:"settings,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCadenceCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a billing Cadence object.
type V2BillingCadenceRetrieveParams struct {
	Params `form:"*"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
}

// The payer determines the entity financially responsible for the bill.
type V2BillingCadenceUpdatePayerParams struct {
	// The ID of the Billing Profile object which determines how a bill will be paid.
	BillingProfile *string `form:"billing_profile" json:"billing_profile,omitempty"`
}

// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices. If null is provided, the current bill settings will be removed from the billing cadence.
type V2BillingCadenceUpdateSettingsBillParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of Settings to use.
	// If not provided, this will always default to the `live_version` specified on the setting, any time the settings are used.
	// Using a specific version here will prevent the settings from updating, and is discouraged for cadences.
	// To clear a pinned version, set the version to null.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments. If null is provided, the current collection settings will be removed from the billing cadence.
type V2BillingCadenceUpdateSettingsCollectionParams struct {
	// The ID of the referenced Settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of Settings to use.
	// If not provided, this will always default to the `live_version` specified on the setting, any time the settings are used.
	// Using a specific version here will prevent the settings from updating, and is discouraged for cadences.
	// To clear a pinned version, set the version to null.
	Version *string `form:"version" json:"version,omitempty"`
}

// The settings associated with the cadence.
type V2BillingCadenceUpdateSettingsParams struct {
	// Settings that configure bills generation, which includes calculating totals, tax, and presenting invoices. If null is provided, the current bill settings will be removed from the billing cadence.
	Bill *V2BillingCadenceUpdateSettingsBillParams `form:"bill" json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments. If null is provided, the current collection settings will be removed from the billing cadence.
	Collection *V2BillingCadenceUpdateSettingsCollectionParams `form:"collection" json:"collection,omitempty"`
}

// Update a billing Cadence object.
type V2BillingCadenceUpdateParams struct {
	Params `form:"*"`
	// Additional resource to include in the response.
	Include []*string `form:"include" json:"include,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The payer determines the entity financially responsible for the bill.
	Payer *V2BillingCadenceUpdatePayerParams `form:"payer" json:"payer,omitempty"`
	// The settings associated with the cadence.
	Settings *V2BillingCadenceUpdateSettingsParams `form:"settings" json:"settings,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingCadenceUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
