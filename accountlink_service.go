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

// v1AccountLinkService is used to invoke /v1/account_links APIs.
type v1AccountLinkService struct {
	B   Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the platform can redirect their user to in order to take them through the Connect Onboarding flow.
func (c v1AccountLinkService) Create(ctx context.Context, params *AccountLinkCreateParams) (*AccountLink, error) {
	if params == nil {
		params = &AccountLinkCreateParams{}
	}
	params.Context = ctx
	accountlink := &AccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v1/account_links", c.Key, params, accountlink)
	return accountlink, err
}
