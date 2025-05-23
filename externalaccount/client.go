//
//
// File generated from our OpenAPI spec
//
//

// Package externalaccount provides the /v1/external_accounts APIs
package externalaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/external_accounts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Delete a specified external account for a given account.
func DeleteBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().DeleteBankAccount(id, params)
}

// Delete a specified external account for a given account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeleteBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Delete a specified external account for a given account.
func DeleteCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().DeleteCard(id, params)
}

// Delete a specified external account for a given account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeleteCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, card)
	return card, err
}

// Retrieve a specified external account for a given account.
func GetBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().GetBankAccount(id, params)
}

// Retrieve a specified external account for a given account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) GetBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Retrieve a specified external account for a given account.
func GetCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().GetCard(id, params)
}

// Retrieve a specified external account for a given account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) GetCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, card)
	return card, err
}

// Create an external account for a given connected account.
func NewBankAccount(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().NewBankAccount(params)
}

// Create an external account for a given connected account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) NewBankAccount(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	body := &form.Values{}
	params.AppendToAsSourceOrExternalAccount(body)
	bankaccount := &stripe.BankAccount{}
	err := c.B.CallRaw(http.MethodPost, "/v1/external_accounts", c.Key, []byte(body.Encode()), &params.Params, bankaccount)
	return bankaccount, err
}

// Create an external account for a given connected account.
func NewCard(params *stripe.CardParams) (*stripe.Card, error) {
	return getC().NewCard(params)
}

// Create an external account for a given connected account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) NewCard(params *stripe.CardParams) (*stripe.Card, error) {
	body := &form.Values{}
	params.AppendToAsCardSourceOrExternalAccount(body, nil)
	card := &stripe.Card{}
	err := c.B.CallRaw(http.MethodPost, "/v1/external_accounts", c.Key, []byte(body.Encode()), &params.Params, card)
	return card, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func UpdateBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().UpdateBankAccount(id, params)
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateBankAccount(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	bankaccount := &stripe.BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
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
func UpdateCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().UpdateCard(id, params)
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateCard(id string, params *stripe.CardParams) (*stripe.Card, error) {
	path := stripe.FormatURLPath("/v1/external_accounts/%s", id)
	card := &stripe.Card{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, card)
	return card, err
}

// List external accounts for an account.
func ListBankAccount(params *stripe.BankAccountListParams) *Iter {
	return getC().ListBankAccount(params)
}

// List external accounts for an account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListBankAccount(listParams *stripe.BankAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.BankAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/external_accounts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for external accounts.
type Iter struct {
	*stripe.Iter
}

// Card returns the external account which the iterator is currently pointing to.
func (i *Iter) Card() *stripe.Card {
	return i.Current().(*stripe.Card)
}

// BankAccount returns the external account which the iterator is currently pointing to.
func (i *Iter) BankAccount() *stripe.BankAccount {
	return i.Current().(*stripe.BankAccount)
}

// CardList returns the current list object which the iterator is currently using. List objects will change as new API calls are made to continue pagination.
func (i *Iter) CardList() *stripe.CardList {
	return i.List().(*stripe.CardList)
}

// BankAccountList returns the current list object which the iterator is currently using. List objects will change as new API calls are made to continue pagination.
func (i *Iter) BankAccountList() *stripe.BankAccountList {
	return i.List().(*stripe.BankAccountList)
}

// List external accounts for an account.
func ListCard(params *stripe.CardListParams) *Iter {
	return getC().ListCard(params)
}

// List external accounts for an account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListCard(listParams *stripe.CardListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CardList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/external_accounts", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
