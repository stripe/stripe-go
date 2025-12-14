//
//
// File generated from our OpenAPI spec
//
//

// Package manualrule provides the manualrule related APIs
package manualrule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke manualrule related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a ManualRule object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2TaxManualRuleParams) (*stripe.V2TaxManualRule, error) {
	manualrule := &stripe.V2TaxManualRule{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/manual_rules", c.Key, params, manualrule)
	return manualrule, err
}

// Retrieves a ManualRule object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2TaxManualRuleParams) (*stripe.V2TaxManualRule, error) {
	path := stripe.FormatURLPath("/v2/tax/manual_rules/%s", id)
	manualrule := &stripe.V2TaxManualRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, manualrule)
	return manualrule, err
}

// Updates the Tax configuration for a ManualRule object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2TaxManualRuleParams) (*stripe.V2TaxManualRule, error) {
	path := stripe.FormatURLPath("/v2/tax/manual_rules/%s", id)
	manualrule := &stripe.V2TaxManualRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, manualrule)
	return manualrule, err
}

// Deactivates a ManualRule object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.V2TaxManualRuleDeactivateParams) (*stripe.V2TaxManualRule, error) {
	path := stripe.FormatURLPath("/v2/tax/manual_rules/%s/deactivate", id)
	manualrule := &stripe.V2TaxManualRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, manualrule)
	return manualrule, err
}

// List all ManualRule objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2TaxManualRuleListParams) stripe.Seq2[*stripe.V2TaxManualRule, error] {
	return stripe.NewV2List("/v2/tax/manual_rules", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2TaxManualRule], error) {
		page := &stripe.V2Page[*stripe.V2TaxManualRule]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
