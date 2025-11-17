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

// v1BalanceTransferService is used to invoke /v1/balance_transfers APIs.
type v1BalanceTransferService struct {
	B   Backend
	Key string
}

// Creates a balance transfer. For Issuing use cases, funds will be debited immediately from the source balance and credited to the destination balance immediately (if your account is based in the US) or next-business-day (if your account is based in the EU). For Segregated Separate Charges and Transfers use cases, funds will be debited immediately from the source balance and credited immediately to the destination balance.
func (c v1BalanceTransferService) Create(ctx context.Context, params *BalanceTransferCreateParams) (*BalanceTransfer, error) {
	if params == nil {
		params = &BalanceTransferCreateParams{}
	}
	params.Context = ctx
	balancetransfer := &BalanceTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v1/balance_transfers", c.Key, params, balancetransfer)
	return balancetransfer, err
}
