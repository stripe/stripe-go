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

// v1TestHelpersTreasuryReceivedDebitService is used to invoke /v1/treasury/received_debits APIs.
type v1TestHelpersTreasuryReceivedDebitService struct {
	B   Backend
	Key string
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
func (c v1TestHelpersTreasuryReceivedDebitService) Create(ctx context.Context, params *TestHelpersTreasuryReceivedDebitCreateParams) (*TreasuryReceivedDebit, error) {
	if params == nil {
		params = &TestHelpersTreasuryReceivedDebitCreateParams{}
	}
	params.Context = ctx
	receiveddebit := &TreasuryReceivedDebit{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/treasury/received_debits", c.Key, params, receiveddebit)
	return receiveddebit, err
}
