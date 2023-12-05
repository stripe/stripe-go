//
//
// File generated from our OpenAPI spec
//
//

// Package invoicepayment provides the /invoices/{invoice}/payments APIs
package invoicepayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /invoices/{invoice}/payments APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an invoice payment.
func Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	return getC().Get(id, params)
}

// Get returns the details of an invoice payment.
func (c Client) Get(id string, params *stripe.InvoicePaymentParams) (*stripe.InvoicePayment, error) {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/payments/%s",
		stripe.StringValue(params.Invoice),
		id,
	)
	invoicepayment := &stripe.InvoicePayment{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, invoicepayment)
	return invoicepayment, err
}

// List returns a list of invoice payments.
func List(params *stripe.InvoicePaymentListParams) *Iter {
	return getC().List(params)
}

// List returns a list of invoice payments.
func (c Client) List(listParams *stripe.InvoicePaymentListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/invoices/%s/payments",
		stripe.StringValue(listParams.Invoice),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoicePaymentList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   path,
				Key:    c.Key,
			}
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

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
