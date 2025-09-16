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

// v1IssuingCardholderService is used to invoke /v1/issuing/cardholders APIs.
type v1IssuingCardholderService struct {
	B   Backend
	Key string
}

// Creates a new Issuing Cardholder object that can be issued cards.
func (c v1IssuingCardholderService) Create(ctx context.Context, params *IssuingCardholderCreateParams) (*IssuingCardholder, error) {
	if params == nil {
		params = &IssuingCardholderCreateParams{}
	}
	params.Context = ctx
	cardholder := &IssuingCardholder{}
	err := c.B.Call(
		http.MethodPost, "/v1/issuing/cardholders", c.Key, params, cardholder)
	return cardholder, err
}

// Retrieves an Issuing Cardholder object.
func (c v1IssuingCardholderService) Retrieve(ctx context.Context, id string, params *IssuingCardholderRetrieveParams) (*IssuingCardholder, error) {
	if params == nil {
		params = &IssuingCardholderRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &IssuingCardholder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cardholder)
	return cardholder, err
}

// Updates the specified Issuing Cardholder object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingCardholderService) Update(ctx context.Context, id string, params *IssuingCardholderUpdateParams) (*IssuingCardholder, error) {
	if params == nil {
		params = &IssuingCardholderUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &IssuingCardholder{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cardholder)
	return cardholder, err
}

// Returns a list of Issuing Cardholder objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingCardholderService) List(ctx context.Context, listParams *IssuingCardholderListParams) Seq2[*IssuingCardholder, error] {
	if listParams == nil {
		listParams = &IssuingCardholderListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingCardholder, ListContainer, error) {
		list := &IssuingCardholderList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cardholders", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
