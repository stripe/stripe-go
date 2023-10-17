//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /gift_cards/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /gift_cards/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new gift cards transaction.
func New(params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	return getC().New(params)
}

// New creates a new gift cards transaction.
func (c Client) New(params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	transaction := &stripe.GiftCardsTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/gift_cards/transactions",
		c.Key,
		params,
		transaction,
	)
	return transaction, err
}

// Get returns the details of a gift cards transaction.
func Get(id string, params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a gift cards transaction.
func (c Client) Get(id string, params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/transactions/%s", id)
	transaction := &stripe.GiftCardsTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Update updates a gift cards transaction's properties.
func Update(id string, params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	return getC().Update(id, params)
}

// Update updates a gift cards transaction's properties.
func (c Client) Update(id string, params *stripe.GiftCardsTransactionParams) (*stripe.GiftCardsTransaction, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/transactions/%s", id)
	transaction := &stripe.GiftCardsTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Cancel is the method for the `POST /v1/gift_cards/transactions/{id}/cancel` API.
func Cancel(id string, params *stripe.GiftCardsTransactionCancelParams) (*stripe.GiftCardsTransaction, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/gift_cards/transactions/{id}/cancel` API.
func (c Client) Cancel(id string, params *stripe.GiftCardsTransactionCancelParams) (*stripe.GiftCardsTransaction, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/transactions/%s/cancel", id)
	transaction := &stripe.GiftCardsTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Confirm is the method for the `POST /v1/gift_cards/transactions/{id}/confirm` API.
func Confirm(id string, params *stripe.GiftCardsTransactionConfirmParams) (*stripe.GiftCardsTransaction, error) {
	return getC().Confirm(id, params)
}

// Confirm is the method for the `POST /v1/gift_cards/transactions/{id}/confirm` API.
func (c Client) Confirm(id string, params *stripe.GiftCardsTransactionConfirmParams) (*stripe.GiftCardsTransaction, error) {
	path := stripe.FormatURLPath("/v1/gift_cards/transactions/%s/confirm", id)
	transaction := &stripe.GiftCardsTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// List returns a list of gift cards transactions.
func List(params *stripe.GiftCardsTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of gift cards transactions.
func (c Client) List(listParams *stripe.GiftCardsTransactionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.GiftCardsTransactionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/gift_cards/transactions", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for gift cards transactions.
type Iter struct {
	*stripe.Iter
}

// GiftCardsTransaction returns the gift cards transaction which the iterator is currently pointing to.
func (i *Iter) GiftCardsTransaction() *stripe.GiftCardsTransaction {
	return i.Current().(*stripe.GiftCardsTransaction)
}

// GiftCardsTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) GiftCardsTransactionList() *stripe.GiftCardsTransactionList {
	return i.List().(*stripe.GiftCardsTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
