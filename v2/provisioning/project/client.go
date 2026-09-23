//
//
// File generated from our OpenAPI spec
//
//

// Package project provides the project related APIs
package project

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke project related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new project.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2ProvisioningProjectParams) (*stripe.V2ProvisioningProject, error) {
	project := &stripe.V2ProvisioningProject{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/projects", c.Key, params, project)
	return project, err
}
