//
//
// File generated from our OpenAPI spec
//
//

// Package invoice provides the invoice related APIs
package invoice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke invoice related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an empty invoice using the billing group's settings.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingGroupsInvoiceParams) (*stripe.V2BillingGroupInvoice, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/groups/%s/invoices", stripe.StringValue(params.ID))
	groupinvoice := &stripe.V2BillingGroupInvoice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, groupinvoice)
	return groupinvoice, err
}
