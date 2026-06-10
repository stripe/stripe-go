//
//
// File generated from our OpenAPI spec
//
//

// Package taxfund provides the /v1/tax_funds APIs
package taxfund

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/tax_funds APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a tax fund object by its ID.
func Get(id string, params *stripe.TaxFundParams) (*stripe.TaxFund, error) {
	return getC().Get(id, params)
}

// Retrieves a tax fund object by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TaxFundParams) (*stripe.TaxFund, error) {
	path := stripe.FormatURLPath("/v1/tax_funds/%s", id)
	taxfund := &stripe.TaxFund{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxfund)
	return taxfund, err
}

// Returns a list of tax funds in reverse chronological order.
func List(params *stripe.TaxFundListParams) *Iter {
	return getC().List(params)
}

// Returns a list of tax funds in reverse chronological order.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxFundListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxFundList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax_funds", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax funds.
type Iter struct {
	*stripe.Iter
}

// TaxFund returns the tax fund which the iterator is currently pointing to.
func (i *Iter) TaxFund() *stripe.TaxFund {
	return i.Current().(*stripe.TaxFund)
}

// TaxFundList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxFundList() *stripe.TaxFundList {
	return i.List().(*stripe.TaxFundList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
