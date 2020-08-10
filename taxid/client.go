// Package taxid provides the /customers/cus_1 APIs
package taxid

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /tax_ids APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tax id.
func New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().New(params)
}

// New creates a new tax id.
func (c Client) New(params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Customer must be set")
	}

	path := stripe.FormatURLPath("/v1/customers/%s/tax_ids", stripe.StringValue(params.Customer))
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, taxid)
	return taxid, err
}

// Get returns the details of a tax id.
func Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax id.
func (c Client) Get(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Customer must be set")
	}

	path := stripe.FormatURLPath("/v1/customers/%s/tax_ids/%s", stripe.StringValue(params.Customer), id)
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxid)
	return taxid, err
}

// Del removes a tax id.
func Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	return getC().Del(id, params)
}

// Del removes a tax id.
func (c Client) Del(id string, params *stripe.TaxIDParams) (*stripe.TaxID, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Customer must be set")
	}

	path := stripe.FormatURLPath("/v1/customers/%s/tax_ids/%s", stripe.StringValue(params.Customer), id)
	taxid := &stripe.TaxID{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, taxid)
	return taxid, err
}

// List returns a list of tax ids.
func List(params *stripe.TaxIDListParams) *Iter {
	return getC().List(params)
}

// List returns a list of tax ids.
func (c Client) List(listParams *stripe.TaxIDListParams) *Iter {
	path := stripe.FormatURLPath("/v1/customers/%s/tax_ids", stripe.StringValue(listParams.Customer))

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.TaxIDList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for tax ids.
type Iter struct {
	*stripe.Iter
}

// TaxID returns the tax id which the iterator is currently pointing to.
func (i *Iter) TaxID() *stripe.TaxID {
	return i.Current().(*stripe.TaxID)
}

// TaxIDList returns the current list object which the iterator is currently
// using. List objects will change as new API calls are made to continue
// pagination.
func (i *Iter) TaxIDList() *stripe.TaxIDList {
	return i.List().(*stripe.TaxIDList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
