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
	path := FormatURLPath(
		"/v1/test_helpers/terminal/readers/%s/present_payment_method", id)
	reader := &TerminalReader{}
	if params == nil {
		params = &TestHelpersTerminalReaderPresentPaymentMethodParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, reader)
	return reader, err
}
