//
//
// File generated from our OpenAPI spec
//
//

// Package accountinferredbalance provides the /financial_connections/accounts/{account}/inferred_balances APIs
package accountinferredbalance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /financial_connections/accounts/{account}/inferred_balances APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of financial connections account inferred balances.
func List(params *stripe.FinancialConnectionsAccountInferredBalanceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of financial connections account inferred balances.
func (c Client) List(listParams *stripe.FinancialConnectionsAccountInferredBalanceListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/inferred_balances",
		stripe.StringValue(listParams.Account),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialConnectionsAccountInferredBalanceList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for financial connections account inferred balances.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsAccountInferredBalance returns the financial connections account inferred balance which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsAccountInferredBalance() *stripe.FinancialConnectionsAccountInferredBalance {
	return i.Current().(*stripe.FinancialConnectionsAccountInferredBalance)
}

// FinancialConnectionsAccountInferredBalanceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsAccountInferredBalanceList() *stripe.FinancialConnectionsAccountInferredBalanceList {
	return i.List().(*stripe.FinancialConnectionsAccountInferredBalanceList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
