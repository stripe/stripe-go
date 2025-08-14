//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
type V2BillingBillSettingVersionCalculationTaxType string

// List of values that V2BillingBillSettingVersionCalculationTaxType can take
const (
	V2BillingBillSettingVersionCalculationTaxTypeAutomatic V2BillingBillSettingVersionCalculationTaxType = "automatic"
	V2BillingBillSettingVersionCalculationTaxTypeManual    V2BillingBillSettingVersionCalculationTaxType = "manual"
)

// The interval unit for the time until due.
type V2BillingBillSettingVersionInvoiceTimeUntilDueInterval string

// List of values that V2BillingBillSettingVersionInvoiceTimeUntilDueInterval can take
const (
	V2BillingBillSettingVersionInvoiceTimeUntilDueIntervalDay   V2BillingBillSettingVersionInvoiceTimeUntilDueInterval = "day"
	V2BillingBillSettingVersionInvoiceTimeUntilDueIntervalMonth V2BillingBillSettingVersionInvoiceTimeUntilDueInterval = "month"
	V2BillingBillSettingVersionInvoiceTimeUntilDueIntervalWeek  V2BillingBillSettingVersionInvoiceTimeUntilDueInterval = "week"
	V2BillingBillSettingVersionInvoiceTimeUntilDueIntervalYear  V2BillingBillSettingVersionInvoiceTimeUntilDueInterval = "year"
)

// Settings for calculating tax.
type V2BillingBillSettingVersionCalculationTax struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type V2BillingBillSettingVersionCalculationTaxType `json:"type"`
}

// Settings related to calculating a bill.
type V2BillingBillSettingVersionCalculation struct {
	// Settings for calculating tax.
	Tax *V2BillingBillSettingVersionCalculationTax `json:"tax"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingBillSettingVersionInvoiceTimeUntilDue struct {
	// The interval unit for the time until due.
	Interval V2BillingBillSettingVersionInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount int64 `json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingBillSettingVersionInvoice struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingBillSettingVersionInvoiceTimeUntilDue `json:"time_until_due"`
}
type V2BillingBillSettingVersion struct {
	APIResource
	// Settings related to calculating a bill.
	Calculation *V2BillingBillSettingVersionCalculation `json:"calculation"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the BillSettingVersion object.
	ID string `json:"id"`
	// Settings related to invoice behavior.
	Invoice *V2BillingBillSettingVersionInvoice `json:"invoice"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate string `json:"invoice_rendering_template"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
