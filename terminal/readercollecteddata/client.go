//
//
// File generated from our OpenAPI spec
//
//

// Package readercollecteddata provides the /v1/terminal/reader_collected_data APIs
package readercollecteddata

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/terminal/reader_collected_data APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve data collected using Reader hardware.
func Get(id string, params *stripe.TerminalReaderCollectedDataParams) (*stripe.TerminalReaderCollectedData, error) {
	return getC().Get(id, params)
}

// Retrieve data collected using Reader hardware.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TerminalReaderCollectedDataParams) (*stripe.TerminalReaderCollectedData, error) {
	path := stripe.FormatURLPath("/v1/terminal/reader_collected_data/%s", id)
	readercollecteddata := &stripe.TerminalReaderCollectedData{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, readercollecteddata)
	return readercollecteddata, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
