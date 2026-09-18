//
//
// File generated from our OpenAPI spec
//
//

// Package paymentprofile provides the paymentprofile related APIs
package paymentprofile

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke paymentprofile related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the payment profile for the current project.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.V2ProvisioningPaymentProfileParams) (*stripe.V2ProvisioningPaymentProfile, error) {
	paymentprofile := &stripe.V2ProvisioningPaymentProfile{}
	err := c.B.Call(
		http.MethodGet, "/v2/provisioning/payment_profile", c.Key, params, paymentprofile)
	return paymentprofile, err
}
