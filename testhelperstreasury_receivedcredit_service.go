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

// v1TestHelpersTreasuryReceivedCreditService is used to invoke /v1/treasury/received_credits APIs.
type v1TestHelpersTreasuryReceivedCreditService struct {
	B   Backend
	Key string
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
func (c v1TestHelpersTreasuryReceivedCreditService) Create(ctx context.Context, params *TestHelpersTreasuryReceivedCreditCreateParams) (*TreasuryReceivedCredit, error) {
	if params == nil {
		params = &TestHelpersTreasuryReceivedCreditCreateParams{}
	}
	params.Context = ctx
	receivedcredit := &TreasuryReceivedCredit{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/treasury/received_credits", c.Key, params, receivedcredit)
	return receivedcredit, err
}
