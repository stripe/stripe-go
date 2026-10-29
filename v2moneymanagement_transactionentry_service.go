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

// v2MoneyManagementTransactionEntryService is used to invoke transactionentry related APIs.
type v2MoneyManagementTransactionEntryService struct {
	B   Backend
	Key string
}

// Retrieves the details of a TransactionEntry by ID.
func (c v2MoneyManagementTransactionEntryService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementTransactionEntryRetrieveParams) (*V2MoneyManagementTransactionEntry, error) {
	if params == nil {
		params = &V2MoneyManagementTransactionEntryRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/transaction_entries/%s", id)
	transactionentry := &V2MoneyManagementTransactionEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Returns a list of TransactionEntries that match the provided filters.
func (c v2MoneyManagementTransactionEntryService) List(ctx context.Context, listParams *V2MoneyManagementTransactionEntryListParams) *V2List[*V2MoneyManagementTransactionEntry] {
	if listParams == nil {
		listParams = &V2MoneyManagementTransactionEntryListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/transaction_entries", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementTransactionEntry], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementTransactionEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
