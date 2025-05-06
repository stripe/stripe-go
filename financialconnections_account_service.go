//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1FinancialConnectionsAccountService is used to invoke /v1/financial_connections/accounts APIs.
type v1FinancialConnectionsAccountService struct {
	B   Backend
	Key string
}

// Retrieves the details of an Financial Connections Account.
func (c v1FinancialConnectionsAccountService) GetByID(ctx context.Context, id string, params *FinancialConnectionsAccountRetrieveParams) (*FinancialConnectionsAccount, error) {
	if params == nil {
		params = &FinancialConnectionsAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/accounts/%s", id)
	account := &FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Disables your access to a Financial Connections Account. You will no longer be able to access data associated with the account (e.g. balances, transactions).
func (c v1FinancialConnectionsAccountService) Disconnect(ctx context.Context, id string, params *FinancialConnectionsAccountDisconnectParams) (*FinancialConnectionsAccount, error) {
	if params == nil {
		params = &FinancialConnectionsAccountDisconnectParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/accounts/%s/disconnect", id)
	account := &FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Refreshes the data associated with a Financial Connections Account.
func (c v1FinancialConnectionsAccountService) Refresh(ctx context.Context, id string, params *FinancialConnectionsAccountRefreshParams) (*FinancialConnectionsAccount, error) {
	if params == nil {
		params = &FinancialConnectionsAccountRefreshParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/accounts/%s/refresh", id)
	account := &FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Subscribes to periodic refreshes of data associated with a Financial Connections Account.
func (c v1FinancialConnectionsAccountService) Subscribe(ctx context.Context, id string, params *FinancialConnectionsAccountSubscribeParams) (*FinancialConnectionsAccount, error) {
	if params == nil {
		params = &FinancialConnectionsAccountSubscribeParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/accounts/%s/subscribe", id)
	account := &FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Unsubscribes from periodic refreshes of data associated with a Financial Connections Account.
func (c v1FinancialConnectionsAccountService) Unsubscribe(ctx context.Context, id string, params *FinancialConnectionsAccountUnsubscribeParams) (*FinancialConnectionsAccount, error) {
	if params == nil {
		params = &FinancialConnectionsAccountUnsubscribeParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/accounts/%s/unsubscribe", id)
	account := &FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Returns a list of Financial Connections Account objects.
func (c v1FinancialConnectionsAccountService) List(ctx context.Context, listParams *FinancialConnectionsAccountListParams) Seq2[*FinancialConnectionsAccount, error] {
	if listParams == nil {
		listParams = &FinancialConnectionsAccountListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*FinancialConnectionsAccount, ListContainer, error) {
		list := &FinancialConnectionsAccountList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/financial_connections/accounts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// Lists all owners for a given Account
func (c v1FinancialConnectionsAccountService) ListOwners(ctx context.Context, listParams *FinancialConnectionsAccountListOwnersParams) Seq2[*FinancialConnectionsAccountOwner, error] {
	if listParams == nil {
		listParams = &FinancialConnectionsAccountListOwnersParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/financial_connections/accounts/%s/owners", StringValue(
			listParams.Account))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*FinancialConnectionsAccountOwner, ListContainer, error) {
		list := &FinancialConnectionsAccountOwnerList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
