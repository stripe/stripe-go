//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Optional array of Meter segments to filter event dimension keys for billing.
type V2BillingMeteredItemMeterSegmentCondition struct {
	// A Meter dimension.
	Dimension string `json:"dimension"`
	// To count usage towards this metered item, the dimension must have this value.
	Value string `json:"value"`
}

// Stripe Tax details.
type V2BillingMeteredItemTaxDetails struct {
	// Product tax code (PTC).
	TaxCode string `json:"tax_code"`
}

// A Metered Item represents a billable item whose pricing is based on usage, measured by a meter. You can use rate cards
// to specify the pricing and create subscriptions to these items.
type V2BillingMeteredItem struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Optional array of Meter dimensions to group event dimension keys for invoice line items.
	InvoicePresentationDimensions []string `json:"invoice_presentation_dimensions"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular billable item.
	// Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// ID of the Meter that measures usage for this Metered Item.
	Meter string `json:"meter"`
	// Optional array of Meter segments to filter event dimension keys for billing.
	MeterSegmentConditions []*V2BillingMeteredItemMeterSegmentCondition `json:"meter_segment_conditions"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Stripe Tax details.
	TaxDetails *V2BillingMeteredItemTaxDetails `json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "CPU-hour" for Checkout to display "(price) per CPU-hour", or "1 million events" to display "(price) per 1
	// million events".
	// Maximum length of 100 characters.
	UnitLabel string `json:"unit_label,omitempty"`
}
