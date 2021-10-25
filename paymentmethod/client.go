//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethod provides the /payment_methods APIs
package paymentmethod

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /payment_methods APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payment method.
func New(params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().New(params)
}

// New creates a new payment method.
func (c Client) New(params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/payment_methods",
		c.Key,
		params,
		paymentmethod,
	)
	return paymentmethod, err
}

// Get returns the details of a payment method.
func Get(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payment method.
func (c Client) Get(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s", id)
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Update updates a payment method's properties.
func Update(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().Update(id, params)
}

// Update updates a payment method's properties.
func (c Client) Update(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s", id)
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Attach is the method for the `POST /v1/payment_methods/{payment_method}/attach` API.
func Attach(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	return getC().Attach(id, params)
}

// Attach is the method for the `POST /v1/payment_methods/{payment_method}/attach` API.
func (c Client) Attach(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s/attach", id)
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// Detach is the method for the `POST /v1/payment_methods/{payment_method}/detach` API.
func Detach(id string, params *stripe.PaymentMethodDetachParams) (*stripe.PaymentMethod, error) {
	return getC().Detach(id, params)
}

// Detach is the method for the `POST /v1/payment_methods/{payment_method}/detach` API.
func (c Client) Detach(id string, params *stripe.PaymentMethodDetachParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s/detach", id)
	paymentmethod := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentmethod)
	return paymentmethod, err
}

// List returns a list of payment methods.
func List(params *stripe.PaymentMethodListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payment methods.
func (c Client) List(listParams *stripe.PaymentMethodListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentMethodList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_methods", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment methods.
type Iter struct {
	*stripe.Iter
}

// PaymentMethod returns the payment method which the iterator is currently pointing to.
func (i *Iter) PaymentMethod() *stripe.PaymentMethod {
	return i.Current().(*stripe.PaymentMethod)
}

// PaymentMethodList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentMethodList() *stripe.PaymentMethodList {
	return i.List().(*stripe.PaymentMethodList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
