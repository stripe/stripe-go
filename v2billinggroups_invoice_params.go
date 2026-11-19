//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Create an empty invoice using the billing group's settings.
type V2BillingGroupsInvoiceParams struct {
	Params `form:"*"`
	// The ID of the BillingGroup object to create an invoice for.
	ID *string `form:"-" json:"-"` // Included in URL
}

// Create an empty invoice using the billing group's settings.
type V2BillingGroupsInvoiceCreateParams struct {
	Params `form:"*"`
	// The ID of the BillingGroup object to create an invoice for.
	ID *string `form:"-" json:"-"` // Included in URL
}
