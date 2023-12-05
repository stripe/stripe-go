//
//
// File generated from our OpenAPI spec
//
//

// Package receivedcredit provides the /treasury/received_credits APIs
package receivedcredit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /treasury/received_credits APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury received credit.
func New(params *stripe.TestHelpersTreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	return getC().New(params)
}

// New creates a new treasury received credit.
func (c Client) New(params *stripe.TestHelpersTreasuryReceivedCreditParams) (*stripe.TreasuryReceivedCredit, error) {
	receivedcredit := &stripe.TreasuryReceivedCredit{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/test_helpers/treasury/received_credits", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, receivedcredit)
	return receivedcredit, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
