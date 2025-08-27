//
//
// File generated from our OpenAPI spec
//
//

// Package mandate provides the /v1/mandates APIs
package mandate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/mandates APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Mandate object.
func Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	return getC().Get(id, params)
}

// Retrieves a Mandate object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	path := stripe.FormatURLPath("/v1/mandates/%s", id)
	mandate := &stripe.Mandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mandate)
	return mandate, err
}

// Retrieves a list of Mandates for a given PaymentMethod.
func List(params *stripe.MandateListParams) *Iter {
	return getC().List(params)
}

// Retrieves a list of Mandates for a given PaymentMethod.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.MandateListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.MandateList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/mandates", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for mandates.
type Iter struct {
	*stripe.Iter
}

// Mandate returns the mandate which the iterator is currently pointing to.
func (i *Iter) Mandate() *stripe.Mandate {
	return i.Current().(*stripe.Mandate)
}

// MandateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) MandateList() *stripe.MandateList {
	return i.List().(*stripe.MandateList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
