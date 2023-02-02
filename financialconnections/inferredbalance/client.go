//
//
// File generated from our OpenAPI spec
//
//

// Package inferredbalance provides the /financial_connections/accounts/{account}/inferred_balances APIs
package inferredbalance

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /financial_connections/accounts/{account}/inferred_balances APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// List returns a list of financial connections inferred balances.
func List(params *stripe.FinancialConnectionsInferredBalanceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of financial connections inferred balances.
func (c Client) List(listParams *stripe.FinancialConnectionsInferredBalanceListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/inferred_balances",
		stripe.StringValue(listParams.Account),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialConnectionsInferredBalanceList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for financial connections inferred balances.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsInferredBalance returns the financial connections inferred balance which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsInferredBalance() *stripe.FinancialConnectionsInferredBalance {
	return i.Current().(*stripe.FinancialConnectionsInferredBalance)
}

// FinancialConnectionsInferredBalanceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsInferredBalanceList() *stripe.FinancialConnectionsInferredBalanceList {
	return i.List().(*stripe.FinancialConnectionsInferredBalanceList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
