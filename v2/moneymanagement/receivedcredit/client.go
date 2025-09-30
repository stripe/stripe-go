//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the receivedcredit related APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke receivedcredit related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a ReceivedCredit by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementReceivedCreditParams) (*stripe.V2MoneyManagementReceivedCredit, error) {
	path := stripe.FormatURLPath("/v2/money_management/received_credits/%s", id)
	receivedcredit := &stripe.V2MoneyManagementReceivedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Retrieves a list of ReceivedCredits.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2MoneyManagementReceivedCreditListParams) stripe.Seq2[*stripe.V2MoneyManagementReceivedCredit, error] {
	return stripe.NewV2List("/v2/money_management/received_credits", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementReceivedCredit], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementReceivedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
