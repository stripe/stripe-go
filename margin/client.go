//
//
// File generated from our OpenAPI spec
//
//

// Package margin provides the /billing/margins APIs
package margin

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/form"
)

// Client is used to invoke /billing/margins APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a margin object to be used with invoices, invoice items, and invoice line items for a customer to represent a partner discount. A margin has a percent_off which is the percent that will be taken off the subtotal after all items and other discounts and promotions) of any invoices for a customer. Calculation of prorations do not include any partner margins applied on the original invoice item.
func New(params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().New(params)
}

// Create a margin object to be used with invoices, invoice items, and invoice line items for a customer to represent a partner discount. A margin has a percent_off which is the percent that will be taken off the subtotal after all items and other discounts and promotions) of any invoices for a customer. Calculation of prorations do not include any partner margins applied on the original invoice item.
func (c Client) New(params *stripe.MarginParams) (*stripe.Margin, error) {
	margin := &stripe.Margin{}
	err := c.B.Call(http.MethodPost, "/v1/billing/margins", c.Key, params, margin)
	return margin, err
}

// Retrieve a margin object with the given ID.
func Get(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().Get(id, params)
}

// Retrieve a margin object with the given ID.
func (c Client) Get(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	path := stripe.FormatURLPath("/v1/billing/margins/%s", id)
	margin := &stripe.Margin{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, margin)
	return margin, err
}

// Update the specified margin object. Certain fields of the margin object are not editable.
func Update(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().Update(id, params)
}

// Update the specified margin object. Certain fields of the margin object are not editable.
func (c Client) Update(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	path := stripe.FormatURLPath("/v1/billing/margins/%s", id)
	margin := &stripe.Margin{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, margin)
	return margin, err
}

// Retrieve a list of your margins.
func List(params *stripe.MarginListParams) *Iter {
	return getC().List(params)
}

// Retrieve a list of your margins.
func (c Client) List(listParams *stripe.MarginListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.MarginList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/billing/margins", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for margins.
type Iter struct {
	*stripe.Iter
}

// Margin returns the margin which the iterator is currently pointing to.
func (i *Iter) Margin() *stripe.Margin {
	return i.Current().(*stripe.Margin)
}

// MarginList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) MarginList() *stripe.MarginList {
	return i.List().(*stripe.MarginList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
