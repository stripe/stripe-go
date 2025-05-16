//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1InvoiceLineItemService is used to invoke /v1/invoices/{invoice}/lines APIs.
type v1InvoiceLineItemService struct {
	B   Backend
	Key string
}

// Updates an invoice's line item. Some fields, such as tax_amounts, only live on the invoice line item,
// so they can only be updated through this endpoint. Other fields, such as amount, live on both the invoice
// item and the invoice line item, so updates on this endpoint will propagate to the invoice item as well.
// Updating an invoice's line item is only possible before the invoice is finalized.
func (c v1InvoiceLineItemService) Update(ctx context.Context, id string, params *InvoiceLineItemUpdateParams) (*InvoiceLineItem, error) {
	if params == nil {
		params = &InvoiceLineItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/invoices/%s/lines/%s", StringValue(params.Invoice), id)
	invoicelineitem := &InvoiceLineItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoicelineitem)
	return invoicelineitem, err
}
