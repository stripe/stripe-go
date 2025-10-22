//
//
// File generated from our OpenAPI spec
//
//

// Package usbankaccount provides the usbankaccount related APIs
package usbankaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke usbankaccount related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a USBankAccount object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/us_bank_accounts", c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Retrieve a USBankAccount object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Update a USBankAccount object. This is limited to supplying a previously empty routing_number field.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Archive a USBankAccount object. USBankAccount objects will not be automatically archived by Stripe.
// Archived USBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Archive(id string, params *stripe.V2CoreVaultUSBankAccountArchiveParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s/archive", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Confirm microdeposits amounts or descriptor code that you have received from the Send Microdeposits request. Once you correctly confirm this, this US Bank Account will be verified and eligible to transfer funds with.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ConfirmMicrodeposits(id string, params *stripe.V2CoreVaultUSBankAccountConfirmMicrodepositsParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath(
		"/v2/core/vault/us_bank_accounts/%s/confirm_microdeposits", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Send microdeposits in order to verify your US Bank Account so it is eligible to transfer funds. This will start the verification process and you must Confirm Microdeposits to successfully verify your US Bank Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SendMicrodeposits(id string, params *stripe.V2CoreVaultUSBankAccountSendMicrodepositsParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath(
		"/v2/core/vault/us_bank_accounts/%s/send_microdeposits", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// List USBankAccount objects. Optionally filter by verification status.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreVaultUSBankAccountListParams) stripe.Seq2[*stripe.V2CoreVaultUSBankAccount, error] {
	return stripe.NewV2List("/v2/core/vault/us_bank_accounts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreVaultUSBankAccount], error) {
		page := &stripe.V2Page[*stripe.V2CoreVaultUSBankAccount]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
