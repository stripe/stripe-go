//
//
// File generated from our OpenAPI spec
//
//

// Package usbankaccount provides the usbankaccount related APIs
package usbankaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke usbankaccount related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a UsBankAccount object.
func (c Client) New(params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/us_bank_accounts", c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Retrieve a UsBankAccount object.
func (c Client) Get(id string, params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Update a UsBankAccount object. This is limited to supplying a previously empty routing_number field.
func (c Client) Update(id string, params *stripe.V2CoreVaultUSBankAccountParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}

// Archive a UsBankAccount object. UsBankAccount objects will not be automatically archived by Stripe.
// Archived UsBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c Client) Archive(id string, params *stripe.V2CoreVaultUSBankAccountArchiveParams) (*stripe.V2CoreVaultUSBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s/archive", id)
	usbankaccount := &stripe.V2CoreVaultUSBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, usbankaccount)
	return usbankaccount, err
}
