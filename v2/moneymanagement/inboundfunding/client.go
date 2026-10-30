//
//
// File generated from our OpenAPI spec
//
//

// Package inboundfunding provides the inboundfunding related APIs
package inboundfunding

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke inboundfunding related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an InboundFunding by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementInboundFundingParams) (*stripe.V2MoneyManagementInboundFunding, error) {
	path := stripe.FormatURLPath("/v2/money_management/inbound_fundings/%s", id)
	inboundfunding := &stripe.V2MoneyManagementInboundFunding{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundfunding)
	return inboundfunding, err
}

// Lists InboundFundings.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementInboundFundingListParams) stripe.Seq2[*stripe.V2MoneyManagementInboundFunding, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementInboundFundingListParams{}
	}
	return stripe.NewV2List("/v2/money_management/inbound_fundings", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementInboundFunding], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementInboundFunding]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
