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

// v1TreasuryTransactionService is used to invoke /v1/treasury/transactions APIs.
type v1TreasuryTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing Transaction.
func (c v1TreasuryTransactionService) Retrieve(ctx context.Context, id string, params *TreasuryTransactionRetrieveParams) (*TreasuryTransaction, error) {
	if params == nil {
		params = &TreasuryTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/transactions/%s", id)
	transaction := &TreasuryTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Retrieves a list of Transaction objects.
func (c v1TreasuryTransactionService) List(ctx context.Context, listParams *TreasuryTransactionListParams) Seq2[*TreasuryTransaction, error] {
	if listParams == nil {
		listParams = &TreasuryTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryTransaction, ListContainer, error) {
		list := &TreasuryTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
