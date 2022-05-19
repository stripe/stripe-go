//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpayment provides the /treasury/outbound_payments APIs
package outboundpayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /treasury/outbound_payments APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury outbound payment.
func New(params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().New(params)
}

// New creates a new treasury outbound payment.
func (c Client) New(params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/treasury/outbound_payments",
		c.Key,
		params,
		outboundpayment,
	)
	return outboundpayment, err
}

// Get returns the details of a treasury outbound payment.
func Get(id string, params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury outbound payment.
func (c Client) Get(id string, params *stripe.TreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath("/v1/treasury/outbound_payments/%s", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Cancel is the method for the `POST /v1/treasury/outbound_payments/{id}/cancel` API.
func Cancel(id string, params *stripe.TreasuryOutboundPaymentCancelParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/treasury/outbound_payments/{id}/cancel` API.
func (c Client) Cancel(id string, params *stripe.TreasuryOutboundPaymentCancelParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath("/v1/treasury/outbound_payments/%s/cancel", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// List returns a list of treasury outbound payments.
func List(params *stripe.TreasuryOutboundPaymentListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury outbound payments.
func (c Client) List(listParams *stripe.TreasuryOutboundPaymentListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryOutboundPaymentList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/outbound_payments", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury outbound payments.
type Iter struct {
	*stripe.Iter
}

// TreasuryOutboundPayment returns the treasury outbound payment which the iterator is currently pointing to.
func (i *Iter) TreasuryOutboundPayment() *stripe.TreasuryOutboundPayment {
	return i.Current().(*stripe.TreasuryOutboundPayment)
}

// TreasuryOutboundPaymentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryOutboundPaymentList() *stripe.TreasuryOutboundPaymentList {
	return i.List().(*stripe.TreasuryOutboundPaymentList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
