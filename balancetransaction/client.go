// Package balancetransaction provides the /balance_transactions APIs
package balancetransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /balance_transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a balance transaction.
func Get(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a balance transaction.
func (c Client) Get(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	path := stripe.FormatURLPath("/v1/balance_transactions/%s", id)
	transaction := &stripe.BalanceTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// List returns a list of balance transactions.
func List(params *stripe.BalanceTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of balance transactions.
func (c Client) List(listParams *stripe.BalanceTransactionListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.BalanceTransactionList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/balance_transactions", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for balance transactions.
type Iter struct {
	*stripe.Iter
}

// BalanceTransaction returns the balance transaction which the iterator is currently pointing to.
func (i *Iter) BalanceTransaction() *stripe.BalanceTransaction {
	return i.Current().(*stripe.BalanceTransaction)
}

// BalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) BalanceTransactionList() *stripe.BalanceTransactionList {
	return i.List().(*stripe.BalanceTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
