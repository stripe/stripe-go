// Package balance provides the /balance APIs
package balance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /balance and transaction-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of your balance.
// For more details see https://stripe.com/docs/api#retrieve_balance.
func Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	return getC().Get(params)
}

func (c Client) Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	balance := &stripe.Balance{}
	err := c.B.Call(http.MethodGet, "/balance", c.Key, params, balance)
	return balance, err
}

// GetBalanceTransaction returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func GetBalanceTransaction(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	return getC().GetBalanceTransaction(id, params)
}

func (c Client) GetBalanceTransaction(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	path := stripe.FormatURLPath("/balance/history/%s", id)
	balance := &stripe.BalanceTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, balance)
	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func List(params *stripe.BalanceTransactionListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.BalanceTransactionListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.BalanceTransactionList{}
		err := c.B.CallRaw(http.MethodGet, "/balance/history", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BalanceTransactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Charge returns the most recent BalanceTransaction
// visited by a call to Next.
func (i *Iter) BalanceTransaction() *stripe.BalanceTransaction {
	return i.Current().(*stripe.BalanceTransaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
