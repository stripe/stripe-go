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

// v1GiftCardsTransactionService is used to invoke /v1/gift_cards/transactions APIs.
type v1GiftCardsTransactionService struct {
	B   Backend
	Key string
}

// Create a gift card transaction
func (c v1GiftCardsTransactionService) Create(ctx context.Context, params *GiftCardsTransactionCreateParams) (*GiftCardsTransaction, error) {
	transaction := &GiftCardsTransaction{}
	if params == nil {
		params = &GiftCardsTransactionCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/gift_cards/transactions", c.Key, params, transaction)
	return transaction, err
}

// Retrieves the gift card transaction.
func (c v1GiftCardsTransactionService) Retrieve(ctx context.Context, id string, params *GiftCardsTransactionRetrieveParams) (*GiftCardsTransaction, error) {
	path := FormatURLPath("/v1/gift_cards/transactions/%s", id)
	transaction := &GiftCardsTransaction{}
	if params == nil {
		params = &GiftCardsTransactionRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Update a gift card transaction
func (c v1GiftCardsTransactionService) Update(ctx context.Context, id string, params *GiftCardsTransactionUpdateParams) (*GiftCardsTransaction, error) {
	path := FormatURLPath("/v1/gift_cards/transactions/%s", id)
	transaction := &GiftCardsTransaction{}
	if params == nil {
		params = &GiftCardsTransactionUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Cancel a gift card transaction
func (c v1GiftCardsTransactionService) Cancel(ctx context.Context, id string, params *GiftCardsTransactionCancelParams) (*GiftCardsTransaction, error) {
	path := FormatURLPath("/v1/gift_cards/transactions/%s/cancel", id)
	transaction := &GiftCardsTransaction{}
	if params == nil {
		params = &GiftCardsTransactionCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Confirm a gift card transaction
func (c v1GiftCardsTransactionService) Confirm(ctx context.Context, id string, params *GiftCardsTransactionConfirmParams) (*GiftCardsTransaction, error) {
	path := FormatURLPath("/v1/gift_cards/transactions/%s/confirm", id)
	transaction := &GiftCardsTransaction{}
	if params == nil {
		params = &GiftCardsTransactionConfirmParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// List gift card transactions for a gift card
func (c v1GiftCardsTransactionService) List(ctx context.Context, listParams *GiftCardsTransactionListParams) Seq2[*GiftCardsTransaction, error] {
	if listParams == nil {
		listParams = &GiftCardsTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*GiftCardsTransaction, ListContainer, error) {
		list := &GiftCardsTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/gift_cards/transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
