//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the receiveddebit related APIs
package receiveddebit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke receiveddebit related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a single ReceivedDebit by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementReceivedDebitParams) (*stripe.V2MoneyManagementReceivedDebit, error) {
	path := stripe.FormatURLPath("/v2/money_management/received_debits/%s", id)
	receiveddebit := &stripe.V2MoneyManagementReceivedDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebit)
	return receiveddebit, err
}

// Retrieves a list of ReceivedDebits, given the selected filters.
func (c Client) All(listParams *stripe.V2MoneyManagementReceivedDebitListParams) stripe.Seq2[stripe.V2MoneyManagementReceivedDebit, error] {
	return stripe.NewV2List("/v2/money_management/received_debits", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementReceivedDebit], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementReceivedDebit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
