//
//
// File generated from our OpenAPI spec
//
//

// Package association provides the association related APIs
package association

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke association related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Finds a tax association object by PaymentIntent id.
func Find(params *stripe.TaxAssociationFindParams) (*stripe.TaxAssociation, error) {
	return getC().Find(params)
}

// Finds a tax association object by PaymentIntent id.
func (c Client) Find(params *stripe.TaxAssociationFindParams) (*stripe.TaxAssociation, error) {
	association := &stripe.TaxAssociation{}
	err := c.B.Call(
		http.MethodGet, "/v1/tax/associations/find", c.Key, params, association)
	return association, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
