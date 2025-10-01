//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the receiveddebit related APIs
package receiveddebit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke receiveddebit related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a single ReceivedDebit by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementReceivedDebitParams) (*stripe.V2MoneyManagementReceivedDebit, error) {
	path := stripe.FormatURLPath("/v2/money_management/received_debits/%s", id)
	receiveddebit := &stripe.V2MoneyManagementReceivedDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebit)
	return receiveddebit, err
}

// Retrieves a list of ReceivedDebits, given the selected filters.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementReceivedDebitListParams) stripe.Seq2[*stripe.V2MoneyManagementReceivedDebit, error] {
	return stripe.NewV2List("/v2/money_management/received_debits", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementReceivedDebit], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementReceivedDebit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
