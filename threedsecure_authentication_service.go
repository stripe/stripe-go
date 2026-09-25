//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1ThreeDSecureAuthenticationService is used to invoke /v1/three_d_secure/authentications APIs.
type v1ThreeDSecureAuthenticationService struct {
	B   Backend
	Key string
}

// This endpoint creates a 3DS Authentication. Refer to the [Create a 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#create-a-3ds-authentication-object) for more information.
//
// You can pass the submit parameter to automatically submit the 3DS Authentication object when you create it. Refer to the [Submit at creation section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
func (c v1ThreeDSecureAuthenticationService) Create(ctx context.Context, params *ThreeDSecureAuthenticationCreateParams) (*ThreeDSecureAuthentication, error) {
	if params == nil {
		params = &ThreeDSecureAuthenticationCreateParams{}
	}
	params.Context = ctx
	authentication := &ThreeDSecureAuthentication{}
	err := c.B.Call(
		http.MethodPost, "/v1/three_d_secure/authentications", c.Key, params, authentication)
	return authentication, err
}

// This endpoint retrieves a 3DS Authentication.
func (c v1ThreeDSecureAuthenticationService) Retrieve(ctx context.Context, id string, params *ThreeDSecureAuthenticationRetrieveParams) (*ThreeDSecureAuthentication, error) {
	if params == nil {
		params = &ThreeDSecureAuthenticationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/three_d_secure/authentications/%s", id)
	authentication := &ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, authentication)
	return authentication, err
}

// This endpoint cancels a 3DS Authentication. You can cancel a 3DS Authentication object when it's in a non-final status:
// requires_submission or requires_challenge.
func (c v1ThreeDSecureAuthenticationService) Cancel(ctx context.Context, id string, params *ThreeDSecureAuthenticationCancelParams) (*ThreeDSecureAuthentication, error) {
	if params == nil {
		params = &ThreeDSecureAuthenticationCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/three_d_secure/authentications/%s/cancel", id)
	authentication := &ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authentication)
	return authentication, err
}

// This endpoint submits a 3DS Authentication. You can submit a 3DS Authentication object when it has status requires_submission. Refer to the [Submit the 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-the-3ds-authentication-object) for more information.
func (c v1ThreeDSecureAuthenticationService) Submit(ctx context.Context, id string, params *ThreeDSecureAuthenticationSubmitParams) (*ThreeDSecureAuthentication, error) {
	if params == nil {
		params = &ThreeDSecureAuthenticationSubmitParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/three_d_secure/authentications/%s/submit", id)
	authentication := &ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authentication)
	return authentication, err
}

// Returns a list of 3D Secure Authentications.
func (c v1ThreeDSecureAuthenticationService) List(ctx context.Context, listParams *ThreeDSecureAuthenticationListParams) *V1List[*ThreeDSecureAuthentication] {
	if listParams == nil {
		listParams = &ThreeDSecureAuthenticationListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*ThreeDSecureAuthentication], error) {
		list := &v1Page[*ThreeDSecureAuthentication]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/three_d_secure/authentications", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
