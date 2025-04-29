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

// v1CapitalFinancingTransactionService is used to invoke /v1/capital/financing_transactions APIs.
type v1CapitalFinancingTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a financing transaction for a financing offer.
func (c v1CapitalFinancingTransactionService) Retrieve(ctx context.Context, id string, params *CapitalFinancingTransactionRetrieveParams) (*CapitalFinancingTransaction, error) {
	path := FormatURLPath("/v1/capital/financing_transactions/%s", id)
	financingtransaction := &CapitalFinancingTransaction{}
	if params == nil {
		params = &CapitalFinancingTransactionRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingtransaction)
	return financingtransaction, err
}

// Returns a list of financing transactions. The transactions are returned in sorted order,
// with the most recent transactions appearing first.
func (c v1CapitalFinancingTransactionService) List(ctx context.Context, listParams *CapitalFinancingTransactionListParams) Seq2[*CapitalFinancingTransaction, error] {
	if listParams == nil {
		listParams = &CapitalFinancingTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CapitalFinancingTransaction, ListContainer, error) {
		list := &CapitalFinancingTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
