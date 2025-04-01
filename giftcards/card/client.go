//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /v1/gift_cards/cards APIs
package card

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/form"
)

// Client is used to invoke /v1/gift_cards/cards APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new gift card object.
func New(params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	return getC().New(params)
}

// Creates a new gift card object.
func (c Client) New(params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	card := &stripe.GiftCardsCard{}
	err := c.B.Call(http.MethodPost, "/v1/gift_cards/cards", c.Key, params, card)
	return card, err
}

// Retrieve a gift card by id
func Get(id string, params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	return getC().Get(id, params)
}

// Retrieve a gift card by id
func (c Client) Get(id string, params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/cards/%s", id)
	card := &stripe.GiftCardsCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Update a gift card
func Update(id string, params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	return getC().Update(id, params)
}

// Update a gift card
func (c Client) Update(id string, params *stripe.GiftCardsCardParams) (*stripe.GiftCardsCard, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/cards/%s", id)
	card := &stripe.GiftCardsCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Validates a gift card code, returning the matching gift card object if it exists.
func Validate(params *stripe.GiftCardsCardValidateParams) (*stripe.GiftCardsCard, error) {
	return getC().Validate(params)
}

// Validates a gift card code, returning the matching gift card object if it exists.
func (c Client) Validate(params *stripe.GiftCardsCardValidateParams) (*stripe.GiftCardsCard, error) {
	card := &stripe.GiftCardsCard{}
	err := c.B.Call(
		http.MethodPost, "/v1/gift_cards/cards/validate", c.Key, params, card)
	return card, err
}

// List gift cards for an account
func List(params *stripe.GiftCardsCardListParams) *Iter {
	return getC().List(params)
}

// List gift cards for an account
func (c Client) List(listParams *stripe.GiftCardsCardListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.GiftCardsCardList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/gift_cards/cards", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for gift cards cards.
type Iter struct {
	*stripe.Iter
}

// GiftCardsCard returns the gift cards card which the iterator is currently pointing to.
func (i *Iter) GiftCardsCard() *stripe.GiftCardsCard {
	return i.Current().(*stripe.GiftCardsCard)
}

// GiftCardsCardList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) GiftCardsCardList() *stripe.GiftCardsCardList {
	return i.List().(*stripe.GiftCardsCardList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
