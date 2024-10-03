//
//
// File generated from our OpenAPI spec
//
//

// Package readercollecteddata provides the /terminal/reader_collected_data APIs
package readercollecteddata

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v80"
)

// Client is used to invoke /terminal/reader_collected_data APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve data collected using Reader hardware.
func Get(id string, params *stripe.TerminalReaderCollectedDataParams) (*stripe.TerminalReaderCollectedData, error) {
	return getC().Get(id, params)
}

// Retrieve data collected using Reader hardware.
func (c Client) Get(id string, params *stripe.TerminalReaderCollectedDataParams) (*stripe.TerminalReaderCollectedData, error) {
	path := stripe.FormatURLPath("/v1/terminal/reader_collected_data/%s", id)
	readercollecteddata := &stripe.TerminalReaderCollectedData{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, readercollecteddata)
	return readercollecteddata, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
