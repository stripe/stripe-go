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

// v1BalanceService is used to invoke /v1/balance APIs.
type v1BalanceService struct {
	B   Backend
	Key string
}

// Retrieves the current account balance, based on the authentication that was used to make the request.
//
//	For a sample request, see [Accounting for negative balances](https://docs.stripe.com/docs/connect/account-balances#accounting-for-negative-balances).
func (c v1BalanceService) Retrieve(ctx context.Context, params *BalanceRetrieveParams) (*Balance, error) {
	if params == nil {
		params = &BalanceRetrieveParams{}
	}
	params.Context = ctx
	balance := &Balance{}
	err := c.B.Call(http.MethodGet, "/v1/balance", c.Key, params, balance)
	return balance, err
}
