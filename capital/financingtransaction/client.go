//
//
// File generated from our OpenAPI spec
//
//

// Package financingtransaction provides the /capital/financing_transactions APIs
package financingtransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/form"
)

// Client is used to invoke /capital/financing_transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a financing transaction for a financing offer.
func Get(id string, params *stripe.CapitalFinancingTransactionParams) (*stripe.CapitalFinancingTransaction, error) {
	return getC().Get(id, params)
}

// Retrieves a financing transaction for a financing offer.
func (c Client) Get(id string, params *stripe.CapitalFinancingTransactionParams) (*stripe.CapitalFinancingTransaction, error) {
	path := stripe.FormatURLPath("/v1/capital/financing_transactions/%s", id)
	financingtransaction := &stripe.CapitalFinancingTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingtransaction)
	return financingtransaction, err
}

// Returns a list of financing transactions. The transactions are returned in sorted order,
// with the most recent transactions appearing first.
func List(params *stripe.CapitalFinancingTransactionListParams) *Iter {
	return getC().List(params)
}

// Returns a list of financing transactions. The transactions are returned in sorted order,
// with the most recent transactions appearing first.
func (c Client) List(listParams *stripe.CapitalFinancingTransactionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CapitalFinancingTransactionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_transactions", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for capital financing transactions.
type Iter struct {
	*stripe.Iter
}

// CapitalFinancingTransaction returns the capital financing transaction which the iterator is currently pointing to.
func (i *Iter) CapitalFinancingTransaction() *stripe.CapitalFinancingTransaction {
	return i.Current().(*stripe.CapitalFinancingTransaction)
}

// CapitalFinancingTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CapitalFinancingTransactionList() *stripe.CapitalFinancingTransactionList {
	return i.List().(*stripe.CapitalFinancingTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
