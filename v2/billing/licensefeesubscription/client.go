//
//
// File generated from our OpenAPI spec
//
//

// Package licensefeesubscription provides the licensefeesubscription related APIs
package licensefeesubscription

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke licensefeesubscription related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a License Fee Subscription object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingLicenseFeeSubscriptionParams) (*stripe.V2BillingLicenseFeeSubscription, error) {
	path := stripe.FormatURLPath("/v2/billing/license_fee_subscriptions/%s", id)
	licensefeesubscription := &stripe.V2BillingLicenseFeeSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefeesubscription)
	return licensefeesubscription, err
}
