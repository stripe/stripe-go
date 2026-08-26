//
//
// File generated from our OpenAPI spec
//
//

// Package accountactivity provides the accountactivity related APIs
package accountactivity

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke accountactivity related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new account activity to report account registration, login, or evaluation follow-up activity.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2SignalsAccountActivityParams) (*stripe.V2SignalsAccountActivity, error) {
	accountactivity := &stripe.V2SignalsAccountActivity{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/account_activity", c.Key, params, accountactivity)
	return accountactivity, err
}

// Retrieves an AccountActivity by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2SignalsAccountActivityParams) (*stripe.V2SignalsAccountActivity, error) {
	path := stripe.FormatURLPath("/v2/signals/account_activity/%s", id)
	accountactivity := &stripe.V2SignalsAccountActivity{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountactivity)
	return accountactivity, err
}

// Deletes an AccountActivity by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2SignalsAccountActivityParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath("/v2/signals/account_activity/%s", id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}
