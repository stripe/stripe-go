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

// v1BalanceTransactionService is used to invoke /v1/balance_transactions APIs.
type v1BalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the balance transaction with the given ID.
//
// Note that this endpoint previously used the path /v1/balance/history/:id.
func (c v1BalanceTransactionService) Retrieve(ctx context.Context, id string, params *BalanceTransactionRetrieveParams) (*BalanceTransaction, error) {
	if params == nil {
		params = &BalanceTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/balance_transactions/%s", id)
	balancetransaction := &BalanceTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, balancetransaction)
	return balancetransaction, err
}

// Returns a list of transactions that have contributed to the Stripe account balance (e.g., charges, transfers, and so forth). The transactions are returned in sorted order, with the most recent transactions appearing first.
//
// Note that this endpoint was previously called “Balance history” and used the path /v1/balance/history.
func (c v1BalanceTransactionService) List(ctx context.Context, listParams *BalanceTransactionListParams) Seq2[*BalanceTransaction, error] {
	if listParams == nil {
		listParams = &BalanceTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BalanceTransaction, ListContainer, error) {
		list := &BalanceTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/balance_transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
