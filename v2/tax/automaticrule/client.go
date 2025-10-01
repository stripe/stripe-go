//
//
// File generated from our OpenAPI spec
//
//

// Package automaticrule provides the automaticrule related APIs
package automaticrule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke automaticrule related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an AutomaticRule object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2TaxAutomaticRuleParams) (*stripe.V2TaxAutomaticRule, error) {
	automaticrule := &stripe.V2TaxAutomaticRule{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/automatic_rules", c.Key, params, automaticrule)
	return automaticrule, err
}

// Retrieves an AutomaticRule object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2TaxAutomaticRuleParams) (*stripe.V2TaxAutomaticRule, error) {
	path := stripe.FormatURLPath("/v2/tax/automatic_rules/%s", id)
	automaticrule := &stripe.V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Updates the automatic Tax configuration for an AutomaticRule object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2TaxAutomaticRuleParams) (*stripe.V2TaxAutomaticRule, error) {
	path := stripe.FormatURLPath("/v2/tax/automatic_rules/%s", id)
	automaticrule := &stripe.V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Deactivates an AutomaticRule object. Deactivated AutomaticRule objects are ignored in future tax calculations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.V2TaxAutomaticRuleDeactivateParams) (*stripe.V2TaxAutomaticRule, error) {
	path := stripe.FormatURLPath("/v2/tax/automatic_rules/%s/deactivate", id)
	automaticrule := &stripe.V2TaxAutomaticRule{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, automaticrule)
	return automaticrule, err
}

// Finds an AutomaticRule object by BillableItem ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Find(params *stripe.V2TaxAutomaticRuleFindParams) (*stripe.V2TaxAutomaticRule, error) {
	automaticrule := &stripe.V2TaxAutomaticRule{}
	err := c.B.Call(
		http.MethodGet, "/v2/tax/automatic_rules/find", c.Key, params, automaticrule)
	return automaticrule, err
}
