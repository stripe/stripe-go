//
//
// File generated from our OpenAPI spec
//
//

// Package dispute provides the dispute related APIs
package dispute

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke dispute related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a Dispute for the specified money movement.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementDisputeParams) (*stripe.V2MoneyManagementDispute, error) {
	dispute := &stripe.V2MoneyManagementDispute{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/disputes", c.Key, params, dispute)
	return dispute, err
}

// Retrieves a Dispute.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementDisputeParams) (*stripe.V2MoneyManagementDispute, error) {
	path := stripe.FormatURLPath("/v2/money_management/disputes/%s", id)
	dispute := &stripe.V2MoneyManagementDispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, dispute)
	return dispute, err
}

// Retrieves a list of Disputes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementDisputeListParams) stripe.Seq2[*stripe.V2MoneyManagementDispute, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementDisputeListParams{}
	}
	return stripe.NewV2List("/v2/money_management/disputes", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementDispute], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementDispute]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
