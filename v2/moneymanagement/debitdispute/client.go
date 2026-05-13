//
//
// File generated from our OpenAPI spec
//
//

// Package debitdispute provides the debitdispute related APIs
package debitdispute

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke debitdispute related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new DebitDispute for a ReceivedDebit.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementDebitDisputeParams) (*stripe.V2MoneyManagementDebitDispute, error) {
	debitdispute := &stripe.V2MoneyManagementDebitDispute{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/debit_disputes", c.Key, params, debitdispute)
	return debitdispute, err
}

// Retrieves the details of an existing DebitDispute.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementDebitDisputeParams) (*stripe.V2MoneyManagementDebitDispute, error) {
	path := stripe.FormatURLPath("/v2/money_management/debit_disputes/%s", id)
	debitdispute := &stripe.V2MoneyManagementDebitDispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, debitdispute)
	return debitdispute, err
}

// Retrieves a list of DebitDisputes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementDebitDisputeListParams) stripe.Seq2[*stripe.V2MoneyManagementDebitDispute, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementDebitDisputeListParams{}
	}
	return stripe.NewV2List("/v2/money_management/debit_disputes", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementDebitDispute], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementDebitDispute]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
