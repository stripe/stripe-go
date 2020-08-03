// Package taxrate provides the /tax_rates APIs
package taxrate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /tax_rates APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tr.
func New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().New(params)
}

// New creates a new tr.
func (c Client) New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	tr := &stripe.TaxRate{}
	err := c.B.Call(http.MethodPost, "/v1/tax_rates", c.Key, params, tr)
	return tr, err
}

// Get returns the details of a tax rate.
func Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax rate.
func (c Client) Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	path := stripe.FormatURLPath("/v1/tax_rates/%s", id)
	tr := &stripe.TaxRate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, tr)
	return tr, err
}

// Update updates a tax rate's properties.
func Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().Update(id, params)
}

// Update updates a tax rate's properties.
func (c Client) Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	path := stripe.FormatURLPath("/v1/tax_rates/%s", id)
	tr := &stripe.TaxRate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, tr)
	return tr, err
}

// Del removes a tax rate.
func Del(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().Del(id, params)
}

// Del removes a tax rate.
func (c Client) Del(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	path := stripe.FormatURLPath("/v1/tax_rates/%s", id)
	tr := &stripe.TaxRate{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, tr)
	return tr, err
}

// List returns a list of trs.
func List(params *stripe.TaxRateListParams) *Iter {
	return getC().List(params)
}

// List returns a list of trs.
func (c Client) List(listParams *stripe.TaxRateListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.TaxRateList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/tax_rates", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for trs.
type Iter struct {
	*stripe.Iter
}

// TaxRate returns the tr which the iterator is currently pointing to.
func (i *Iter) TaxRate() *stripe.TaxRate {
	return i.Current().(*stripe.TaxRate)
}

// TaxRateList returns the current list object which the iterator is currently
// using. List objects will change as new API calls are made to continue
// pagination.
func (i *Iter) TaxRateList() *stripe.TaxRateList {
	return i.List().(*stripe.TaxRateList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
