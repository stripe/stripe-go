//
//
// File generated from our OpenAPI spec
//
//

// Package eligibility provides the eligibility related APIs
package eligibility

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke eligibility related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Checks whether a project is eligible to provision resources with a provider, including
// any outstanding KYC requirements that must be satisfied first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.V2ProvisioningEligibilityParams) (*stripe.V2ProvisioningEligibility, error) {
	eligibility := &stripe.V2ProvisioningEligibility{}
	err := c.B.Call(
		http.MethodGet, "/v2/provisioning/eligibility", c.Key, params, eligibility)
	return eligibility, err
}
