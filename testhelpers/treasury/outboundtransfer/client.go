//
//
// File generated from our OpenAPI spec
//
//

// Package outboundtransfer provides the /treasury/outbound_transfers APIs
package outboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /treasury/outbound_transfers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Fail is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/fail` API.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Fail(id, params)
}

// Fail is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/fail` API.
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/fail",
		id,
	)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Post is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/post` API.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Post(id, params)
}

// Post is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/post` API.
func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/post",
		id,
	)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// ReturnOutboundTransfer is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/return` API.
func ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().ReturnOutboundTransfer(id, params)
}

// ReturnOutboundTransfer is the method for the `POST /v1/test_helpers/treasury/outbound_transfers/{outbound_transfer}/return` API.
func (c Client) ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/return",
		id,
	)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
