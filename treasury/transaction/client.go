//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /v1/treasury/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/form"
)

// Client is used to invoke /v1/treasury/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing Transaction.
func Get(id string, params *stripe.TreasuryTransactionParams) (*stripe.TreasuryTransaction, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing Transaction.
func (c Client) Get(id string, params *stripe.TreasuryTransactionParams) (*stripe.TreasuryTransaction, error) {
	path := stripe.FormatURLPath("/v1/treasury/transactions/%s", id)
	transaction := &stripe.TreasuryTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Retrieves a list of Transaction objects.
func List(params *stripe.TreasuryTransactionListParams) *Iter {
	return getC().List(params)
}

// Retrieves a list of Transaction objects.
func (c Client) List(listParams *stripe.TreasuryTransactionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryTransactionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/transactions", c.Key, []byte(b.Encode()), p, list)

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
