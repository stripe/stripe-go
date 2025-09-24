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

// Retrieve a BillSettingVersion by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingBillSettingsVersionParams) (*stripe.V2BillingBillSettingVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/bill_settings/%s/versions/%s", stripe.StringValue(
			params.BillSettingID), id)
	billsettingversion := &stripe.V2BillingBillSettingVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, billsettingversion)
	return billsettingversion, err
}

// List all BillSettingVersions by BillSetting ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingBillSettingsVersionListParams) stripe.Seq2[*stripe.V2BillingBillSettingVersion, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/bill_settings/%s/versions", stripe.StringValue(
			listParams.BillSettingID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingBillSettingVersion], error) {
		page := &stripe.V2Page[*stripe.V2BillingBillSettingVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
