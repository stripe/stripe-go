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

// v2MoneyManagementSpendingMandateService is used to invoke spendingmandate related APIs.
type v2MoneyManagementSpendingMandateService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing SpendingMandate.
func (c v2MoneyManagementSpendingMandateService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementSpendingMandateRetrieveParams) (*V2MoneyManagementSpendingMandate, error) {
	if params == nil {
		params = &V2MoneyManagementSpendingMandateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/spending_mandates/%s", id)
	spendingmandate := &V2MoneyManagementSpendingMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, spendingmandate)
	return spendingmandate, err
}

// Cancels an active SpendingMandate.
func (c v2MoneyManagementSpendingMandateService) Cancel(ctx context.Context, id string, params *V2MoneyManagementSpendingMandateCancelParams) (*V2MoneyManagementSpendingMandate, error) {
	if params == nil {
		params = &V2MoneyManagementSpendingMandateCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/spending_mandates/%s/cancel", id)
	spendingmandate := &V2MoneyManagementSpendingMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, spendingmandate)
	return spendingmandate, err
}

// Returns a list of SpendingMandates.
func (c v2MoneyManagementSpendingMandateService) List(ctx context.Context, listParams *V2MoneyManagementSpendingMandateListParams) *V2List[*V2MoneyManagementSpendingMandate] {
	if listParams == nil {
		listParams = &V2MoneyManagementSpendingMandateListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/spending_mandates", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementSpendingMandate], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementSpendingMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
