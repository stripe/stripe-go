//
//
// File generated from our OpenAPI spec
//
//

// Package version provides the version related APIs
package version

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke version related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a CollectionSetting Version by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingCollectionSettingsVersionParams) (*stripe.V2BillingCollectionSettingVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/collection_settings/%s/versions/%s", stripe.StringValue(
			params.CollectionSettingID), id)
	collectionsettingversion := &stripe.V2BillingCollectionSettingVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, collectionsettingversion)
	return collectionsettingversion, err
}

// List all CollectionSettingVersions by CollectionSetting ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingCollectionSettingsVersionListParams) stripe.Seq2[*stripe.V2BillingCollectionSettingVersion, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/collection_settings/%s/versions", stripe.StringValue(
			listParams.CollectionSettingID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingCollectionSettingVersion], error) {
		page := &stripe.V2Page[*stripe.V2BillingCollectionSettingVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
