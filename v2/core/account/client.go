//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the account related APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke account related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// An Account is a representation of a company, individual or other entity that a user interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
func (c Client) New(params *stripe.V2CoreAccountParams) (*stripe.V2CoreAccount, error) {
	account := &stripe.V2CoreAccount{}
	err := c.B.Call(http.MethodPost, "/v2/core/accounts", c.Key, params, account)
	return account, err
}

// Retrieves the details of an Account.
func (c Client) Get(id string, params *stripe.V2CoreAccountParams) (*stripe.V2CoreAccount, error) {
	path := stripe.FormatURLPath("/v2/core/accounts/%s", id)
	account := &stripe.V2CoreAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Updates the details of an Account.
func (c Client) Update(id string, params *stripe.V2CoreAccountParams) (*stripe.V2CoreAccount, error) {
	path := stripe.FormatURLPath("/v2/core/accounts/%s", id)
	account := &stripe.V2CoreAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Removes access to the Account and its associated resources.
func (c Client) Close(id string, params *stripe.V2CoreAccountCloseParams) (*stripe.V2CoreAccount, error) {
	path := stripe.FormatURLPath("/v2/core/accounts/%s/close", id)
	account := &stripe.V2CoreAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of Accounts.
func (c Client) All(listParams *stripe.V2CoreAccountListParams) stripe.Seq2[stripe.V2CoreAccount, error] {
	return stripe.NewV2List("/v2/core/accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2CoreAccount], error) {
		page := &stripe.V2Page[stripe.V2CoreAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
