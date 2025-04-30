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

// v1ExternalAccountService is used to invoke /v1/external_accounts APIs.
type v1ExternalAccountService struct {
	B   Backend
	Key string
}

// Create an external account for a given connected account.
func (c v1ExternalAccountService) CreateBankAccount(ctx context.Context, params *BankAccountParams) (*BankAccount, error) {
	bankaccount := &BankAccount{}
	if params == nil {
		params = &BankAccountParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/external_accounts", c.Key, params, bankaccount)
	return bankaccount, err
}

// Create an external account for a given connected account.
func (c v1ExternalAccountService) CreateCard(ctx context.Context, params *CardParams) (*Card, error) {
	card := &Card{}
	if params == nil {
		params = &CardParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, "/v1/external_accounts", c.Key, params, card)
	return card, err
}

// Delete a specified external account for a given account.
func (c v1ExternalAccountService) DeleteBankAccount(ctx context.Context, id string, params *BankAccountParams) (*BankAccount, error) {
	path := FormatURLPath("/v1/external_accounts/%s", id)
	bankaccount := &BankAccount{}
	if params == nil {
		params = &BankAccountParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Delete a specified external account for a given account.
func (c v1ExternalAccountService) DeleteCard(ctx context.Context, id string, params *CardParams) (*Card, error) {
	path := FormatURLPath("/v1/external_accounts/%s", id)
	card := &Card{}
	if params == nil {
		params = &CardParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}

// Retrieve a specified external account for a given account.
func (c v1ExternalAccountService) RetrieveBankAccount(ctx context.Context, id string, params *BankAccountParams) (*BankAccount, error) {
	body := &form.Values{}
	params.AppendToAsSourceOrExternalAccount(body)
	bankaccount := &BankAccount{}
	err := c.B.CallRaw(http.MethodPost, "/v1/external_accounts", c.Key, []byte(body.Encode()), &params.Params, bankaccount)
	return bankaccount, err
}

// Retrieve a specified external account for a given account.
func (c v1ExternalAccountService) RetrieveCard(ctx context.Context, id string, params *CardParams) (*Card, error) {
	body := &form.Values{}
	params.AppendToAsCardSourceOrExternalAccount(body, nil)
	card := &Card{}
	err := c.B.CallRaw(http.MethodPost, "/v1/external_accounts", c.Key, []byte(body.Encode()), &params.Params, card)
	return card, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func (c v1ExternalAccountService) UpdateBankAccount(ctx context.Context, id string, params *BankAccountParams) (*BankAccount, error) {
	path := FormatURLPath("/v1/external_accounts/%s", id)
	bankaccount := &BankAccount{}
	if params == nil {
		params = &BankAccountParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func (c v1ExternalAccountService) UpdateCard(ctx context.Context, id string, params *CardParams) (*Card, error) {
	path := FormatURLPath("/v1/external_accounts/%s", id)
	card := &Card{}
	if params == nil {
		params = &CardParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// List external accounts for an account
func (c v1ExternalAccountService) ListBankAccount(ctx context.Context, listParams *BankAccountListParams) Seq2[*BankAccount, error] {
	if listParams == nil {
		listParams = &BankAccountListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BankAccount, ListContainer, error) {
		list := &BankAccountList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/external_accounts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}

// List external accounts for an account
func (c v1ExternalAccountService) ListCard(ctx context.Context, listParams *CardListParams) Seq2[*Card, error] {
	if listParams == nil {
		listParams = &CardListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Card, ListContainer, error) {
		list := &CardList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/external_accounts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
