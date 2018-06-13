package sku

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /skus APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new SKU.
// For more details see https://stripe.com/docs/api#create_sku.
func New(params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().New(params)
}

// New POSTs a new SKU.
// For more details see https://stripe.com/docs/api#create_sku.
func (c Client) New(params *stripe.SKUParams) (*stripe.SKU, error) {
	s := &stripe.SKU{}
	err := c.B.Call("POST", "/skus", c.Key, params, s)
	return s, err
}

// Update updates a SKU's properties.
// For more details see https://stripe.com/docs/api#update_sku.
func Update(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Update(id, params)
}

// Update updates a SKU's properties.
// For more details see https://stripe.com/docs/api#update_sku.
func (c Client) Update(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/skus/%s", id)
	s := &stripe.SKU{}
	err := c.B.Call("POST", path, c.Key, params, s)
	return s, err
}

// Get returns the details of an sku
// For more details see https://stripe.com/docs/api#retrieve_sku.
func Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/skus/%s", id)
	s := &stripe.SKU{}
	err := c.B.Call("GET", path, c.Key, params, s)
	return s, err
}

// List returns a list of skus.
// For more details see https://stripe.com/docs/api#list_skus
func List(params *stripe.SKUListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.SKUListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.SKUList{}
		err := c.B.CallRaw("GET", "/skus", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of SKUs.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// SKU returns the most recent SKU
// visited by a call to Next.
func (i *Iter) SKU() *stripe.SKU {
	return i.Current().(*stripe.SKU)
}

// Delete destroys a SKU.
// For more details see https://stripe.com/docs/api#delete_sku.
func Del(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Del(id, params)
}

// Delete destroys a SKU.
// For more details see https://stripe.com/docs/api#delete_sku.
func (c Client) Del(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/skus/%s", id)
	s := &stripe.SKU{}
	err := c.B.Call("DELETE", path, c.Key, params, s)

	return s, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
