//
//
// File generated from our OpenAPI spec
//
//

// Package paymentlocationcapability provides the paymentlocationcapability related APIs
package paymentlocationcapability

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke paymentlocationcapability related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates a specified Payment Location Capability. Request or remove a payment location capability by updating its requested parameter.
func Update(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	return getC().Update(id, params)
}

// Updates a specified Payment Location Capability. Request or remove a payment location capability by updating its requested parameter.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	path := stripe.FormatURLPath("/v1/payment_location_capabilities/%s", id)
	paymentlocationcapability := &stripe.PaymentLocationCapability{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, paymentlocationcapability)
	return paymentlocationcapability, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
