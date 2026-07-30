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

// v2MoneyManagementReceivedDebitMandateService is used to invoke receiveddebitmandate related APIs.
type v2MoneyManagementReceivedDebitMandateService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebitMandate.
func (c v2MoneyManagementReceivedDebitMandateService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementReceivedDebitMandateRetrieveParams) (*V2MoneyManagementReceivedDebitMandate, error) {
	if params == nil {
		params = &V2MoneyManagementReceivedDebitMandateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/received_debit_mandates/%s", id)
	receiveddebitmandate := &V2MoneyManagementReceivedDebitMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebitmandate)
	return receiveddebitmandate, err
}

// Cancels an active ReceivedDebitMandate.
func (c v2MoneyManagementReceivedDebitMandateService) Cancel(ctx context.Context, id string, params *V2MoneyManagementReceivedDebitMandateCancelParams) (*V2MoneyManagementReceivedDebitMandate, error) {
	if params == nil {
		params = &V2MoneyManagementReceivedDebitMandateCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/received_debit_mandates/%s/cancel", id)
	receiveddebitmandate := &V2MoneyManagementReceivedDebitMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, receiveddebitmandate)
	return receiveddebitmandate, err
}

// Returns a list of ReceivedDebitMandates.
func (c v2MoneyManagementReceivedDebitMandateService) List(ctx context.Context, listParams *V2MoneyManagementReceivedDebitMandateListParams) *V2List[*V2MoneyManagementReceivedDebitMandate] {
	if listParams == nil {
		listParams = &V2MoneyManagementReceivedDebitMandateListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/received_debit_mandates", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementReceivedDebitMandate], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementReceivedDebitMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
