//
//
// File generated from our OpenAPI spec
//
//

// Package person provides the person related APIs
package person

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke person related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Person associated with an Account.
func (c Client) New(params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CorePerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons", stripe.StringValue(params.AccountID))
	person := &stripe.V2CorePerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Retrieves a Person associated with an Account.
func (c Client) Get(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CorePerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	person := &stripe.V2CorePerson{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, person)
	return person, err
}

// Updates a Person associated with an Account.
func (c Client) Update(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CorePerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	person := &stripe.V2CorePerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Delete a Person associated with an Account.
func (c Client) Del(id string, params *stripe.V2CoreAccountsPersonParams) (*stripe.V2CorePerson, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", stripe.StringValue(params.AccountID), id)
	person := &stripe.V2CorePerson{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, person)
	return person, err
}

// Returns a list of Persons associated with an Account.
func (c Client) All(listParams *stripe.V2CoreAccountsPersonListParams) stripe.Seq2[*stripe.V2CorePerson, error] {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/persons", stripe.StringValue(listParams.AccountID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CorePerson], error) {
		page := &stripe.V2Page[*stripe.V2CorePerson]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
