//
//
// File generated from our OpenAPI spec
//
//

// Package authentication provides the /v1/three_d_secure/authentications APIs
package authentication

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/three_d_secure/authentications APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// This endpoint creates a 3DS Authentication. Refer to the [Create a 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#create-a-3ds-authentication-object) for more information.
//
// You can pass the submit parameter to automatically submit the 3DS Authentication object when you create it. Refer to the [Submit at creation section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
func New(params *stripe.ThreeDSecureAuthenticationParams) (*stripe.ThreeDSecureAuthentication, error) {
	return getC().New(params)
}

// This endpoint creates a 3DS Authentication. Refer to the [Create a 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#create-a-3ds-authentication-object) for more information.
//
// You can pass the submit parameter to automatically submit the 3DS Authentication object when you create it. Refer to the [Submit at creation section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.ThreeDSecureAuthenticationParams) (*stripe.ThreeDSecureAuthentication, error) {
	authentication := &stripe.ThreeDSecureAuthentication{}
	err := c.B.Call(
		http.MethodPost, "/v1/three_d_secure/authentications", c.Key, params, authentication)
	return authentication, err
}

// This endpoint retrieves a 3DS Authentication.
func Get(id string, params *stripe.ThreeDSecureAuthenticationParams) (*stripe.ThreeDSecureAuthentication, error) {
	return getC().Get(id, params)
}

// This endpoint retrieves a 3DS Authentication.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ThreeDSecureAuthenticationParams) (*stripe.ThreeDSecureAuthentication, error) {
	path := stripe.FormatURLPath("/v1/three_d_secure/authentications/%s", id)
	authentication := &stripe.ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, authentication)
	return authentication, err
}

// This endpoint cancels a 3DS Authentication. You can cancel a 3DS Authentication object when it's in a non-final status:
// requires_submission or requires_challenge.
func Cancel(id string, params *stripe.ThreeDSecureAuthenticationCancelParams) (*stripe.ThreeDSecureAuthentication, error) {
	return getC().Cancel(id, params)
}

// This endpoint cancels a 3DS Authentication. You can cancel a 3DS Authentication object when it's in a non-final status:
// requires_submission or requires_challenge.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.ThreeDSecureAuthenticationCancelParams) (*stripe.ThreeDSecureAuthentication, error) {
	path := stripe.FormatURLPath(
		"/v1/three_d_secure/authentications/%s/cancel", id)
	authentication := &stripe.ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authentication)
	return authentication, err
}

// This endpoint submits a 3DS Authentication. You can submit a 3DS Authentication object when it has status requires_submission. Refer to the [Submit the 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-the-3ds-authentication-object) for more information.
func Submit(id string, params *stripe.ThreeDSecureAuthenticationSubmitParams) (*stripe.ThreeDSecureAuthentication, error) {
	return getC().Submit(id, params)
}

// This endpoint submits a 3DS Authentication. You can submit a 3DS Authentication object when it has status requires_submission. Refer to the [Submit the 3DS Authentication object section of the Standalone 3DS guide](https://docs.stripe.com/payments/3d-secure/standalone-3d-secure#submit-the-3ds-authentication-object) for more information.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Submit(id string, params *stripe.ThreeDSecureAuthenticationSubmitParams) (*stripe.ThreeDSecureAuthentication, error) {
	path := stripe.FormatURLPath(
		"/v1/three_d_secure/authentications/%s/submit", id)
	authentication := &stripe.ThreeDSecureAuthentication{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authentication)
	return authentication, err
}

// Returns a list of 3D Secure Authentications.
func List(params *stripe.ThreeDSecureAuthenticationListParams) *Iter {
	return getC().List(params)
}

// Returns a list of 3D Secure Authentications.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ThreeDSecureAuthenticationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ThreeDSecureAuthenticationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/three_d_secure/authentications", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for three dsecure authentications.
type Iter struct {
	*stripe.Iter
}

// ThreeDSecureAuthentication returns the three dsecure authentication which the iterator is currently pointing to.
func (i *Iter) ThreeDSecureAuthentication() *stripe.ThreeDSecureAuthentication {
	return i.Current().(*stripe.ThreeDSecureAuthentication)
}

// ThreeDSecureAuthenticationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ThreeDSecureAuthenticationList() *stripe.ThreeDSecureAuthenticationList {
	return i.List().(*stripe.ThreeDSecureAuthenticationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
