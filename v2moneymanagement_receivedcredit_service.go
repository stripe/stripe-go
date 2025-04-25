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
	path := FormatURLPath("/v2/money_management/received_credits/%s", id)
	receivedcredit := &V2MoneyManagementReceivedCredit{}
	if params == nil {
		params = &V2MoneyManagementReceivedCreditRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Retrieves a list of ReceivedCredits.
func (c v2MoneyManagementReceivedCreditService) List(ctx context.Context, listParams *V2MoneyManagementReceivedCreditListParams) Seq2[*V2MoneyManagementReceivedCredit, error] {
	return NewV2List("/v2/money_management/received_credits", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementReceivedCredit], error) {
		page := &V2Page[*V2MoneyManagementReceivedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
