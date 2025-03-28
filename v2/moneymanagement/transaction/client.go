//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the transaction related APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke transaction related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Transaction by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementTransactionParams) (*stripe.V2MoneyManagementTransaction, error) {
	path := stripe.FormatURLPath("/v2/money_management/transactions/%s", id)
	transaction := &stripe.V2MoneyManagementTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, transaction)
	return transaction, err
}

// Returns a list of Transactions that match the provided filters.
func (c Client) All(listParams *stripe.V2MoneyManagementTransactionListParams) stripe.Seq2[stripe.V2MoneyManagementTransaction, error] {
	return stripe.NewV2List("/v2/money_management/transactions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementTransaction], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementTransaction]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
