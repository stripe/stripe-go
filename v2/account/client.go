//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the account related APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke account related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an Account. You can optionally provide information about the associated Legal Entity, such as name and business type. The Account can also be configured as a recipient of OutboundPayments by requesting Features on the recipient configuration.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2AccountParams) (*stripe.V2Account, error) {
	account := &stripe.V2Account{}
	err := c.B.Call(http.MethodPost, "/v2/accounts", c.Key, params, account)
	return account, err
}

// Retrieves the details of an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2AccountParams) (*stripe.V2Account, error) {
	path := stripe.FormatURLPath("/v2/accounts/%s", id)
	account := &stripe.V2Account{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Updates the details of an Account. You can also optionally provide or update the details of the associated Legal Entity and recipient configuration.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2AccountParams) (*stripe.V2Account, error) {
	path := stripe.FormatURLPath("/v2/accounts/%s", id)
	account := &stripe.V2Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Closes and removes access to the Account and its associated resources.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Close(id string, params *stripe.V2AccountCloseParams) (*stripe.V2Account, error) {
	path := stripe.FormatURLPath("/v2/accounts/%s/close", id)
	account := &stripe.V2Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of Accounts. Note that the `include` parameter cannot be passed in on list requests.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2AccountListParams) stripe.Seq2[*stripe.V2Account, error] {
	return stripe.NewV2List("/v2/accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2Account], error) {
		page := &stripe.V2Page[*stripe.V2Account]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
