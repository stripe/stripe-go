//
//
// File generated from our OpenAPI spec
//
//

// Package collectionsetting provides the collectionsetting related APIs
package collectionsetting

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke collectionsetting related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a CollectionSetting object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingCollectionSettingParams) (*stripe.V2BillingCollectionSetting, error) {
	collectionsetting := &stripe.V2BillingCollectionSetting{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/collection_settings", c.Key, params, collectionsetting)
	return collectionsetting, err
}

// Retrieve a CollectionSetting by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingCollectionSettingParams) (*stripe.V2BillingCollectionSetting, error) {
	path := stripe.FormatURLPath("/v2/billing/collection_settings/%s", id)
	collectionsetting := &stripe.V2BillingCollectionSetting{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, collectionsetting)
	return collectionsetting, err
}

// Update fields on an existing CollectionSetting.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingCollectionSettingParams) (*stripe.V2BillingCollectionSetting, error) {
	path := stripe.FormatURLPath("/v2/billing/collection_settings/%s", id)
	collectionsetting := &stripe.V2BillingCollectionSetting{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, collectionsetting)
	return collectionsetting, err
}

// List all CollectionSetting objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingCollectionSettingListParams) stripe.Seq2[*stripe.V2BillingCollectionSetting, error] {
	return stripe.NewV2List("/v2/billing/collection_settings", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingCollectionSetting], error) {
		page := &stripe.V2Page[*stripe.V2BillingCollectionSetting]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
