//
//
// File generated from our OpenAPI spec
//
//

// Package externalaccount provides the /external_accounts APIs
package externalaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/form"
)

// Client is used to invoke /external_accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an external account for a given connected account.
func New(params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	return getC().New(params)
}

// Create an external account for a given connected account.
func (c Client) New(params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	externalaccount := &stripe.ExternalAccount{}
	err := c.B.Call(
		http.MethodPost, "/v1/external_accounts", c.Key, params, externalaccount)
	return externalaccount, err
}

// Retrieve a specified external account for a given account.
func Get(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	return getC().Get(id, params)
}

// Retrieve a specified external account for a given account.
func (c Client) Get(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	externalaccount := &stripe.ExternalAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, externalaccount)
	return externalaccount, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func Update(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	return getC().Update(id, params)
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func (c Client) Update(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	externalaccount := &stripe.ExternalAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, externalaccount)
	return externalaccount, err
}

// Delete a specified external account for a given account.
func Del(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	return getC().Del(id, params)
}

// Delete a specified external account for a given account.
func (c Client) Del(id string, params *stripe.ExternalAccountParams) (*stripe.ExternalAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	externalaccount := &stripe.ExternalAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, externalaccount)
	return externalaccount, err
}

// List external accounts for an account.
func List(params *stripe.ExternalAccountListParams) *Iter {
	return getC().List(params)
}

// List external accounts for an account.
func (c Client) List(listParams *stripe.ExternalAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ExternalAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/external_accounts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for external accounts.
type Iter struct {
	*stripe.Iter
}

// ExternalAccount returns the external account which the iterator is currently pointing to.
func (i *Iter) ExternalAccount() *stripe.ExternalAccount {
	return i.Current().(*stripe.ExternalAccount)
}

// ExternalAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ExternalAccountList() *stripe.ExternalAccountList {
	return i.List().(*stripe.ExternalAccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
