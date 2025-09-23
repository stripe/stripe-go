//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Stripe Tax details.
type V2BillingLicensedItemTaxDetails struct {
	// Product tax code (PTC).
	TaxCode string `json:"tax_code"`
}

// A Licensed Item represents a billable item whose pricing is based on license fees. You can use license fees
// to specify the pricing and create subscriptions to these items.
type V2BillingLicensedItem struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular billable item.
	// Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Stripe Tax details.
	TaxDetails *V2BillingLicensedItemTaxDetails `json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "seat" for Checkout to display "(price) per seat", or "environment" to display "(price) per environment".
	// Maximum length of 100 characters.
	UnitLabel string `json:"unit_label,omitempty"`
}
