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

// v1TestHelpersCustomerService is used to invoke /v1/customers APIs.
type v1TestHelpersCustomerService struct {
	B   Backend
	Key string
}

// Create an incoming testmode bank transfer
func (c v1TestHelpersCustomerService) FundCashBalance(ctx context.Context, id string, params *TestHelpersCustomerFundCashBalanceParams) (*CustomerCashBalanceTransaction, error) {
	if params == nil {
		params = &TestHelpersCustomerFundCashBalanceParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/customers/%s/fund_cash_balance", id)
	customercashbalancetransaction := &CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, customercashbalancetransaction)
	return customercashbalancetransaction, err
}
