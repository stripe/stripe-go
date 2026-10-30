//
//
// File generated from our OpenAPI spec
//
//

// Package fundingintent provides the fundingintent related APIs
package fundingintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke fundingintent related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a FundingIntent that adds funds to a FinancialAccount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementFundingIntentParams) (*stripe.V2MoneyManagementFundingIntent, error) {
	fundingintent := &stripe.V2MoneyManagementFundingIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/funding_intents", c.Key, params, fundingintent)
	return fundingintent, err
}

// Retrieves a FundingIntent by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementFundingIntentParams) (*stripe.V2MoneyManagementFundingIntent, error) {
	path := stripe.FormatURLPath("/v2/money_management/funding_intents/%s", id)
	fundingintent := &stripe.V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Updates the mutable details of a FundingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2MoneyManagementFundingIntentParams) (*stripe.V2MoneyManagementFundingIntent, error) {
	path := stripe.FormatURLPath("/v2/money_management/funding_intents/%s", id)
	fundingintent := &stripe.V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Cancels a FundingIntent that can still be canceled.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementFundingIntentCancelParams) (*stripe.V2MoneyManagementFundingIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/funding_intents/%s/cancel", id)
	fundingintent := &stripe.V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Confirms a FundingIntent that is ready for processing.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Confirm(id string, params *stripe.V2MoneyManagementFundingIntentConfirmParams) (*stripe.V2MoneyManagementFundingIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/funding_intents/%s/confirm", id)
	fundingintent := &stripe.V2MoneyManagementFundingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, fundingintent)
	return fundingintent, err
}

// Lists FundingIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementFundingIntentListParams) stripe.Seq2[*stripe.V2MoneyManagementFundingIntent, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementFundingIntentListParams{}
	}
	return stripe.NewV2List("/v2/money_management/funding_intents", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementFundingIntent], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementFundingIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
