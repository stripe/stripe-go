// Package exchangerate provides the /exchange_rates APIs
package exchangerate

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /exchange_rates and exchangerates-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns an ExchangeRate for a given currency
func Get(currency string) (*stripe.ExchangeRate, error) {
	return getC().Get(currency)
}

func (c Client) Get(currency string) (*stripe.ExchangeRate, error) {
	path := stripe.FormatURLPath("/exchange_rates/%s", currency)
	exchangeRate := &stripe.ExchangeRate{}
	err := c.B.Call("GET", path, c.Key, nil, exchangeRate)

	return exchangeRate, err
}

// List lists available ExchangeRates.
func List(params *stripe.ExchangeRateListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.ExchangeRateListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ExchangeRateList{}
		err := c.B.CallRaw("GET", "/exchange_rates", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of ExchangeRates.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// ExchangeRate returns the most recent ExchangeRate
// visited by a call to Next.
func (i *Iter) ExchangeRate() *stripe.ExchangeRate {
	return i.Current().(*stripe.ExchangeRate)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
