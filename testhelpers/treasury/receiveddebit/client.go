//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the /v1/treasury/received_debits APIs
package receiveddebit

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke /v1/treasury/received_debits APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
func New(params *stripe.TestHelpersTreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	return getC().New(params)
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
func (c Client) New(params *stripe.TestHelpersTreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	receiveddebit := &stripe.TreasuryReceivedDebit{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/treasury/received_debits", c.Key, params, receiveddebit)
	return receiveddebit, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
