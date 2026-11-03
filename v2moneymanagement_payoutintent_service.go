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

// v2MoneyManagementPayoutIntentService is used to invoke payoutintent related APIs.
type v2MoneyManagementPayoutIntentService struct {
	B   Backend
	Key string
}

// Creates a PayoutIntent.
func (c v2MoneyManagementPayoutIntentService) Create(ctx context.Context, params *V2MoneyManagementPayoutIntentCreateParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentCreateParams{}
	}
	params.Context = ctx
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/payout_intents", c.Key, params, payoutintent)
	return payoutintent, err
}

// Retrieves the details of an existing PayoutIntent by passing the unique PayoutIntent ID.
func (c v2MoneyManagementPayoutIntentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentRetrieveParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/payout_intents/%s", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Updates an editable PayoutIntent in the `pending` or `requires_action` state.
func (c v2MoneyManagementPayoutIntentService) Update(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentUpdateParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/payout_intents/%s", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Cancels a PayoutIntent in the `pending` or `requires_action` state, or a `processing` PayoutIntent with a cancelable OutboundPayment.
func (c v2MoneyManagementPayoutIntentService) Cancel(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentCancelParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/payout_intents/%s/cancel", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Confirms a PayoutIntent in the `requires_action` state, transitioning it to `pending`.
func (c v2MoneyManagementPayoutIntentService) Confirm(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentConfirmParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentConfirmParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/payout_intents/%s/confirm", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Refreshes the FX quote for a PayoutIntent.
func (c v2MoneyManagementPayoutIntentService) RefreshFxQuote(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentRefreshFxQuoteParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentRefreshFxQuoteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/payout_intents/%s/refresh_fx_quote", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Refreshes the recipient verification for a PayoutIntent.
func (c v2MoneyManagementPayoutIntentService) RefreshRecipientVerification(ctx context.Context, id string, params *V2MoneyManagementPayoutIntentRefreshRecipientVerificationParams) (*V2MoneyManagementPayoutIntent, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutIntentRefreshRecipientVerificationParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/payout_intents/%s/refresh_recipient_verification", id)
	payoutintent := &V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Returns a list of PayoutIntents.
func (c v2MoneyManagementPayoutIntentService) List(ctx context.Context, listParams *V2MoneyManagementPayoutIntentListParams) *V2List[*V2MoneyManagementPayoutIntent] {
	if listParams == nil {
		listParams = &V2MoneyManagementPayoutIntentListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/payout_intents", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementPayoutIntent], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementPayoutIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
