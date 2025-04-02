//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the receivedcredit related APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke receivedcredit related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a ReceivedCredit by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementReceivedCreditParams) (*stripe.V2MoneyManagementReceivedCredit, error) {
	path := stripe.FormatURLPath("/v2/money_management/received_credits/%s", id)
	receivedcredit := &stripe.V2MoneyManagementReceivedCredit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receivedcredit)
	return receivedcredit, err
}

// Retrieves a list of ReceivedCredits.
func (c Client) All(listParams *stripe.V2MoneyManagementReceivedCreditListParams) stripe.Seq2[stripe.V2MoneyManagementReceivedCredit, error] {
	return stripe.NewV2List("/v2/money_management/received_credits", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementReceivedCredit], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementReceivedCredit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
