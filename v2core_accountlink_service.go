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

// v2CoreAccountLinkService is used to invoke accountlink related APIs.
type v2CoreAccountLinkService struct {
	B   Backend
	Key string
}

// Creates an AccountLink object that includes a single-use Stripe URL that the merchant can redirect their user to in order to take them to a Stripe-hosted application such as Recipient Onboarding.
func (c v2CoreAccountLinkService) Create(ctx context.Context, params *V2CoreAccountLinkCreateParams) (*V2CoreAccountLink, error) {
	if params == nil {
		params = &V2CoreAccountLinkCreateParams{}
	}
	params.Context = ctx
	accountlink := &V2CoreAccountLink{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_links", c.Key, params, accountlink)
	return accountlink, err
}
