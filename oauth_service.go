package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// oauthService is used to invoke /oauth and related APIs.
type oauthService struct {
	B   Backend
	Key string
}

// AuthorizeURL builds an OAuth authorize URL.
func (c oauthService) AuthorizeURL(params *AuthorizeURLParams) string {
	express := ""
	if BoolValue(params.Express) {
		express = "/express"
	}
	qs := &form.Values{}
	form.AppendTo(qs, params)
	return fmt.Sprintf("%s%s/oauth/authorize?%s", ConnectURL, express, qs.Encode())
}

// Create creates an OAuth token using a code after successful redirection back.
func (c oauthService) Create(ctx context.Context, params *OAuthTokenParams) (*OAuthToken, error) {
	if params == nil {
		params = &OAuthTokenParams{}
	}
	params.Context = ctx
	if params.ClientSecret == nil {
		params.ClientSecret = String(c.Key)
	}
	oauthToken := &OAuthToken{}
	err := c.B.Call(http.MethodPost, "/oauth/token", c.Key, params, oauthToken)
	return oauthToken, err
}

// Delete deauthorizes a connected account.
func (c oauthService) Delete(ctx context.Context, params *DeauthorizeParams) (*Deauthorize, error) {
	if params == nil {
		params = &DeauthorizeParams{}
	}
	params.Context = ctx
	deauthorization := &Deauthorize{}
	err := c.B.Call(http.MethodPost, "/oauth/deauthorize", c.Key, params, deauthorization)
	return deauthorization, err
}
