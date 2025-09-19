//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
type V2BillingBillSettingCalculationTaxType string

// List of values that V2BillingBillSettingCalculationTaxType can take
const (
	V2BillingBillSettingCalculationTaxTypeAutomatic V2BillingBillSettingCalculationTaxType = "automatic"
	V2BillingBillSettingCalculationTaxTypeManual    V2BillingBillSettingCalculationTaxType = "manual"
)

// The interval unit for the time until due.
type V2BillingBillSettingInvoiceTimeUntilDueInterval string

// List of values that V2BillingBillSettingInvoiceTimeUntilDueInterval can take
const (
	V2BillingBillSettingInvoiceTimeUntilDueIntervalDay   V2BillingBillSettingInvoiceTimeUntilDueInterval = "day"
	V2BillingBillSettingInvoiceTimeUntilDueIntervalMonth V2BillingBillSettingInvoiceTimeUntilDueInterval = "month"
	V2BillingBillSettingInvoiceTimeUntilDueIntervalWeek  V2BillingBillSettingInvoiceTimeUntilDueInterval = "week"
	V2BillingBillSettingInvoiceTimeUntilDueIntervalYear  V2BillingBillSettingInvoiceTimeUntilDueInterval = "year"
)

// Settings for calculating tax.
type V2BillingBillSettingCalculationTax struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type V2BillingBillSettingCalculationTaxType `json:"type"`
}

// Settings related to calculating a bill.
type V2BillingBillSettingCalculation struct {
	// Settings for calculating tax.
	Tax *V2BillingBillSettingCalculationTax `json:"tax,omitempty"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingBillSettingInvoiceTimeUntilDue struct {
	// The interval unit for the time until due.
	Interval V2BillingBillSettingInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount int64 `json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingBillSettingInvoice struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingBillSettingInvoiceTimeUntilDue `json:"time_until_due,omitempty"`
}

// BillSetting is responsible for settings which dictate generating bills, which include settings for calculating totals on bills, tax on bill items, as well as how to generate and present invoices.
type V2BillingBillSetting struct {
	APIResource
	// Settings related to calculating a bill.
	Calculation *V2BillingBillSettingCalculation `json:"calculation,omitempty"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// An optional field for adding a display name for the BillSetting object.
	DisplayName string `json:"display_name,omitempty"`
	// The ID of the BillSetting object.
	ID string `json:"id"`
	// Settings related to invoice behavior.
	Invoice *V2BillingBillSettingInvoice `json:"invoice,omitempty"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate string `json:"invoice_rendering_template,omitempty"`
	// The latest version of the current settings object. This will be
	// Updated every time an attribute of the settings is updated.
	LatestVersion string `json:"latest_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The current live version of the settings object. This can be different from
	// latest_version if settings are updated without setting live_version='latest'.
	LiveVersion string `json:"live_version"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
