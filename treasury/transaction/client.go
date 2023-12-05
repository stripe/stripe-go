//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /treasury/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /treasury/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a treasury transaction.
func Get(id string, params *stripe.TreasuryTransactionParams) (*stripe.TreasuryTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury transaction.
func (c Client) Get(id string, params *stripe.TreasuryTransactionParams) (*stripe.TreasuryTransaction, error) {
	path := stripe.FormatURLPath("/v1/treasury/transactions/%s", id)
	transaction := &stripe.TreasuryTransaction{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// List returns a list of treasury transactions.
func List(params *stripe.TreasuryTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury transactions.
func (c Client) List(listParams *stripe.TreasuryTransactionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryTransactionList{}
			err := c.B.Call(stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/treasury/transactions",
				Key:    c.Key,
				Params: p,
				Body:   b,
			},
				list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury transactions.
type Iter struct {
	*stripe.Iter
}

// TreasuryTransaction returns the treasury transaction which the iterator is currently pointing to.
func (i *Iter) TreasuryTransaction() *stripe.TreasuryTransaction {
	return i.Current().(*stripe.TreasuryTransaction)
}

// TreasuryTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryTransactionList() *stripe.TreasuryTransactionList {
	return i.List().(*stripe.TreasuryTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
