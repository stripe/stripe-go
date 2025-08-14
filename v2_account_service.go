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

// v2AccountService is used to invoke account related APIs.
type v2AccountService struct {
	B   Backend
	Key string
}

// Creates an Account. You can optionally provide information about the associated Legal Entity, such as name and business type. The Account can also be configured as a recipient of OutboundPayments by requesting Features on the recipient configuration.
func (c v2AccountService) Create(ctx context.Context, params *V2AccountCreateParams) (*V2Account, error) {
	if params == nil {
		params = &V2AccountCreateParams{}
	}
	params.Context = ctx
	account := &V2Account{}
	err := c.B.Call(http.MethodPost, "/v2/accounts", c.Key, params, account)
	return account, err
}

// Retrieves the details of an Account.
func (c v2AccountService) Retrieve(ctx context.Context, id string, params *V2AccountRetrieveParams) (*V2Account, error) {
	if params == nil {
		params = &V2AccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/accounts/%s", id)
	account := &V2Account{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Updates the details of an Account. You can also optionally provide or update the details of the associated Legal Entity and recipient configuration.
func (c v2AccountService) Update(ctx context.Context, id string, params *V2AccountUpdateParams) (*V2Account, error) {
	if params == nil {
		params = &V2AccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/accounts/%s", id)
	account := &V2Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Closes and removes access to the Account and its associated resources.
func (c v2AccountService) Close(ctx context.Context, id string, params *V2AccountCloseParams) (*V2Account, error) {
	if params == nil {
		params = &V2AccountCloseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/accounts/%s/close", id)
	account := &V2Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of Accounts. Note that the `include` parameter cannot be passed in on list requests.
func (c v2AccountService) List(ctx context.Context, listParams *V2AccountListParams) Seq2[*V2Account, error] {
	if listParams == nil {
		listParams = &V2AccountListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/accounts", listParams, func(path string, p ParamsContainer) (*V2Page[*V2Account], error) {
		page := &V2Page[*V2Account]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
