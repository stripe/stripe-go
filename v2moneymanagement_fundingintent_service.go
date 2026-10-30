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

// v2MoneyManagementFundingIntentService is used to invoke fundingintent related APIs.
type v2MoneyManagementFundingIntentService struct {
	B   Backend
	Key string
}

// Creates a FundingIntent that adds funds to a FinancialAccount.
func (c v2MoneyManagementFundingIntentService) Create(ctx context.Context, params *V2MoneyManagementFundingIntentCreateParams) (*V2MoneyManagementFundingIntent, error) {
	if params == nil {
		params = &V2MoneyManagementFundingIntentCreateParams{}
	}
	params.Context = ctx
	fundingintent := &V2MoneyManagementFundingIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/funding_intents", c.Key, params, fundingintent)
	return fundingintent, err
}

// Retrieves a FundingIntent by ID.
func (c v2MoneyManagementFundingIntentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementFundingIntentRetrieveParams) (*V2MoneyManagementFundingIntent, error) {
	if params == nil {
		params = &V2MoneyManagementFundingIntentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/funding_intents/%s", id)
	fundingintent := &V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Updates the mutable details of a FundingIntent.
func (c v2MoneyManagementFundingIntentService) Update(ctx context.Context, id string, params *V2MoneyManagementFundingIntentUpdateParams) (*V2MoneyManagementFundingIntent, error) {
	if params == nil {
		params = &V2MoneyManagementFundingIntentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/funding_intents/%s", id)
	fundingintent := &V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Cancels a FundingIntent that can still be canceled.
func (c v2MoneyManagementFundingIntentService) Cancel(ctx context.Context, id string, params *V2MoneyManagementFundingIntentCancelParams) (*V2MoneyManagementFundingIntent, error) {
	if params == nil {
		params = &V2MoneyManagementFundingIntentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/funding_intents/%s/cancel", id)
	fundingintent := &V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Confirms a FundingIntent that is ready for processing.
func (c v2MoneyManagementFundingIntentService) Confirm(ctx context.Context, id string, params *V2MoneyManagementFundingIntentConfirmParams) (*V2MoneyManagementFundingIntent, error) {
	if params == nil {
		params = &V2MoneyManagementFundingIntentConfirmParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/funding_intents/%s/confirm", id)
	fundingintent := &V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Lists FundingIntents.
func (c v2MoneyManagementFundingIntentService) List(ctx context.Context, listParams *V2MoneyManagementFundingIntentListParams) *V2List[*V2MoneyManagementFundingIntent] {
	if listParams == nil {
		listParams = &V2MoneyManagementFundingIntentListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/funding_intents", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementFundingIntent], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementFundingIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
