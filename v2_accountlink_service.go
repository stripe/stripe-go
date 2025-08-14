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

// v2AccountLinkService is used to invoke accountlink related APIs.
type v2AccountLinkService struct {
	B   Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
func (c v2AccountLinkService) Create(ctx context.Context, params *V2AccountLinkCreateParams) (*V2AccountLink, error) {
	if params == nil {
		params = &V2AccountLinkCreateParams{}
	}
	params.Context = ctx
	accountlink := &V2AccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v2/account_links", c.Key, params, accountlink)
	return accountlink, err
}
