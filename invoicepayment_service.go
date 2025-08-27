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

// v1InvoicePaymentService is used to invoke /v1/invoice_payments APIs.
type v1InvoicePaymentService struct {
	B   Backend
	Key string
}

// Retrieves the invoice payment with the given ID.
func (c v1InvoicePaymentService) Retrieve(ctx context.Context, id string, params *InvoicePaymentRetrieveParams) (*InvoicePayment, error) {
	if params == nil {
		params = &InvoicePaymentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/invoice_payments/%s", id)
	invoicepayment := &InvoicePayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoicepayment)
	return invoicepayment, err
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
func (c v1InvoicePaymentService) List(ctx context.Context, listParams *InvoicePaymentListParams) Seq2[*InvoicePayment, error] {
	if listParams == nil {
		listParams = &InvoicePaymentListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*InvoicePayment, ListContainer, error) {
		list := &InvoicePaymentList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/invoice_payments", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
