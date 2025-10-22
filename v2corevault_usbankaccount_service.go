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
	if params == nil {
		params = &V2CoreVaultUSBankAccountCreateParams{}
	}
	params.Context = ctx
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/us_bank_accounts", c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Retrieve a USBankAccount object.
func (c v2CoreVaultUSBankAccountService) Retrieve(ctx context.Context, id string, params *V2CoreVaultUSBankAccountRetrieveParams) (*V2CoreVaultUSBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultUSBankAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Update a USBankAccount object. This is limited to supplying a previously empty routing_number field.
func (c v2CoreVaultUSBankAccountService) Update(ctx context.Context, id string, params *V2CoreVaultUSBankAccountUpdateParams) (*V2CoreVaultUSBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultUSBankAccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Archive a USBankAccount object. USBankAccount objects will not be automatically archived by Stripe.
// Archived USBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c v2CoreVaultUSBankAccountService) Archive(ctx context.Context, id string, params *V2CoreVaultUSBankAccountArchiveParams) (*V2CoreVaultUSBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultUSBankAccountArchiveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/us_bank_accounts/%s/archive", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Confirm microdeposits amounts or descriptor code that you have received from the Send Microdeposits request. Once you correctly confirm this, this US Bank Account will be verified and eligible to transfer funds with.
func (c v2CoreVaultUSBankAccountService) ConfirmMicrodeposits(ctx context.Context, id string, params *V2CoreVaultUSBankAccountConfirmMicrodepositsParams) (*V2CoreVaultUSBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultUSBankAccountConfirmMicrodepositsParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/vault/us_bank_accounts/%s/confirm_microdeposits", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Send microdeposits in order to verify your US Bank Account so it is eligible to transfer funds. This will start the verification process and you must Confirm Microdeposits to successfully verify your US Bank Account.
func (c v2CoreVaultUSBankAccountService) SendMicrodeposits(ctx context.Context, id string, params *V2CoreVaultUSBankAccountSendMicrodepositsParams) (*V2CoreVaultUSBankAccount, error) {
	if params == nil {
		params = &V2CoreVaultUSBankAccountSendMicrodepositsParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/vault/us_bank_accounts/%s/send_microdeposits", id)
	usbankaccount := &V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// List USBankAccount objects. Optionally filter by verification status.
func (c v2CoreVaultUSBankAccountService) List(ctx context.Context, listParams *V2CoreVaultUSBankAccountListParams) Seq2[*V2CoreVaultUSBankAccount, error] {
	if listParams == nil {
		listParams = &V2CoreVaultUSBankAccountListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/core/vault/us_bank_accounts", listParams, func(path string, p ParamsContainer) (*V2Page[*V2CoreVaultUSBankAccount], error) {
		page := &V2Page[*V2CoreVaultUSBankAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
