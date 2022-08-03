//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the /treasury/inbound_transfers APIs
package inboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /treasury/inbound_transfers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Fail is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/fail` API.
func Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Fail(id, params)
}

// Fail is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/fail` API.
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/fail",
		id,
	)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// ReturnInboundTransfer is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/return` API.
func ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().ReturnInboundTransfer(id, params)
}

// ReturnInboundTransfer is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/return` API.
func (c Client) ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/return",
		id,
	)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Succeed is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/succeed` API.
func Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Succeed(id, params)
}

// Succeed is the method for the `POST /v1/test_helpers/treasury/inbound_transfers/{id}/succeed` API.
func (c Client) Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/succeed",
		id,
	)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
