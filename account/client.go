//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the /accounts APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new account.
func New(params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().New(params)
}

// New creates a new account.
func (c Client) New(params *stripe.AccountParams) (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, "/v1/accounts", c.Key, params, account)
	return account, err
}

// Get retrieves the authenticating account.
func Get() (*stripe.Account, error) {
	return getC().Get()
}

// Get retrieves the authenticating account.
func (c Client) Get() (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, "/v1/account", c.Key, nil, account)
	return account, err
}

// GetByID returns the details of an account.
func GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().GetByID(id, params)
}

// GetByID returns the details of an account.
func (c Client) GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Update updates an account's properties.
func Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Update(id, params)
}

// Update updates an account's properties.
func (c Client) Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Del removes an account.
func Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Del(id, params)
}

// Del removes an account.
func (c Client) Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, account)
	return account, err
}

// Reject is the method for the `POST /v1/accounts/{account}/reject` API.
func Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	return getC().Reject(id, params)
}

// Reject is the method for the `POST /v1/accounts/{account}/reject` API.
func (c Client) Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s/reject", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// List returns a list of accounts.
func List(params *stripe.AccountListParams) *Iter {
	return getC().List(params)
}

// List returns a list of accounts.
func (c Client) List(listParams *stripe.AccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/accounts", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for accounts.
type Iter struct {
	*stripe.Iter
}

// Account returns the account which the iterator is currently pointing to.
func (i *Iter) Account() *stripe.Account {
	return i.Current().(*stripe.Account)
}

// AccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AccountList() *stripe.AccountList {
	return i.List().(*stripe.AccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
