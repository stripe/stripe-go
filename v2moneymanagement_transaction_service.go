//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2MoneyManagementTransactionService is used to invoke transaction related APIs.
type v2MoneyManagementTransactionService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Transaction by ID.
func (c v2MoneyManagementTransactionService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementTransactionRetrieveParams) (*V2MoneyManagementTransaction, error) {
	path := FormatURLPath("/v2/money_management/transactions/%s", id)
	transaction := &V2MoneyManagementTransaction{}
	if params == nil {
		params = &V2MoneyManagementTransactionRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Returns a list of Transactions that match the provided filters.
func (c v2MoneyManagementTransactionService) List(ctx context.Context, listParams *V2MoneyManagementTransactionListParams) Seq2[*V2MoneyManagementTransaction, error] {
	return NewV2List("/v2/money_management/transactions", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementTransaction], error) {
		page := &V2Page[*V2MoneyManagementTransaction]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
