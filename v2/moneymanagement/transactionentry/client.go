//
//
// File generated from our OpenAPI spec
//
//

// Package transactionentry provides the transactionentry related APIs
package transactionentry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke transactionentry related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a TransactionEntry by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementTransactionEntryParams) (*stripe.V2MoneyManagementTransactionEntry, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/transaction_entries/%s", id)
	transactionentry := &stripe.V2MoneyManagementTransactionEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Returns a list of TransactionEntries that match the provided filters.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementTransactionEntryListParams) stripe.Seq2[*stripe.V2MoneyManagementTransactionEntry, error] {
	return stripe.NewV2List("/v2/money_management/transaction_entries", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementTransactionEntry], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementTransactionEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
