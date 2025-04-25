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

// v2MoneyManagementReceivedDebitService is used to invoke receiveddebit related APIs.
type v2MoneyManagementReceivedDebitService struct {
	B   Backend
	Key string
}

// Retrieves a single ReceivedDebit by ID.
func (c v2MoneyManagementReceivedDebitService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementReceivedDebitRetrieveParams) (*V2MoneyManagementReceivedDebit, error) {
	path := FormatURLPath("/v2/money_management/received_debits/%s", id)
	receiveddebit := &V2MoneyManagementReceivedDebit{}
	if params == nil {
		params = &V2MoneyManagementReceivedDebitRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebit)
	return receiveddebit, err
}

// Retrieves a list of ReceivedDebits, given the selected filters.
func (c v2MoneyManagementReceivedDebitService) List(ctx context.Context, listParams *V2MoneyManagementReceivedDebitListParams) Seq2[*V2MoneyManagementReceivedDebit, error] {
	return NewV2List("/v2/money_management/received_debits", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementReceivedDebit], error) {
		page := &V2Page[*V2MoneyManagementReceivedDebit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
