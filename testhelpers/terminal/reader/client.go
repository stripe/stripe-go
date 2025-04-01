//
//
// File generated from our OpenAPI spec
//
//

// Package reader provides the /v1/terminal/readers APIs
package reader

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke /v1/terminal/readers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
func PresentPaymentMethod(id string, params *stripe.TestHelpersTerminalReaderPresentPaymentMethodParams) (*stripe.TerminalReader, error) {
	return getC().PresentPaymentMethod(id, params)
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
func (c Client) PresentPaymentMethod(id string, params *stripe.TestHelpersTerminalReaderPresentPaymentMethodParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/present_payment_method", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Use this endpoint to trigger a successful input collection on a simulated reader.
func SucceedInputCollection(id string, params *stripe.TestHelpersTerminalReaderSucceedInputCollectionParams) (*stripe.TerminalReader, error) {
	return getC().SucceedInputCollection(id, params)
}

// Use this endpoint to trigger a successful input collection on a simulated reader.
func (c Client) SucceedInputCollection(id string, params *stripe.TestHelpersTerminalReaderSucceedInputCollectionParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/succeed_input_collection", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Use this endpoint to complete an input collection with a timeout error on a simulated reader.
func TimeoutInputCollection(id string, params *stripe.TestHelpersTerminalReaderTimeoutInputCollectionParams) (*stripe.TerminalReader, error) {
	return getC().TimeoutInputCollection(id, params)
}

// Use this endpoint to complete an input collection with a timeout error on a simulated reader.
func (c Client) TimeoutInputCollection(id string, params *stripe.TestHelpersTerminalReaderTimeoutInputCollectionParams) (*stripe.TerminalReader, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/timeout_input_collection", id)
	reader := &stripe.TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
