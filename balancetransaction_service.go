//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
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
func (c v1BalanceTransactionService) List(ctx context.Context, listParams *BalanceTransactionListParams) *V1List[*BalanceTransaction] {
	if listParams == nil {
		listParams = &BalanceTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*BalanceTransaction], error) {
		list := &v1Page[*BalanceTransaction]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/balance_transactions", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
