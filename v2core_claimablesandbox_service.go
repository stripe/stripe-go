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

// v2CoreClaimableSandboxService is used to invoke claimablesandbox related APIs.
type v2CoreClaimableSandboxService struct {
	B   Backend
	Key string
}

// Create an anonymous, claimable sandbox. This sandbox can be prefilled with data. The response will include
// a claim URL that allow a user to claim the account.
func (c v2CoreClaimableSandboxService) Create(ctx context.Context, params *V2CoreClaimableSandboxCreateParams) (*V2CoreClaimableSandbox, error) {
	if params == nil {
		params = &V2CoreClaimableSandboxCreateParams{}
	}
	params.Context = ctx
	claimablesandbox := &V2CoreClaimableSandbox{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/claimable_sandboxes", c.Key, params, claimablesandbox)
	return claimablesandbox, err
}
