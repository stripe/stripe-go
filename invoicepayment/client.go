//
//
// File generated from our OpenAPI spec
//
//

// Package invoicepayment provides the /v1/invoice_payments APIs
package invoicepayment

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/form"
)

// Client is used to invoke /v1/invoice_payments APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the invoice payment with the given ID.
func Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	return getC().Get(id, params)
}

// Retrieves the invoice payment with the given ID.
func (c Client) Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	path := stripe.FormatURLPath("/v1/invoice_payments/%s", id)
	invoicepayment := &stripe.InvoicePayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoicepayment)
	return invoicepayment, err
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
func List(params *stripe.InvoicePaymentListParams) *Iter {
	return getC().List(params)
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
func (c Client) List(listParams *stripe.InvoicePaymentListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoicePaymentList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoice_payments", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for invoice payments.
type Iter struct {
	*stripe.Iter
}

// InvoicePayment returns the invoice payment which the iterator is currently pointing to.
func (i *Iter) InvoicePayment() *stripe.InvoicePayment {
	return i.Current().(*stripe.InvoicePayment)
}

// InvoicePaymentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoicePaymentList() *stripe.InvoicePaymentList {
	return i.List().(*stripe.InvoicePaymentList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
