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
	path := FormatURLPath("/v2/money_management/transaction_entries/%s", id)
	transactionentry := &V2MoneyManagementTransactionEntry{}
	if params == nil {
		params = &V2MoneyManagementTransactionEntryRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Returns a list of TransactionEntries that match the provided filters.
func (c v2MoneyManagementTransactionEntryService) List(ctx context.Context, listParams *V2MoneyManagementTransactionEntryListParams) Seq2[*V2MoneyManagementTransactionEntry, error] {
	return NewV2List("/v2/money_management/transaction_entries", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementTransactionEntry], error) {
		page := &V2Page[*V2MoneyManagementTransactionEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
