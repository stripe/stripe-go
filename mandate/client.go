// Package mandate provides the /mandates APIs
package mandate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke mandates APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a Mandate.
func Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	return getC().Get(id, params)
}

// Get returns the details of a Mandate.
func (c Client) Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	path := stripe.FormatURLPath("/v1/mandates/%s", id)
	mandate := &stripe.Mandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mandate)
	return mandate, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
