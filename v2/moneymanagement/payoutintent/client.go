//
//
// File generated from our OpenAPI spec
//
//

// Package payoutintent provides the payoutintent related APIs
package payoutintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke payoutintent related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a PayoutIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementPayoutIntentParams) (*stripe.V2MoneyManagementPayoutIntent, error) {
	payoutintent := &stripe.V2MoneyManagementPayoutIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/payout_intents", c.Key, params, payoutintent)
	return payoutintent, err
}

// Retrieves the details of an existing PayoutIntent by passing the unique PayoutIntent ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementPayoutIntentParams) (*stripe.V2MoneyManagementPayoutIntent, error) {
	path := stripe.FormatURLPath("/v2/money_management/payout_intents/%s", id)
	payoutintent := &stripe.V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Updates a PayoutIntent. Only pending or requires_action PayoutIntents that are editable can be updated.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2MoneyManagementPayoutIntentParams) (*stripe.V2MoneyManagementPayoutIntent, error) {
	path := stripe.FormatURLPath("/v2/money_management/payout_intents/%s", id)
	payoutintent := &stripe.V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Cancels a PayoutIntent. Only pending PayoutIntents or processing PayoutIntents with cancelable OutboundPayment/Transfer can be canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementPayoutIntentCancelParams) (*stripe.V2MoneyManagementPayoutIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/payout_intents/%s/cancel", id)
	payoutintent := &stripe.V2MoneyManagementPayoutIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, payoutintent)
	return payoutintent, err
}

// Returns a list of PayoutIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementPayoutIntentListParams) stripe.Seq2[*stripe.V2MoneyManagementPayoutIntent, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementPayoutIntentListParams{}
	}
	return stripe.NewV2List("/v2/money_management/payout_intents", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementPayoutIntent], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementPayoutIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
