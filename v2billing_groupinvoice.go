//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A Billing Group Invoice references an empty invoice created with the billing group's settings.
type V2BillingGroupInvoice struct {
	APIResource
	// Timestamp of when the BillingGroupInvoice object was created.
	Created time.Time `json:"created"`
	// The ID of the BillingGroupInvoice object.
	ID string `json:"id"`
	// The ID of the created Invoice object.
	InvoiceID string `json:"invoice_id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
