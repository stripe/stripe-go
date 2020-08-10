// Package customerbalancetransaction provides the /balance_transactions APIs
package customerbalancetransaction

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /balance_transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new customer balance transaction.
func New(params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	return getC().New(params)
}

// New creates a new customer balance transaction.
func (c Client) New(params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Customer must be set")
	}

	path := stripe.FormatURLPath("/v1/customers/%s/balance_transactions", stripe.StringValue(params.Customer))
	transaction := &stripe.CustomerBalanceTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Get returns the details of a customer balance transaction.
func Get(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a customer balance transaction.
func (c Client) Get(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.Customer must be set")
	}

	path := stripe.FormatURLPath("/v1/customers/%s/balance_transactions/%s", stripe.StringValue(params.Customer), id)
	transaction := &stripe.CustomerBalanceTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// List returns a list of customer balance transactions.
func List(params *stripe.CustomerBalanceTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of customer balance transactions.
func (c Client) List(listParams *stripe.CustomerBalanceTransactionListParams) *Iter {
	path := stripe.FormatURLPath("/v1/customers/%s/balance_transactions", stripe.StringValue(listParams.Customer))

	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.CustomerBalanceTransactionList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Update updates a customer balance transaction.
func Update(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	return getC().Update(id, params)
}

// Update updates a customer balance transaction.
func (c Client) Update(id string, params *stripe.CustomerBalanceTransactionParams) (*stripe.CustomerBalanceTransaction, error) {
	path := stripe.FormatURLPath("/v1/customers/%s/balance_transactions/%s", stripe.StringValue(params.Customer), id)
	transaction := &stripe.CustomerBalanceTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

// Iter is an iterator for customer balance transactions.
type Iter struct {
	*stripe.Iter
}

// CustomerBalanceTransaction returns the customer balance transaction which the iterator is
// currently pointing to.
func (i *Iter) CustomerBalanceTransaction() *stripe.CustomerBalanceTransaction {
	return i.Current().(*stripe.CustomerBalanceTransaction)
}

// CustomerBalanceTransactionList returns the current list object which the
// iterator is currently using. List objects will change as new API calls are
// made to continue pagination.
func (i *Iter) CustomerBalanceTransactionList() *stripe.CustomerBalanceTransactionList {
	return i.List().(*stripe.CustomerBalanceTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
