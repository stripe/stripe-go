//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Stripe Tax details.
type V2BillingOneTimeItemTaxDetails struct {
	// Product tax code (PTC).
	TaxCode string `json:"tax_code"`
}

// A one-time item represents a product that can be charged as a one-time line item,
// used for overage charges when custom pricing unit credits are exhausted.
type V2BillingOneTimeItem struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Description that customers see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular one-time item.
	// Maximum length of 200 characters.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Stripe Tax details.
	TaxDetails *V2BillingOneTimeItemTaxDetails `json:"tax_details,omitempty"`
	// The unit to use when displaying prices for this one-time item. For example, set this field
	// to "credit" for the display to show "(price) per credit".
	// Maximum length of 100 characters.
	UnitLabel string `json:"unit_label,omitempty"`
}
