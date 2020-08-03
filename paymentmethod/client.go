// Package paymentmethod provides the /payment_methods APIs
package paymentmethod

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke pms APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Attach attaches a PaymentMethod.
func Attach(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	return getC().Attach(id, params)
}

// Attach attaches a PaymentMethod.
func (c Client) Attach(id string, params *stripe.PaymentMethodAttachParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s/attach", id)
	pm := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pm)
	return pm, err
}

// Detach detaches a PaymentMethod.
func Detach(id string, params *stripe.PaymentMethodDetachParams) (*stripe.PaymentMethod, error) {
	return getC().Detach(id, params)
}

// Detach detaches a PaymentMethod.
func (c Client) Detach(id string, params *stripe.PaymentMethodDetachParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s/detach", id)
	pm := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pm)
	return pm, err
}

// Get returns the details of a PaymentMethod.
func Get(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().Get(id, params)
}

// Get returns the details of a PaymentMethod.
func (c Client) Get(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s", id)
	fee := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fee)
	return fee, err
}

// List returns a list of PaymentMethods.
func List(params *stripe.PaymentMethodListParams) *Iter {
	return getC().List(params)
}

// List returns a list of PaymentMethods.
func (c Client) List(listParams *stripe.PaymentMethodListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.PaymentMethodList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_methods", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// New creates a new PaymentMethod.
func New(params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().New(params)
}

// New creates a new PaymentMethod.
func (c Client) New(params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	pm := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, "/v1/payment_methods", c.Key, params, pm)
	return pm, err
}

// Update updates a PaymentMethod.
func Update(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	return getC().Update(id, params)
}

// Update updates a PaymentMethod.
func (c Client) Update(id string, params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	path := stripe.FormatURLPath("/v1/payment_methods/%s", id)
	pm := &stripe.PaymentMethod{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pm)
	return pm, err
}

// Iter is an iterator for PaymentMethods.
type Iter struct {
	*stripe.Iter
}

// PaymentMethod returns the application fee which the iterator is currently pointing to.
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
