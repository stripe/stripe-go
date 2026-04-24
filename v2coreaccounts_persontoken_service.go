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

// v2CoreAccountsPersonTokenService is used to invoke persontoken related APIs.
type v2CoreAccountsPersonTokenService struct {
	B   Backend
	Key string
}

// Creates a single-use token that represents the details for a person. Use this when you create or update persons associated with an Account v2. Learn more about [account tokens](https://docs.stripe.com/connect/account-tokens).
// You can only create person tokens with your application's publishable key and in live mode. You can use your application's secret key to create person tokens only in test mode.
func (c v2CoreAccountsPersonTokenService) Create(ctx context.Context, params *V2CoreAccountsPersonTokenCreateParams) (*V2CoreAccountPersonToken, error) {
	if params == nil {
		params = &V2CoreAccountsPersonTokenCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/person_tokens", StringValue(params.AccountID))
	accountpersontoken := &V2CoreAccountPersonToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountpersontoken)
	return accountpersontoken, err
}

// Retrieves a Person Token associated with an Account.
func (c v2CoreAccountsPersonTokenService) Retrieve(ctx context.Context, id string, params *V2CoreAccountsPersonTokenRetrieveParams) (*V2CoreAccountPersonToken, error) {
	if params == nil {
		params = &V2CoreAccountsPersonTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/accounts/%s/person_tokens/%s", StringValue(params.AccountID), id)
	accountpersontoken := &V2CoreAccountPersonToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountpersontoken)
	return accountpersontoken, err
}
