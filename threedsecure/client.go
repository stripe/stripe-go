// Package threedsecure provides the /3d_secure APIs
// This package is deprecated and should not be used anymore.
// It is still here for backwards compatibility for now.
// To use 3D Secure, please use the source package instead.
package threedsecure

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /3d_secure APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new 3D Secure object.
func New(params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	return getC().New(params)
}

// New creates a new 3D Secure object.
func (c Client) New(params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	tds := &stripe.ThreeDSecure{}
	err := c.B.Call(http.MethodPost, "/v1/3d_secure", c.Key, params, tds)
	return tds, err
}

// Get returns the details of a 3D Secure object.
func Get(id string, params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	return getC().Get(id, params)
}

// Get returns the details of a 3D Secure object.
func (c Client) Get(id string, params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	path := stripe.FormatURLPath("/v1/3d_secure/%s", id)
	tds := &stripe.ThreeDSecure{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, tds)

	return tds, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
