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

// v2BillingIntentsActionService is used to invoke action related APIs.
type v2BillingIntentsActionService struct {
	B   Backend
	Key string
}

// Retrieve a Billing Intent Action.
func (c v2BillingIntentsActionService) Retrieve(ctx context.Context, id string, params *V2BillingIntentsActionRetrieveParams) (*V2BillingIntentAction, error) {
	if params == nil {
		params = &V2BillingIntentsActionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/intents/%s/actions/%s", StringValue(params.IntentID), id)
	intentaction := &V2BillingIntentAction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, intentaction)
	return intentaction, err
}

// List Billing Intent Actions.
func (c v2BillingIntentsActionService) List(ctx context.Context, listParams *V2BillingIntentsActionListParams) Seq2[*V2BillingIntentAction, error] {
	if listParams == nil {
		listParams = &V2BillingIntentsActionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/intents/%s/actions", StringValue(listParams.IntentID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingIntentAction], error) {
		page := &V2Page[*V2BillingIntentAction]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
