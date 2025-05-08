//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2CoreAccountService is used to invoke account related APIs.
type v2CoreAccountService struct {
	B   Backend
	Key string
}

// An Account is a representation of a company, individual or other entity that a user interacts with. Accounts contain identifying information about the entity, and configurations that store the features an account has access to. An account can be configured as any or all of the following configurations: Customer, Merchant and/or Recipient.
func (c v2CoreAccountService) Create(ctx context.Context, params *V2CoreAccountCreateParams) (*V2CoreAccount, error) {
	if params == nil {
		params = &V2CoreAccountCreateParams{}
	}
	params.Context = ctx
	account := &V2CoreAccount{}
	err := c.B.Call(http.MethodPost, "/v2/core/accounts", c.Key, params, account)
	return account, err
}

// Retrieves the details of an Account.
func (c v2CoreAccountService) Retrieve(ctx context.Context, id string, params *V2CoreAccountRetrieveParams) (*V2CoreAccount, error) {
	if params == nil {
		params = &V2CoreAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/accounts/%s", id)
	account := &V2CoreAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Updates the details of an Account.
func (c v2CoreAccountService) Update(ctx context.Context, id string, params *V2CoreAccountUpdateParams) (*V2CoreAccount, error) {
	if params == nil {
		params = &V2CoreAccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/accounts/%s", id)
	account := &V2CoreAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Removes access to the Account and its associated resources.
func (c v2CoreAccountService) Close(ctx context.Context, id string, params *V2CoreAccountCloseParams) (*V2CoreAccount, error) {
	if params == nil {
		params = &V2CoreAccountCloseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/accounts/%s/close", id)
	account := &V2CoreAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of Accounts.
func (c v2CoreAccountService) List(ctx context.Context, listParams *V2CoreAccountListParams) Seq2[*V2CoreAccount, error] {
	if listParams == nil {
		listParams = &V2CoreAccountListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/core/accounts", listParams, func(path string, p ParamsContainer) (*V2Page[*V2CoreAccount], error) {
		page := &V2Page[*V2CoreAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
