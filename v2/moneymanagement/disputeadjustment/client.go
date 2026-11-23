//
//
// File generated from our OpenAPI spec
//
//

// Package disputeadjustment provides the disputeadjustment related APIs
package disputeadjustment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke disputeadjustment related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a DisputeAdjustment by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementDisputeAdjustmentParams) (*stripe.V2MoneyManagementDisputeAdjustment, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/dispute_adjustments/%s", id)
	disputeadjustment := &stripe.V2MoneyManagementDisputeAdjustment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, disputeadjustment)
	return disputeadjustment, err
}

// Retrieves a list of DisputeAdjustments.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementDisputeAdjustmentListParams) stripe.Seq2[*stripe.V2MoneyManagementDisputeAdjustment, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementDisputeAdjustmentListParams{}
	}
	return stripe.NewV2List("/v2/money_management/dispute_adjustments", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementDisputeAdjustment], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementDisputeAdjustment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
