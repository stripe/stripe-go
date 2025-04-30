//
//
// File generated from our OpenAPI spec
//
//

// Package cardholder provides the /v1/issuing/cardholders APIs
// For more details, see: https://stripe.com/docs/api/?lang=go#issuing_cardholders
package cardholder

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/issuing/cardholders APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Issuing Cardholder object that can be issued cards.
func New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().New(params)
}

// Creates a new Issuing Cardholder object that can be issued cards.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/cardholders", c.Key, params, cardholder)
	return cardholder, err
}

// Retrieves an Issuing Cardholder object.
func Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().Get(id, params)
}

// Retrieves an Issuing Cardholder object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	path := stripe.FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cardholder)
	return cardholder, err
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().Update(id, params)
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	path := stripe.FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cardholder)
	return cardholder, err
}

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingCardholderListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingCardholderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCardholderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cardholders", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing cardholders.
type Iter struct {
	*stripe.Iter
}

// IssuingCardholder returns the issuing cardholder which the iterator is currently pointing to.
func (i *Iter) IssuingCardholder() *stripe.IssuingCardholder {
	return i.Current().(*stripe.IssuingCardholder)
}

// IssuingCardholderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardholderList() *stripe.IssuingCardholderList {
	return i.List().(*stripe.IssuingCardholderList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
