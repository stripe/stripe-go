//
//
// File generated from our OpenAPI spec
//
//

// Package person provides the person related APIs
package person

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke person related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Person. Adds an individual to an Account's identity. You can set relationship attributes and identity information at creation.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons", stripe.StringValue(params.AccountID))
	accountperson := &stripe.V2CoreAccountPerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountperson)
	return accountperson, err
}

// Retrieves a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	accountperson := &stripe.V2CoreAccountPerson{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountperson)
	return accountperson, err
}

// Updates a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CoreAccountPerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	accountperson := &stripe.V2CoreAccountPerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountperson)
	return accountperson, err
}

// Delete a Person associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// Returns a paginated list of Persons associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreAccountsPersonListParams) stripe.Seq2[*stripe.V2CoreAccountPerson, error] {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons", stripe.StringValue(listParams.AccountID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreAccountPerson], error) {
		page := &stripe.V2Page[*stripe.V2CoreAccountPerson]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
