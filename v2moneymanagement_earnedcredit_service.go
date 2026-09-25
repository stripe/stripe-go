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

// v2MoneyManagementEarnedCreditService is used to invoke earnedcredit related APIs.
type v2MoneyManagementEarnedCreditService struct {
	B   Backend
	Key string
}

// Retrieves an EarnedCredit.
func (c v2MoneyManagementEarnedCreditService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementEarnedCreditRetrieveParams) (*V2MoneyManagementEarnedCredit, error) {
	if params == nil {
		params = &V2MoneyManagementEarnedCreditRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/earned_credits/%s", id)
	earnedcredit := &V2MoneyManagementEarnedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, earnedcredit)
	return earnedcredit, err
}

// Returns a list of EarnedCredits.
func (c v2MoneyManagementEarnedCreditService) List(ctx context.Context, listParams *V2MoneyManagementEarnedCreditListParams) *V2List[*V2MoneyManagementEarnedCredit] {
	if listParams == nil {
		listParams = &V2MoneyManagementEarnedCreditListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/earned_credits", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementEarnedCredit], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementEarnedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
