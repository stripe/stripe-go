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

// v2MoneyManagementPayoutMethodService is used to invoke payoutmethod related APIs.
type v2MoneyManagementPayoutMethodService struct {
	B   Backend
	Key string
}

// Retrieve a PayoutMethod object.
func (c v2MoneyManagementPayoutMethodService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementPayoutMethodRetrieveParams) (*V2MoneyManagementPayoutMethod, error) {
	path := FormatURLPath("/v2/money_management/payout_methods/%s", id)
	payoutmethod := &V2MoneyManagementPayoutMethod{}
	if params == nil {
		params = &V2MoneyManagementPayoutMethodRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Archive a PayoutMethod object. Archived objects cannot be used as payout methods
// and will not appear in the payout method list.
func (c v2MoneyManagementPayoutMethodService) Archive(ctx context.Context, id string, params *V2MoneyManagementPayoutMethodArchiveParams) (*V2MoneyManagementPayoutMethod, error) {
	path := FormatURLPath("/v2/money_management/payout_methods/%s/archive", id)
	payoutmethod := &V2MoneyManagementPayoutMethod{}
	if params == nil {
		params = &V2MoneyManagementPayoutMethodArchiveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// Unarchive an PayoutMethod object.
func (c v2MoneyManagementPayoutMethodService) Unarchive(ctx context.Context, id string, params *V2MoneyManagementPayoutMethodUnarchiveParams) (*V2MoneyManagementPayoutMethod, error) {
	path := FormatURLPath("/v2/money_management/payout_methods/%s/unarchive", id)
	payoutmethod := &V2MoneyManagementPayoutMethod{}
	if params == nil {
		params = &V2MoneyManagementPayoutMethodUnarchiveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutmethod)
	return payoutmethod, err
}

// List objects that adhere to the PayoutMethod interface.
func (c v2MoneyManagementPayoutMethodService) List(ctx context.Context, listParams *V2MoneyManagementPayoutMethodListParams) Seq2[*V2MoneyManagementPayoutMethod, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementPayoutMethodListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/payout_methods", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementPayoutMethod], error) {
		page := &V2Page[*V2MoneyManagementPayoutMethod]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
