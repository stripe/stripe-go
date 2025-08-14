//
//
// File generated from our OpenAPI spec
//
//

// Package licensefee provides the licensefee related APIs
package licensefee

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke licensefee related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a LicenseFee object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingLicenseFeeParams) (*stripe.V2BillingLicenseFee, error) {
	licensefee := &stripe.V2BillingLicenseFee{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/license_fees", c.Key, params, licensefee)
	return licensefee, err
}

// Retrieve a LicenseFee object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingLicenseFeeParams) (*stripe.V2BillingLicenseFee, error) {
	path := stripe.FormatURLPath("/v2/billing/license_fees/%s", id)
	licensefee := &stripe.V2BillingLicenseFee{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefee)
	return licensefee, err
}

// Update a LicenseFee object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingLicenseFeeParams) (*stripe.V2BillingLicenseFee, error) {
	path := stripe.FormatURLPath("/v2/billing/license_fees/%s", id)
	licensefee := &stripe.V2BillingLicenseFee{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, licensefee)
	return licensefee, err
}

// List all LicenseFee objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingLicenseFeeListParams) stripe.Seq2[*stripe.V2BillingLicenseFee, error] {
	return stripe.NewV2List("/v2/billing/license_fees", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingLicenseFee], error) {
		page := &stripe.V2Page[*stripe.V2BillingLicenseFee]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
