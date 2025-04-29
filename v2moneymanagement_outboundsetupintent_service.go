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

// v2MoneyManagementOutboundSetupIntentService is used to invoke outboundsetupintent related APIs.
type v2MoneyManagementOutboundSetupIntentService struct {
	B   Backend
	Key string
}

// Create an OutboundSetupIntent object.
func (c v2MoneyManagementOutboundSetupIntentService) Create(ctx context.Context, params *V2MoneyManagementOutboundSetupIntentCreateParams) (*V2MoneyManagementOutboundSetupIntent, error) {
	outboundsetupintent := &V2MoneyManagementOutboundSetupIntent{}
	if params == nil {
		params = &V2MoneyManagementOutboundSetupIntentCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_setup_intents", c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Retrieve an OutboundSetupIntent object.
func (c v2MoneyManagementOutboundSetupIntentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementOutboundSetupIntentRetrieveParams) (*V2MoneyManagementOutboundSetupIntent, error) {
	path := FormatURLPath("/v2/money_management/outbound_setup_intents/%s", id)
	outboundsetupintent := &V2MoneyManagementOutboundSetupIntent{}
	if params == nil {
		params = &V2MoneyManagementOutboundSetupIntentRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Update an OutboundSetupIntent object.
func (c v2MoneyManagementOutboundSetupIntentService) Update(ctx context.Context, id string, params *V2MoneyManagementOutboundSetupIntentUpdateParams) (*V2MoneyManagementOutboundSetupIntent, error) {
	path := FormatURLPath("/v2/money_management/outbound_setup_intents/%s", id)
	outboundsetupintent := &V2MoneyManagementOutboundSetupIntent{}
	if params == nil {
		params = &V2MoneyManagementOutboundSetupIntentUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Cancel an OutboundSetupIntent object.
func (c v2MoneyManagementOutboundSetupIntentService) Cancel(ctx context.Context, id string, params *V2MoneyManagementOutboundSetupIntentCancelParams) (*V2MoneyManagementOutboundSetupIntent, error) {
	path := FormatURLPath(
		"/v2/money_management/outbound_setup_intents/%s/cancel", id)
	outboundsetupintent := &V2MoneyManagementOutboundSetupIntent{}
	if params == nil {
		params = &V2MoneyManagementOutboundSetupIntentCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// List the OutboundSetupIntent objects.
func (c v2MoneyManagementOutboundSetupIntentService) List(ctx context.Context, listParams *V2MoneyManagementOutboundSetupIntentListParams) Seq2[*V2MoneyManagementOutboundSetupIntent, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementOutboundSetupIntentListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/outbound_setup_intents", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementOutboundSetupIntent], error) {
		page := &V2Page[*V2MoneyManagementOutboundSetupIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
