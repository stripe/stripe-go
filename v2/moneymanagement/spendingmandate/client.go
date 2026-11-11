//
//
// File generated from our OpenAPI spec
//
//

// Package spendingmandate provides the spendingmandate related APIs
package spendingmandate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke spendingmandate related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing SpendingMandate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementSpendingMandateParams) (*stripe.V2MoneyManagementSpendingMandate, error) {
	path := stripe.FormatURLPath("/v2/money_management/spending_mandates/%s", id)
	spendingmandate := &stripe.V2MoneyManagementSpendingMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, spendingmandate)
	return spendingmandate, err
}

// Cancels an active SpendingMandate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementSpendingMandateCancelParams) (*stripe.V2MoneyManagementSpendingMandate, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/spending_mandates/%s/cancel", id)
	spendingmandate := &stripe.V2MoneyManagementSpendingMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, spendingmandate)
	return spendingmandate, err
}

// Returns a list of SpendingMandates.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementSpendingMandateListParams) stripe.Seq2[*stripe.V2MoneyManagementSpendingMandate, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementSpendingMandateListParams{}
	}
	return stripe.NewV2List("/v2/money_management/spending_mandates", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementSpendingMandate], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementSpendingMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
