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

// v2CoreVaultUSBankAccountService is used to invoke usbankaccount related APIs.
type v2CoreVaultUSBankAccountService struct {
	B   Backend
	Key string
}

// Create a USBankAccount object.
func (c v2CoreVaultUSBankAccountService) Create(ctx context.Context, params *V2CoreVaultUSBankAccountCreateParams) (*V2CoreVaultUSBankAccount, error) {
	usbankaccount := &V2CoreVaultUSBankAccount{}
	if params == nil {
		params = &V2CoreVaultUSBankAccountCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/us_bank_accounts", c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Retrieve a USBankAccount object.
func (c v2CoreVaultUSBankAccountService) Retrieve(ctx context.Context, id string, params *V2CoreVaultUSBankAccountRetrieveParams) (*V2CoreVaultUSBankAccount, error) {
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	if params == nil {
		params = &V2CoreVaultUSBankAccountRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Update a USBankAccount object. This is limited to supplying a previously empty routing_number field.
func (c v2CoreVaultUSBankAccountService) Update(ctx context.Context, id string, params *V2CoreVaultUSBankAccountUpdateParams) (*V2CoreVaultUSBankAccount, error) {
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	if params == nil {
		params = &V2CoreVaultUSBankAccountUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Archive a USBankAccount object. USBankAccount objects will not be automatically archived by Stripe.
// Archived USBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c v2CoreVaultUSBankAccountService) Archive(ctx context.Context, id string, params *V2CoreVaultUSBankAccountArchiveParams) (*V2CoreVaultUSBankAccount, error) {
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s/archive", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	if params == nil {
		params = &V2CoreVaultUSBankAccountArchiveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}
