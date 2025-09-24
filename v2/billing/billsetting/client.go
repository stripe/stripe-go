//
//
// File generated from our OpenAPI spec
//
//

// Package billsetting provides the billsetting related APIs
package billsetting

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke billsetting related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a BillSetting object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingBillSettingParams) (*stripe.V2BillingBillSetting, error) {
	billsetting := &stripe.V2BillingBillSetting{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/bill_settings", c.Key, params, billsetting)
	return billsetting, err
}

// Retrieve a BillSetting object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingBillSettingParams) (*stripe.V2BillingBillSetting, error) {
	path := stripe.FormatURLPath("/v2/billing/bill_settings/%s", id)
	billsetting := &stripe.V2BillingBillSetting{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, billsetting)
	return billsetting, err
}

// Update fields on an existing BillSetting object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingBillSettingParams) (*stripe.V2BillingBillSetting, error) {
	path := stripe.FormatURLPath("/v2/billing/bill_settings/%s", id)
	billsetting := &stripe.V2BillingBillSetting{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, billsetting)
	return billsetting, err
}

// List all BillSetting objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingBillSettingListParams) stripe.Seq2[*stripe.V2BillingBillSetting, error] {
	return stripe.NewV2List("/v2/billing/bill_settings", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingBillSetting], error) {
		page := &stripe.V2Page[*stripe.V2BillingBillSetting]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
