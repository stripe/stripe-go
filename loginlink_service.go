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

// v1LoginLinkService is used to invoke /v1/accounts/{account}/login_links APIs.
type v1LoginLinkService struct {
	B   Backend
	Key string
}

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
func (c v1LoginLinkService) Create(ctx context.Context, params *LoginLinkCreateParams) (*LoginLink, error) {
	if params == nil {
		params = &LoginLinkCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/accounts/%s/login_links", StringValue(params.Account))
	loginlink := &LoginLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, loginlink)
	return loginlink, err
}
