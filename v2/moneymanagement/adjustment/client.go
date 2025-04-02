//
//
// File generated from our OpenAPI spec
//
//

// Package adjustment provides the adjustment related APIs
package adjustment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke adjustment related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an Adjustment by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementAdjustmentParams) (*stripe.V2MoneyManagementAdjustment, error) {
	path := stripe.FormatURLPath("/v2/money_management/adjustments/%s", id)
	adjustment := &stripe.V2MoneyManagementAdjustment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, adjustment)
	return adjustment, err
}

// Returns a list of Adjustments that match the provided filters.
func (c Client) All(listParams *stripe.V2MoneyManagementAdjustmentListParams) stripe.Seq2[stripe.V2MoneyManagementAdjustment, error] {
	return stripe.NewV2List("/v2/money_management/adjustments", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementAdjustment], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementAdjustment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
