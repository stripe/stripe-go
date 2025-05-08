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

// v1IssuingCardService is used to invoke /v1/issuing/cards APIs.
type v1IssuingCardService struct {
	B   Backend
	Key string
}

// Creates an Issuing Card object.
func (c v1IssuingCardService) Create(ctx context.Context, params *IssuingCardCreateParams) (*IssuingCard, error) {
	if params == nil {
		params = &IssuingCardCreateParams{}
	}
	params.Context = ctx
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, "/v1/issuing/cards", c.Key, params, card)
	return card, err
}

// Retrieves an Issuing Card object.
func (c v1IssuingCardService) Retrieve(ctx context.Context, id string, params *IssuingCardRetrieveParams) (*IssuingCard, error) {
	if params == nil {
		params = &IssuingCardRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/cards/%s", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Updates the specified Issuing Card object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingCardService) Update(ctx context.Context, id string, params *IssuingCardUpdateParams) (*IssuingCard, error) {
	if params == nil {
		params = &IssuingCardUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/cards/%s", id)
	card := &IssuingCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Returns a list of Issuing Card objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingCardService) List(ctx context.Context, listParams *IssuingCardListParams) Seq2[*IssuingCard, error] {
	if listParams == nil {
		listParams = &IssuingCardListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingCard, ListContainer, error) {
		list := &IssuingCardList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cards", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
