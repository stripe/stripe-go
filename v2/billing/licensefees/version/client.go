//
//
// File generated from our OpenAPI spec
//
//

// Package version provides the version related APIs
package version

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke version related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a LicenseFeeVersion object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingLicenseFeesVersionParams) (*stripe.V2BillingLicenseFeeVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/license_fees/%s/versions/%s", stripe.StringValue(
			params.LicenseFeeID), id)
	licensefeeversion := &stripe.V2BillingLicenseFeeVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefeeversion)
	return licensefeeversion, err
}

// List all versions of a LicenseFee objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingLicenseFeesVersionListParams) stripe.Seq2[*stripe.V2BillingLicenseFeeVersion, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/license_fees/%s/versions", stripe.StringValue(
			listParams.LicenseFeeID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingLicenseFeeVersion], error) {
		page := &stripe.V2Page[*stripe.V2BillingLicenseFeeVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
