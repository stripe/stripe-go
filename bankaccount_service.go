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

// v1BankAccountService is used to invoke bankaccount related APIs.
type v1BankAccountService struct {
	B   Backend
	Key string
}

// New creates a new bank account
func (c v1BankAccountService) Create(ctx context.Context, params *BankAccountCreateParams) (*BankAccount, error) {
	if params == nil {
		params = &BankAccountCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts", StringValue(params.Token), StringValue(
			params.Customer), StringValue(params.Account))
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Get returns the details of a bank account.
func (c v1BankAccountService) Retrieve(ctx context.Context, id string, params *BankAccountRetrieveParams) (*BankAccount, error) {
	if params == nil {
		params = &BankAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func (c v1BankAccountService) Update(ctx context.Context, id string, params *BankAccountUpdateParams) (*BankAccount, error) {
	if params == nil {
		params = &BankAccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Delete a specified external account for a given account.
func (c v1BankAccountService) Delete(ctx context.Context, id string, params *BankAccountDeleteParams) (*BankAccount, error) {
	if params == nil {
		params = &BankAccountDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}
func (c v1BankAccountService) List(ctx context.Context, listParams *BankAccountListParams) Seq2[*BankAccount, error] {
	if listParams == nil {
		listParams = &BankAccountListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/external_accounts", StringValue(
			listParams.Account), StringValue(listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BankAccount, ListContainer, error) {
		list := &BankAccountList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
