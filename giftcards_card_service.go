//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1GiftCardsCardService is used to invoke /v1/gift_cards/cards APIs.
type v1GiftCardsCardService struct {
	B   Backend
	Key string
}

// Creates a new gift card object.
func (c v1GiftCardsCardService) Create(ctx context.Context, params *GiftCardsCardCreateParams) (*GiftCardsCard, error) {
	card := &GiftCardsCard{}
	if params == nil {
		params = &GiftCardsCardCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, "/v1/gift_cards/cards", c.Key, params, card)
	return card, err
}

// Retrieve a gift card by id
func (c v1GiftCardsCardService) Retrieve(ctx context.Context, id string, params *GiftCardsCardRetrieveParams) (*GiftCardsCard, error) {
	path := FormatURLPath("/v1/gift_cards/cards/%s", id)
	card := &GiftCardsCard{}
	if params == nil {
		params = &GiftCardsCardRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Update a gift card
func (c v1GiftCardsCardService) Update(ctx context.Context, id string, params *GiftCardsCardUpdateParams) (*GiftCardsCard, error) {
	path := FormatURLPath("/v1/gift_cards/cards/%s", id)
	card := &GiftCardsCard{}
	if params == nil {
		params = &GiftCardsCardUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Validates a gift card code, returning the matching gift card object if it exists.
func (c v1GiftCardsCardService) Validate(ctx context.Context, params *GiftCardsCardValidateParams) (*GiftCardsCard, error) {
	card := &GiftCardsCard{}
	if params == nil {
		params = &GiftCardsCardValidateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/gift_cards/cards/validate", c.Key, params, card)
	return card, err
}

// List gift cards for an account
func (c v1GiftCardsCardService) List(ctx context.Context, listParams *GiftCardsCardListParams) Seq2[*GiftCardsCard, error] {
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*GiftCardsCard, ListContainer, error) {
		list := &GiftCardsCardList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/gift_cards/cards", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
