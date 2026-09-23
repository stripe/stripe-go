//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethodrequest provides the paymentmethodrequest related APIs
package paymentmethodrequest

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke paymentmethodrequest related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a request for a customer to authorize a new payment method.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2ProvisioningPaymentMethodRequestParams) (*stripe.V2ProvisioningPaymentMethodRequest, error) {
	paymentmethodrequest := &stripe.V2ProvisioningPaymentMethodRequest{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/payment_method_requests", c.Key, params, paymentmethodrequest)
	return paymentmethodrequest, err
}
