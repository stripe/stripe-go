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

// v1FinancialConnectionsTransactionService is used to invoke /v1/financial_connections/transactions APIs.
type v1FinancialConnectionsTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Financial Connections Transaction
func (c v1FinancialConnectionsTransactionService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsTransactionRetrieveParams) (*FinancialConnectionsTransaction, error) {
	if params == nil {
		params = &FinancialConnectionsTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/transactions/%s", id)
	transaction := &FinancialConnectionsTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Returns a list of Financial Connections Transaction objects.
func (c v1FinancialConnectionsTransactionService) List(ctx context.Context, listParams *FinancialConnectionsTransactionListParams) Seq2[*FinancialConnectionsTransaction, error] {
	if listParams == nil {
		listParams = &FinancialConnectionsTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*FinancialConnectionsTransaction, ListContainer, error) {
		list := &FinancialConnectionsTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/financial_connections/transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
