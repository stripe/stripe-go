//
//
// File generated from our OpenAPI spec
//
//

// Package spendmodifierrule provides the spendmodifierrule related APIs
package spendmodifierrule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke spendmodifierrule related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a Spend Modifier associated with a Billing Cadence.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingCadencesSpendModifierRuleParams) (*stripe.V2BillingCadenceSpendModifier, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/cadences/%s/spend_modifier_rules/%s", stripe.StringValue(
			params.CadenceID), id)
	cadencespendmodifier := &stripe.V2BillingCadenceSpendModifier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cadencespendmodifier)
	return cadencespendmodifier, err
}

// List all Spend Modifiers associated with a Billing Cadence.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingCadencesSpendModifierRuleListParams) stripe.Seq2[*stripe.V2BillingCadenceSpendModifier, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingCadencesSpendModifierRuleListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/billing/cadences/%s/spend_modifier_rules", stripe.StringValue(
			listParams.CadenceID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingCadenceSpendModifier], error) {
		page := &stripe.V2Page[*stripe.V2BillingCadenceSpendModifier]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
