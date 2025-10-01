//
//
// File generated from our OpenAPI spec
//
//

// Package action provides the action related APIs
package action

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke action related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a Billing Intent Action.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingIntentsActionParams) (*stripe.V2BillingIntentAction, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/intents/%s/actions/%s", stripe.StringValue(
			params.IntentID), id)
	intentaction := &stripe.V2BillingIntentAction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, intentaction)
	return intentaction, err
}

// List Billing Intent Actions.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingIntentsActionListParams) stripe.Seq2[*stripe.V2BillingIntentAction, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/intents/%s/actions", stripe.StringValue(listParams.IntentID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingIntentAction], error) {
		page := &stripe.V2Page[*stripe.V2BillingIntentAction]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
