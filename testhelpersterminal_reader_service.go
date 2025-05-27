//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TestHelpersTerminalReaderService is used to invoke /v1/terminal/readers APIs.
type v1TestHelpersTerminalReaderService struct {
	B   Backend
	Key string
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
func (c v1TestHelpersTerminalReaderService) PresentPaymentMethod(ctx context.Context, id string, params *TestHelpersTerminalReaderPresentPaymentMethodParams) (*TerminalReader, error) {
	if params == nil {
		params = &TestHelpersTerminalReaderPresentPaymentMethodParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/present_payment_method", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Use this endpoint to trigger a successful input collection on a simulated reader.
func (c v1TestHelpersTerminalReaderService) SucceedInputCollection(ctx context.Context, id string, params *TestHelpersTerminalReaderSucceedInputCollectionParams) (*TerminalReader, error) {
	if params == nil {
		params = &TestHelpersTerminalReaderSucceedInputCollectionParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/succeed_input_collection", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}

// Use this endpoint to complete an input collection with a timeout error on a simulated reader.
func (c v1TestHelpersTerminalReaderService) TimeoutInputCollection(ctx context.Context, id string, params *TestHelpersTerminalReaderTimeoutInputCollectionParams) (*TerminalReader, error) {
	if params == nil {
		params = &TestHelpersTerminalReaderTimeoutInputCollectionParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/timeout_input_collection", id)
	reader := &TerminalReader{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}
