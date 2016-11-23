// Package threedsecure provides the /3d_secure APIs
package threedsecure

import (
	"strconv"

	"github.com/stripe/stripe-go"
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
	body := &stripe.RequestValues{}
	body.Add("amount", strconv.FormatUint(params.Amount, 10))
	body.Add("card", params.Card)
	body.Add("currency", string(params.Currency))
	body.Add("return_url", params.ReturnURL)

	if len(params.Customer) > 0 {
		body.Add("customer", params.Customer)
	}

	params.AppendTo(body)

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
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		commonParams = &params.Params
		params.AppendTo(body)
	}

	tds := &stripe.ThreeDSecure{}
	err := c.B.Call("GET", "/3d_secure/"+id, c.Key, body, commonParams, tds)

	return tds, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
