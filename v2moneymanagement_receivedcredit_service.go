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

// v2MoneyManagementReceivedCreditService is used to invoke receivedcredit related APIs.
type v2MoneyManagementReceivedCreditService struct {
	B   Backend
	Key string
}

// Retrieve a ReceivedCredit by ID.
func (c v2MoneyManagementReceivedCreditService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementReceivedCreditRetrieveParams) (*V2MoneyManagementReceivedCredit, error) {
	if params == nil {
		params = &V2MoneyManagementReceivedCreditRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/received_credits/%s", id)
	receivedcredit := &V2MoneyManagementReceivedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Retrieves a list of ReceivedCredits.
func (c v2MoneyManagementReceivedCreditService) List(ctx context.Context, listParams *V2MoneyManagementReceivedCreditListParams) *V2List[*V2MoneyManagementReceivedCredit] {
	if listParams == nil {
		listParams = &V2MoneyManagementReceivedCreditListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/received_credits", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementReceivedCredit], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementReceivedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
