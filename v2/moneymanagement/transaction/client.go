//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the transaction related APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke transaction related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Transaction by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementTransactionParams) (*stripe.V2MoneyManagementTransaction, error) {
	path := stripe.FormatURLPath("/v2/money_management/transactions/%s", id)
	transaction := &stripe.V2MoneyManagementTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Returns a list of Transactions that match the provided filters.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementTransactionListParams) stripe.Seq2[*stripe.V2MoneyManagementTransaction, error] {
	return stripe.NewV2List("/v2/money_management/transactions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementTransaction], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementTransaction]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
