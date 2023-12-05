//
//
// File generated from our OpenAPI spec
//
//

// Package reader provides the /terminal/readers APIs
package reader

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /terminal/readers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// PresentPaymentMethod is the method for the `POST /v1/test_helpers/terminal/readers/{reader}/present_payment_method` API.
func PresentPaymentMethod(id string, params *stripe.TestHelpersTerminalReaderPresentPaymentMethodParams) (*stripe.TerminalReader, error) {
	return getC().PresentPaymentMethod(id, params)
}

// PresentPaymentMethod is the method for the `POST /v1/test_helpers/terminal/readers/{reader}/present_payment_method` API.
func (c Client) PresentPaymentMethod(id string, params *stripe.TestHelpersTerminalReaderPresentPaymentMethodParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/present_payment_method",
		id,
	)
	reader := &stripe.TerminalReader{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, reader)
	return reader, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
