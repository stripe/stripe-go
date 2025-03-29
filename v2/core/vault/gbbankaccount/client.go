//
//
// File generated from our OpenAPI spec
//
//

// Package gbbankaccount provides the gbbankaccount related APIs
package gbbankaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke gbbankaccount related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a GB bank account.
func (c Client) New(params *stripe.V2CoreVaultGBBankAccountParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/gb_bank_accounts", c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Retrieve a GB bank account.
func (c Client) Get(id string, params *stripe.V2CoreVaultGBBankAccountParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Confirm that you have received the result of the Confirmation of Payee request, and that you are okay with
// proceeding to pay out to this bank account despite the account not matching, partially matching, or the service
// being unavailable. Once you confirm this, you will be able to send OutboundPayments, but this may lead to
// funds being sent to the wrong account, which we might not be able to recover.
func (c Client) AcknowledgeConfirmationOfPayee(id string, params *stripe.V2CoreVaultGBBankAccountAcknowledgeConfirmationOfPayeeParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath(
		"/v2/core/vault/gb_bank_accounts/%s/acknowledge_confirmation_of_payee", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Archive a GbBankAccount object. Archived GbBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c Client) Archive(id string, params *stripe.V2CoreVaultGBBankAccountArchiveParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s/archive", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Initiate Confirmation of Payee (CoP) in order to verify that the owner of a UK bank account matches
// who you expect. This must be done on all UK bank accounts before sending domestic OutboundPayments. If
// the result is a partial match or a non match, explicit acknowledgement using AcknowledgeConfirmationOfPayee
// is required before sending funds.
func (c Client) InitiateConfirmationOfPayee(id string, params *stripe.V2CoreVaultGBBankAccountInitiateConfirmationOfPayeeParams) (*stripe.V2CoreVaultGBBankAccount, error) {
	path := stripe.FormatURLPath(
		"/v2/core/vault/gb_bank_accounts/%s/initiate_confirmation_of_payee", id)
	gbbankaccount := &stripe.V2CoreVaultGBBankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}
