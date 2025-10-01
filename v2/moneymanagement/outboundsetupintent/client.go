//
//
// File generated from our OpenAPI spec
//
//

// Package outboundsetupintent provides the outboundsetupintent related APIs
package outboundsetupintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke outboundsetupintent related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an OutboundSetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementOutboundSetupIntentParams) (*stripe.V2MoneyManagementOutboundSetupIntent, error) {
	outboundsetupintent := &stripe.V2MoneyManagementOutboundSetupIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_setup_intents", c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Retrieve an OutboundSetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementOutboundSetupIntentParams) (*stripe.V2MoneyManagementOutboundSetupIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_setup_intents/%s", id)
	outboundsetupintent := &stripe.V2MoneyManagementOutboundSetupIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Update an OutboundSetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2MoneyManagementOutboundSetupIntentParams) (*stripe.V2MoneyManagementOutboundSetupIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_setup_intents/%s", id)
	outboundsetupintent := &stripe.V2MoneyManagementOutboundSetupIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// Cancel an OutboundSetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementOutboundSetupIntentCancelParams) (*stripe.V2MoneyManagementOutboundSetupIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_setup_intents/%s/cancel", id)
	outboundsetupintent := &stripe.V2MoneyManagementOutboundSetupIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundsetupintent)
	return outboundsetupintent, err
}

// List the OutboundSetupIntent objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementOutboundSetupIntentListParams) stripe.Seq2[*stripe.V2MoneyManagementOutboundSetupIntent, error] {
	return stripe.NewV2List("/v2/money_management/outbound_setup_intents", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementOutboundSetupIntent], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementOutboundSetupIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
