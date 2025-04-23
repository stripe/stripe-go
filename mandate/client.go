//
//
// File generated from our OpenAPI spec
//
//

// Package mandate provides the /v1/mandates APIs
package mandate

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke /v1/mandates APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Mandate object.
func Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	return getC().Get(id, params)
}

// Retrieves a Mandate object.
func (c Client) Get(id string, params *stripe.MandateParams) (*stripe.Mandate, error) {
	path := stripe.FormatURLPath("/v1/mandates/%s", id)
	mandate := &stripe.Mandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mandate)
	return mandate, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
