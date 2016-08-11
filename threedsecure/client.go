// package threedsecure provides the /3d_secure APIs
package threedsecure

import (
	"strconv"

	"github.com/stripe/stripe-go"
)

type Client struct {
	B   stripe.Backend
	Key string
}

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

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
