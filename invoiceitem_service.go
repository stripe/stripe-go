//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1InvoiceItemService is used to invoke /v1/invoiceitems APIs.
type v1InvoiceItemService struct {
	B   Backend
	Key string
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
func (c v1InvoiceItemService) Create(ctx context.Context, params *InvoiceItemCreateParams) (*InvoiceItem, error) {
	if params == nil {
		params = &InvoiceItemCreateParams{}
	}
	params.Context = ctx
	invoiceitem := &InvoiceItem{}
	err := c.B.Call(
		http.MethodPost, "/v1/invoiceitems", c.Key, params, invoiceitem)
	return invoiceitem, err
}

// Retrieves the invoice item with the given ID.
func (c v1InvoiceItemService) Retrieve(ctx context.Context, id string, params *InvoiceItemRetrieveParams) (*InvoiceItem, error) {
	if params == nil {
		params = &InvoiceItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoiceitems/%s", id)
	invoiceitem := &InvoiceItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoiceitem)
	return invoiceitem, err
}

// Updates the amount or description of an invoice item on an upcoming invoice. Updating an invoice item is only possible before the invoice it's attached to is closed.
func (c v1InvoiceItemService) Update(ctx context.Context, id string, params *InvoiceItemUpdateParams) (*InvoiceItem, error) {
	if params == nil {
		params = &InvoiceItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoiceitems/%s", id)
	invoiceitem := &InvoiceItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, invoiceitem)
	return invoiceitem, err
}

// Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they're not attached to invoices, or if it's attached to a draft invoice.
func (c v1InvoiceItemService) Delete(ctx context.Context, id string, params *InvoiceItemDeleteParams) (*InvoiceItem, error) {
	if params == nil {
		params = &InvoiceItemDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoiceitems/%s", id)
	invoiceitem := &InvoiceItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, invoiceitem)
	return invoiceitem, err
}

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
func (c v1InvoiceItemService) List(ctx context.Context, listParams *InvoiceItemListParams) Seq2[*InvoiceItem, error] {
	if listParams == nil {
		listParams = &InvoiceItemListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*InvoiceItem, ListContainer, error) {
		list := &InvoiceItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/invoiceitems", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
