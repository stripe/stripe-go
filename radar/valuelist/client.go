// Package valuelist provides API functions related to value lists.
//
// For more details, see: https://stripe.com/docs/api/radar/lists?lang=go
package valuelist

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /radar/value_lists APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new value list.
func New(params *stripe.ValueListParams) (*stripe.ValueList, error) {
	return getC().New(params)
}

// New creates a new value list.
func (c Client) New(params *stripe.ValueListParams) (*stripe.ValueList, error) {
	vl := &stripe.ValueList{}
	err := c.B.Call(http.MethodPost, "/radar/value_lists", c.Key, params, vl)
	return vl, err
}

// Get returns the details of a value list.
func Get(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	return getC().Get(id, params)
}

// Get returns the details of a value list.
func (c Client) Get(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	path := stripe.FormatURLPath("/radar/value_lists/%s", id)
	vl := &stripe.ValueList{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, vl)
	return vl, err
}

// Update updates a vl's properties.
func Update(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	return getC().Update(id, params)
}

// Update updates a vl's properties.
func (c Client) Update(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	path := stripe.FormatURLPath("/radar/value_lists/%s", id)
	vl := &stripe.ValueList{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, vl)
	return vl, err
}

// Del removes a value list.
func Del(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	return getC().Del(id, params)
}

// Del removes a value list.
func (c Client) Del(id string, params *stripe.ValueListParams) (*stripe.ValueList, error) {
	path := stripe.FormatURLPath("/radar/value_lists/%s", id)
	vl := &stripe.ValueList{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, vl)
	return vl, err
}

// List returns a list of vls.
func List(params *stripe.ValueListListParams) *Iter {
	return getC().List(params)
}

// List returns a list of vls.
func (c Client) List(listParams *stripe.ValueListListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ValueListList{}
		err := c.B.CallRaw(http.MethodGet, "/radar/value_lists", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for vls.
type Iter struct {
	*stripe.Iter
}

// ValueList returns the vl which the iterator is currently pointing to.
func (i *Iter) ValueList() *stripe.ValueList {
	return i.Current().(*stripe.ValueList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
