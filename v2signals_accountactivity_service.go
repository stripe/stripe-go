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

// v2SignalsAccountActivityService is used to invoke accountactivity related APIs.
type v2SignalsAccountActivityService struct {
	B   Backend
	Key string
}

// Creates a new account activity to report account registration, login, or evaluation follow-up activity.
func (c v2SignalsAccountActivityService) Create(ctx context.Context, params *V2SignalsAccountActivityCreateParams) (*V2SignalsAccountActivity, error) {
	if params == nil {
		params = &V2SignalsAccountActivityCreateParams{}
	}
	params.Context = ctx
	accountactivity := &V2SignalsAccountActivity{}
	err := c.B.Call(
		http.MethodPost, "/v2/signals/account_activity", c.Key, params, accountactivity)
	return accountactivity, err
}

// Retrieves an AccountActivity by its ID.
func (c v2SignalsAccountActivityService) Retrieve(ctx context.Context, id string, params *V2SignalsAccountActivityRetrieveParams) (*V2SignalsAccountActivity, error) {
	if params == nil {
		params = &V2SignalsAccountActivityRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/account_activity/%s", id)
	accountactivity := &V2SignalsAccountActivity{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountactivity)
	return accountactivity, err
}

// Deletes an AccountActivity by its ID.
func (c v2SignalsAccountActivityService) Delete(ctx context.Context, id string, params *V2SignalsAccountActivityDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2SignalsAccountActivityDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/account_activity/%s", id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}
