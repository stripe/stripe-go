// Package threedsecure provides the /3d_secure APIs
package threedsecure

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /3d_secure APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new 3D Secure auths.
// For more details see https://stripe.com/docs/api#create_three_d_secure.
func New(params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	tds := &stripe.ThreeDSecure{}
	err := c.B.Call("POST", "/3d_secure", c.Key, body, &params.Params, tds)
	return tds, err
}

// Get returns the details of a 3D Secure auth.
// For more details see https://stripe.com/docs/api#retrieve_three_d_secure.
func Get(id string, params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ThreeDSecureParams) (*stripe.ThreeDSecure, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	tds := &stripe.ThreeDSecure{}
	err := c.B.Call("GET", "/3d_secure/"+id, c.Key, body, commonParams, tds)

	return tds, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
