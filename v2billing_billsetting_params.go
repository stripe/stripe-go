//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all BillSetting objects.
type V2BillingBillSettingListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Only return the settings with these lookup_keys, if any exist.
	// You can specify up to 10 lookup_keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys,omitempty"`
}

// Settings for calculating tax.
type V2BillingBillSettingCalculationTaxParams struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type *string `form:"type" json:"type"`
}

// Settings related to calculating a bill.
type V2BillingBillSettingCalculationParams struct {
	// Settings for calculating tax.
	Tax *V2BillingBillSettingCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingBillSettingInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingBillSettingInvoiceParams struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingBillSettingInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// Create a BillSetting object.
type V2BillingBillSettingParams struct {
	Params `form:"*"`
	// Settings related to calculating a bill.
	Calculation *V2BillingBillSettingCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// An optional customer-facing display name for the BillSetting object.
	// To remove the display name, set it to an empty string in the request.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Settings related to invoice behavior.
	Invoice *V2BillingBillSettingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
	// Optionally change the live version of the BillSetting. Providing `live_version = "latest"` will set the
	// BillSetting' `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// Settings for calculating tax.
type V2BillingBillSettingCreateCalculationTaxParams struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type *string `form:"type" json:"type"`
}

// Settings related to calculating a bill.
type V2BillingBillSettingCreateCalculationParams struct {
	// Settings for calculating tax.
	Tax *V2BillingBillSettingCreateCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingBillSettingCreateInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingBillSettingCreateInvoiceParams struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingBillSettingCreateInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// Create a BillSetting object.
type V2BillingBillSettingCreateParams struct {
	Params `form:"*"`
	// Settings related to calculating a bill.
	Calculation *V2BillingBillSettingCreateCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// An optional customer-facing display name for the CollectionSetting object.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Settings related to invoice behavior.
	Invoice *V2BillingBillSettingCreateInvoiceParams `form:"invoice" json:"invoice,omitempty"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// Retrieve a BillSetting object by ID.
type V2BillingBillSettingRetrieveParams struct {
	Params `form:"*"`
}

// Settings for calculating tax.
type V2BillingBillSettingUpdateCalculationTaxParams struct {
	// Determines if tax will be calculated automatically based on a PTC or manually based on rules defined by the merchant. Defaults to "manual".
	Type *string `form:"type" json:"type"`
}

// Settings related to calculating a bill.
type V2BillingBillSettingUpdateCalculationParams struct {
	// Settings for calculating tax.
	Tax *V2BillingBillSettingUpdateCalculationTaxParams `form:"tax" json:"tax,omitempty"`
}

// The amount of time until the invoice will be overdue for payment.
type V2BillingBillSettingUpdateInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30,
	// the invoice will be due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Settings related to invoice behavior.
type V2BillingBillSettingUpdateInvoiceParams struct {
	// The amount of time until the invoice will be overdue for payment.
	TimeUntilDue *V2BillingBillSettingUpdateInvoiceTimeUntilDueParams `form:"time_until_due" json:"time_until_due,omitempty"`
}

// Update fields on an existing BillSetting object.
type V2BillingBillSettingUpdateParams struct {
	Params `form:"*"`
	// Settings related to calculating a bill.
	Calculation *V2BillingBillSettingUpdateCalculationParams `form:"calculation" json:"calculation,omitempty"`
	// An optional customer-facing display name for the BillSetting object.
	// To remove the display name, set it to an empty string in the request.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Settings related to invoice behavior.
	Invoice *V2BillingBillSettingUpdateInvoiceParams `form:"invoice" json:"invoice,omitempty"`
	// The ID of the invoice rendering template to be used when generating invoices.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
	// Optionally change the live version of the BillSetting. Providing `live_version = "latest"` will set the
	// BillSetting' `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// A lookup key used to retrieve settings dynamically from a static string.
	// This may be up to 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}
