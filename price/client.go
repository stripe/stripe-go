// Package price provides the /prices APIs
package price

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /prices APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new price.
func New(params *stripe.PriceParams) (*stripe.Price, error) {
	return getC().New(params)
}

// New creates a new price.
func (c Client) New(params *stripe.PriceParams) (*stripe.Price, error) {
	price := &stripe.Price{}
	err := c.B.Call(http.MethodPost, "/v1/prices", c.Key, params, price)
	return price, err
}

// Get returns the details of a price.
func Get(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	return getC().Get(id, params)
}

// Get returns the details of a price.
func (c Client) Get(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	path := stripe.FormatURLPath("/v1/prices/%s", id)
	price := &stripe.Price{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, price)
	return price, err
}

// Update updates a price's properties.
func Update(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	return getC().Update(id, params)
}

// Update updates a price's properties.
func (c Client) Update(id string, params *stripe.PriceParams) (*stripe.Price, error) {
	path := stripe.FormatURLPath("/v1/prices/%s", id)
	price := &stripe.Price{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, price)
	return price, err
}

// List returns a list of prices.
func List(params *stripe.PriceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of prices.
func (c Client) List(listParams *stripe.PriceListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.PriceList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/prices", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for prices.
type Iter struct {
	*stripe.Iter
}

// Price returns the price which the iterator is currently pointing to.
func (i *Iter) Price() *stripe.Price {
	return i.Current().(*stripe.Price)
}

// PriceList returns the current list object which the iterator is currently
// using. List objects will change as new API calls are made to continue
// pagination.
func (i *Iter) PriceList() *stripe.PriceList {
	return i.List().(*stripe.PriceList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
