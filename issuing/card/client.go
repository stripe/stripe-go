//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /v1/issuing/cards APIs
package card

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/issuing/cards APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an Issuing Card object.
func New(params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().New(params)
}

// Creates an Issuing Card object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, "/v1/issuing/cards", c.Key, params, card)
	return card, err
}

// Retrieves an Issuing Card object.
func Get(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().Get(id, params)
}

// Retrieves an Issuing Card object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath("/v1/issuing/cards/%s", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Updates the specified Issuing Card object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	return getC().Update(id, params)
}

// Updates the specified Issuing Card object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingCardParams) (*stripe.IssuingCard, error) {
	path := stripe.FormatURLPath("/v1/issuing/cards/%s", id)
	card := &stripe.IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Returns a list of Issuing Card objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingCardListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Issuing Card objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingCardListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCardList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cards", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing cards.
type Iter struct {
	*stripe.Iter
}

// IssuingCard returns the issuing card which the iterator is currently pointing to.
func (i *Iter) IssuingCard() *stripe.IssuingCard {
	return i.Current().(*stripe.IssuingCard)
}

// IssuingCardList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardList() *stripe.IssuingCardList {
	return i.List().(*stripe.IssuingCardList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
