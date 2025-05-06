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

// v1CardService is used to invoke card related APIs.
type v1CardService struct {
	B   Backend
	Key string
}

// New creates a new card
func (c v1CardService) Create(ctx context.Context, params *CardCreateParams) (*Card, error) {
	if params == nil {
		params = &CardCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts", StringValue(params.Token), StringValue(
			params.Customer), StringValue(params.Account))
	card := &Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Get returns the details of a card.
func (c v1CardService) Retrieve(ctx context.Context, id string, params *CardRetrieveParams) (*Card, error) {
	if params == nil {
		params = &CardRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	card := &Card{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Update a specified source for a given customer.
func (c v1CardService) Update(ctx context.Context, id string, params *CardUpdateParams) (*Card, error) {
	if params == nil {
		params = &CardUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	card := &Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// Delete a specified source for a given customer.
func (c v1CardService) Delete(ctx context.Context, id string, params *CardDeleteParams) (*Card, error) {
	if params == nil {
		params = &CardDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	card := &Card{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}
func (c v1CardService) List(ctx context.Context, listParams *CardListParams) Seq2[*Card, error] {
	if listParams == nil {
		listParams = &CardListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts", StringValue(
			listParams.Account), StringValue(listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Card, ListContainer, error) {
		list := &CardList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
