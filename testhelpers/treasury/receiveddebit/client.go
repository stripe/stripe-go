//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the /treasury/received_debits APIs
package receiveddebit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /treasury/received_debits APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury received debit.
func New(params *stripe.TestHelpersTreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	return getC().New(params)
}

// New creates a new treasury received debit.
func (c Client) New(params *stripe.TestHelpersTreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	receiveddebit := &stripe.TreasuryReceivedDebit{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/test_helpers/treasury/received_debits",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, receiveddebit)
	return receiveddebit, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
