//
//
// File generated from our OpenAPI spec
//
//

// Package transactionentry provides the transactionentry related APIs
package transactionentry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke transactionentry related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a TransactionEntry by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementTransactionEntryParams) (*stripe.V2MoneyManagementTransactionEntry, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/transaction_entries/%s", id)
	transactionentry := &stripe.V2MoneyManagementTransactionEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transactionentry)
	return transactionentry, err
}

// Returns a list of TransactionEntries that match the provided filters.
func (c Client) All(listParams *stripe.V2MoneyManagementTransactionEntryListParams) stripe.Seq2[stripe.V2MoneyManagementTransactionEntry, error] {
	return stripe.NewV2List("/v2/money_management/transaction_entries", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementTransactionEntry], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementTransactionEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
