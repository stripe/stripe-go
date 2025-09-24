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
func (c v2CoreAccountsPersonService) Create(ctx context.Context, params *V2CoreAccountsPersonCreateParams) (*V2CoreAccountPerson, error) {
	if params == nil {
		params = &V2CoreAccountsPersonCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons", StringValue(params.AccountID))
	accountperson := &V2CoreAccountPerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountperson)
	return accountperson, err
}

// Retrieves a Person associated with an Account.
func (c v2CoreAccountsPersonService) Retrieve(ctx context.Context, id string, params *V2CoreAccountsPersonRetrieveParams) (*V2CoreAccountPerson, error) {
	if params == nil {
		params = &V2CoreAccountsPersonRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	accountperson := &V2CoreAccountPerson{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountperson)
	return accountperson, err
}

// Updates a Person associated with an Account.
func (c v2CoreAccountsPersonService) Update(ctx context.Context, id string, params *V2CoreAccountsPersonUpdateParams) (*V2CoreAccountPerson, error) {
	if params == nil {
		params = &V2CoreAccountsPersonUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	accountperson := &V2CoreAccountPerson{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountperson)
	return accountperson, err
}

// Delete a Person associated with an Account.
func (c v2CoreAccountsPersonService) Delete(ctx context.Context, id string, params *V2CoreAccountsPersonDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2CoreAccountsPersonDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons/%s", StringValue(params.AccountID), id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// Returns a list of Persons associated with an Account.
func (c v2CoreAccountsPersonService) List(ctx context.Context, listParams *V2CoreAccountsPersonListParams) Seq2[*V2CoreAccountPerson, error] {
	if listParams == nil {
		listParams = &V2CoreAccountsPersonListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/persons", StringValue(listParams.AccountID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2CoreAccountPerson], error) {
		page := &V2Page[*V2CoreAccountPerson]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
