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

// v2CoreVaultGBBankAccountService is used to invoke gbbankaccount related APIs.
type v2CoreVaultGBBankAccountService struct {
	B   Backend
	Key string
}

// Create a GB bank account.
func (c v2CoreVaultGBBankAccountService) Create(ctx context.Context, params *V2CoreVaultGBBankAccountCreateParams) (*V2CoreVaultGBBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultGBBankAccountCreateParams{}
	}
	params.Context = ctx
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/gb_bank_accounts", c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Retrieve a GB bank account.
func (c v2CoreVaultGBBankAccountService) Retrieve(ctx context.Context, id string, params *V2CoreVaultGBBankAccountRetrieveParams) (*V2CoreVaultGBBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultGBBankAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Archive a GBBankAccount object. Archived GBBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c v2CoreVaultGBBankAccountService) Archive(ctx context.Context, id string, params *V2CoreVaultGBBankAccountArchiveParams) (*V2CoreVaultGBBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultGBBankAccountArchiveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/gb_bank_accounts/%s/archive", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// List objects that can be used as destinations for outbound money movement via OutboundPayment.
func (c v2CoreVaultGBBankAccountService) List(ctx context.Context, listParams *V2CoreVaultGBBankAccountListParams) *V2List[*V2CoreVaultGBBankAccount] {
	if listParams == nil {
		listParams = &V2CoreVaultGBBankAccountListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/vault/gb_bank_accounts", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreVaultGBBankAccount], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreVaultGBBankAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
