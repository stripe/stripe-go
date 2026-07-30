//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebitmandate provides the receiveddebitmandate related APIs
package receiveddebitmandate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke receiveddebitmandate related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebitMandate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementReceivedDebitMandateParams) (*stripe.V2MoneyManagementReceivedDebitMandate, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/received_debit_mandates/%s", id)
	receiveddebitmandate := &stripe.V2MoneyManagementReceivedDebitMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebitmandate)
	return receiveddebitmandate, err
}

// Cancels an active ReceivedDebitMandate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementReceivedDebitMandateCancelParams) (*stripe.V2MoneyManagementReceivedDebitMandate, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/received_debit_mandates/%s/cancel", id)
	receiveddebitmandate := &stripe.V2MoneyManagementReceivedDebitMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, receiveddebitmandate)
	return receiveddebitmandate, err
}

// Returns a list of ReceivedDebitMandates.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementReceivedDebitMandateListParams) stripe.Seq2[*stripe.V2MoneyManagementReceivedDebitMandate, error] {
	if listParams == nil {
		listParams = &stripe.V2MoneyManagementReceivedDebitMandateListParams{}
	}
	return stripe.NewV2List("/v2/money_management/received_debit_mandates", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementReceivedDebitMandate], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementReceivedDebitMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
