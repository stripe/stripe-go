//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the /v1/treasury/received_credits APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/treasury/received_credits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
func New(params *stripe.TestHelpersTreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	return getC().New(params)
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TestHelpersTreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	receivedcredit := &stripe.TreasuryReceivedCredit{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/treasury/received_credits", c.Key, params, receivedcredit)
	return receivedcredit, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
