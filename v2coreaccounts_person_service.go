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

// v2CoreAccountsPersonService is used to invoke person related APIs.
type v2CoreAccountsPersonService struct {
	B   Backend
	Key string
}

// Create a Person associated with an Account.
func (c v2CoreAccountsPersonService) Create(ctx context.Context, params *V2CoreAccountsPersonCreateParams) (*V2CorePerson, error) {
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons", StringValue(params.AccountID))
	person := &V2CorePerson{}
	if params == nil {
		params = &V2CoreAccountsPersonCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Retrieves a Person associated with an Account.
func (c v2CoreAccountsPersonService) Retrieve(ctx context.Context, id string, params *V2CoreAccountsPersonRetrieveParams) (*V2CorePerson, error) {
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	person := &V2CorePerson{}
	if params == nil {
		params = &V2CoreAccountsPersonRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, person)
	return person, err
}

// Updates a Person associated with an Account.
func (c v2CoreAccountsPersonService) Update(ctx context.Context, id string, params *V2CoreAccountsPersonUpdateParams) (*V2CorePerson, error) {
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	person := &V2CorePerson{}
	if params == nil {
		params = &V2CoreAccountsPersonUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, person)
	return person, err
}

// Delete a Person associated with an Account.
func (c v2CoreAccountsPersonService) Delete(ctx context.Context, id string, params *V2CoreAccountsPersonDeleteParams) (*V2CorePerson, error) {
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	person := &V2CorePerson{}
	if params == nil {
		params = &V2CoreAccountsPersonDeleteParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodDelete, path, c.Key, params, person)
	return person, err
}

// Returns a list of Persons associated with an Account.
func (c v2CoreAccountsPersonService) List(ctx context.Context, listParams *V2CoreAccountsPersonListParams) Seq2[*V2CorePerson, error] {
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons", StringValue(listParams.AccountID))
	if listParams == nil {
		listParams = &V2CoreAccountsPersonListParams{}
	}
	listParams.Context = ctx
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2CorePerson], error) {
		page := &V2Page[*V2CorePerson]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
