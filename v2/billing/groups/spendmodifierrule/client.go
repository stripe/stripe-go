//
//
// File generated from our OpenAPI spec
//
//

// Package spendmodifierrule provides the spendmodifierrule related APIs
package spendmodifierrule

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke spendmodifierrule related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a spend modifier associated with a billing group.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingGroupsSpendModifierRuleParams) (*stripe.V2BillingGroupSpendModifier, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/groups/%s/spend_modifier_rules/%s", stripe.StringValue(
			params.GroupID), id)
	groupspendmodifier := &stripe.V2BillingGroupSpendModifier{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, groupspendmodifier)
	return groupspendmodifier, err
}

// List spend modifiers associated with a billing group.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingGroupsSpendModifierRuleListParams) stripe.Seq2[*stripe.V2BillingGroupSpendModifier, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingGroupsSpendModifierRuleListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/billing/groups/%s/spend_modifier_rules", stripe.StringValue(
			listParams.GroupID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingGroupSpendModifier], error) {
		page := &stripe.V2Page[*stripe.V2BillingGroupSpendModifier]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
