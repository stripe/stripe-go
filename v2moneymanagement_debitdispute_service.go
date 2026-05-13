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

// v2MoneyManagementDebitDisputeService is used to invoke debitdispute related APIs.
type v2MoneyManagementDebitDisputeService struct {
	B   Backend
	Key string
}

// Creates a new DebitDispute for a ReceivedDebit.
func (c v2MoneyManagementDebitDisputeService) Create(ctx context.Context, params *V2MoneyManagementDebitDisputeCreateParams) (*V2MoneyManagementDebitDispute, error) {
	if params == nil {
		params = &V2MoneyManagementDebitDisputeCreateParams{}
	}
	params.Context = ctx
	debitdispute := &V2MoneyManagementDebitDispute{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/debit_disputes", c.Key, params, debitdispute)
	return debitdispute, err
}

// Retrieves the details of an existing DebitDispute.
func (c v2MoneyManagementDebitDisputeService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementDebitDisputeRetrieveParams) (*V2MoneyManagementDebitDispute, error) {
	if params == nil {
		params = &V2MoneyManagementDebitDisputeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/debit_disputes/%s", id)
	debitdispute := &V2MoneyManagementDebitDispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, debitdispute)
	return debitdispute, err
}

// Retrieves a list of DebitDisputes.
func (c v2MoneyManagementDebitDisputeService) List(ctx context.Context, listParams *V2MoneyManagementDebitDisputeListParams) *V2List[*V2MoneyManagementDebitDispute] {
	if listParams == nil {
		listParams = &V2MoneyManagementDebitDisputeListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/debit_disputes", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementDebitDispute], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementDebitDispute]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
