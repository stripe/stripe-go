//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the /financial_connections/accounts APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /financial_connections/accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// GetByID returns the details of a financial connections account.
func GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().GetByID(id, params)
}

// GetByID returns the details of a financial connections account.
func (c Client) GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath("/v1/financial_connections/accounts/%s", id)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Disconnect is the method for the `POST /v1/financial_connections/accounts/{account}/disconnect` API.
func Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().Disconnect(id, params)
}

// Disconnect is the method for the `POST /v1/financial_connections/accounts/{account}/disconnect` API.
func (c Client) Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/disconnect",
		id,
	)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Refresh is the method for the `POST /v1/financial_connections/accounts/{account}/refresh` API.
func Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().Refresh(id, params)
}

// Refresh is the method for the `POST /v1/financial_connections/accounts/{account}/refresh` API.
func (c Client) Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/refresh",
		id,
	)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// List returns a list of financial connections accounts.
func List(params *stripe.FinancialConnectionsAccountListParams) *Iter {
	return getC().List(params)
}

// List returns a list of financial connections accounts.
func (c Client) List(listParams *stripe.FinancialConnectionsAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialConnectionsAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/financial_connections/accounts", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for financial connections accounts.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsAccount returns the financial connections account which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsAccount() *stripe.FinancialConnectionsAccount {
	return i.Current().(*stripe.FinancialConnectionsAccount)
}

// FinancialConnectionsAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsAccountList() *stripe.FinancialConnectionsAccountList {
	return i.List().(*stripe.FinancialConnectionsAccountList)
}

// ListOwners is the method for the `GET /v1/financial_connections/accounts/{account}/owners` API.
func ListOwners(params *stripe.FinancialConnectionsAccountListOwnersParams) *OwnerIter {
	return getC().ListOwners(params)
}

// ListOwners is the method for the `GET /v1/financial_connections/accounts/{account}/owners` API.
func (c Client) ListOwners(listParams *stripe.FinancialConnectionsAccountListOwnersParams) *OwnerIter {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/owners",
		stripe.StringValue(listParams.Account),
	)
	return &OwnerIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialConnectionsAccountOwnerList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// OwnerIter is an iterator for financial connections account owners.
type OwnerIter struct {
	*stripe.Iter
}

// FinancialConnectionsAccountOwner returns the financial connections account owner which the iterator is currently pointing to.
func (i *OwnerIter) FinancialConnectionsAccountOwner() *stripe.FinancialConnectionsAccountOwner {
	return i.Current().(*stripe.FinancialConnectionsAccountOwner)
}

// FinancialConnectionsAccountOwnerList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *OwnerIter) FinancialConnectionsAccountOwnerList() *stripe.FinancialConnectionsAccountOwnerList {
	return i.List().(*stripe.FinancialConnectionsAccountOwnerList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
