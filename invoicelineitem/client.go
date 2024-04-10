//
//
// File generated from our OpenAPI spec
//
//

// Package invoicelineitem provides the /invoices/{invoice}/lines APIs
package invoicelineitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v78"
)

// Client is used to invoke /invoices/{invoice}/lines APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Update updates an invoice line item's properties.
func Update(id string, params *stripe.InvoiceLineItemParams) (*stripe.InvoiceLineItem, error) {
	return getC().Update(id, params)
}

// Update updates an invoice line item's properties.
func (c Client) Update(id string, params *stripe.InvoiceLineItemParams) (*stripe.InvoiceLineItem, error) {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/lines/%s",
		stripe.StringValue(params.Invoice),
		id,
	)
	invoicelineitem := &stripe.InvoiceLineItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoicelineitem)
	return invoicelineitem, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
