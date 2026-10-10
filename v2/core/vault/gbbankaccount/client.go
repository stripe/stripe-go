//
//
// File generated from our OpenAPI spec
//
//

// Package gbbankaccount provides the gbbankaccount related APIs
package gbbankaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke gbbankaccount related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a GB bank account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreVaultGBBankAccountParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/gb_bank_accounts", c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Retrieve a GB bank account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreVaultGBBankAccountParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Archive a GBBankAccount object. Archived GBBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Archive(id string, params *stripe.V2CoreVaultGBBankAccountArchiveParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s/archive", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// List objects that can be used as destinations for outbound money movement via OutboundPayment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreVaultGBBankAccountListParams) stripe.Seq2[*stripe.V2CoreVaultGBBankAccount, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreVaultGBBankAccountListParams{}
	}
	return stripe.NewV2List("/v2/core/vault/gb_bank_accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreVaultGBBankAccount], error) {
		page := &stripe.V2Page[*stripe.V2CoreVaultGBBankAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
