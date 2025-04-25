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
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	if params == nil {
		params = &V2CoreVaultGBBankAccountCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/gb_bank_accounts", c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Retrieve a GB bank account.
func (c v2CoreVaultGBBankAccountService) Retrieve(ctx context.Context, id string, params *V2CoreVaultGBBankAccountRetrieveParams) (*V2CoreVaultGBBankAccount, error) {
	path := FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	if params == nil {
		params = &V2CoreVaultGBBankAccountRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Confirm that you have received the result of the Confirmation of Payee request, and that you are okay with
// proceeding to pay out to this bank account despite the account not matching, partially matching, or the service
// being unavailable. Once you confirm this, you will be able to send OutboundPayments, but this may lead to
// funds being sent to the wrong account, which we might not be able to recover.
func (c v2CoreVaultGBBankAccountService) AcknowledgeConfirmationOfPayee(ctx context.Context, id string, params *V2CoreVaultGBBankAccountAcknowledgeConfirmationOfPayeeParams) (*V2CoreVaultGBBankAccount, error) {
	path := FormatURLPath(
		"/v2/core/vault/gb_bank_accounts/%s/acknowledge_confirmation_of_payee", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	if params == nil {
		params = &V2CoreVaultGBBankAccountAcknowledgeConfirmationOfPayeeParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Archive a GBBankAccount object. Archived GBBankAccount objects cannot be used as outbound destinations
// and will not appear in the outbound destination list.
func (c v2CoreVaultGBBankAccountService) Archive(ctx context.Context, id string, params *V2CoreVaultGBBankAccountArchiveParams) (*V2CoreVaultGBBankAccount, error) {
	path := FormatURLPath("/v2/core/vault/gb_bank_accounts/%s/archive", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	if params == nil {
		params = &V2CoreVaultGBBankAccountArchiveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}

// Initiate Confirmation of Payee (CoP) in order to verify that the owner of a UK bank account matches
// who you expect. This must be done on all UK bank accounts before sending domestic OutboundPayments. If
// the result is a partial match or a non match, explicit acknowledgement using AcknowledgeConfirmationOfPayee
// is required before sending funds.
func (c v2CoreVaultGBBankAccountService) InitiateConfirmationOfPayee(ctx context.Context, id string, params *V2CoreVaultGBBankAccountInitiateConfirmationOfPayeeParams) (*V2CoreVaultGBBankAccount, error) {
	path := FormatURLPath(
		"/v2/core/vault/gb_bank_accounts/%s/initiate_confirmation_of_payee", id)
	gbbankaccount := &V2CoreVaultGBBankAccount{}
	if params == nil {
		params = &V2CoreVaultGBBankAccountInitiateConfirmationOfPayeeParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, gbbankaccount)
	return gbbankaccount, err
}
