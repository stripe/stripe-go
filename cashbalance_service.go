//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"
)

// v1CashBalanceService is used to invoke /v1/customers/{id}/cash_balance APIs.
type v1CashBalanceService struct {
	B   Backend
	Key string
}

// Retrieves a customer's cash balance.
func (c v1CashBalanceService) Retrieve(ctx context.Context, params *CashBalanceRetrieveParams) (*CashBalance, error) {
	if params == nil {
		params = &CashBalanceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/cash_balance", StringValue(params.ID))
	cashbalance := &CashBalance{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cashbalance)
	return cashbalance, err
}

// Changes the settings on a customer's cash balance.
func (c v1CashBalanceService) Update(ctx context.Context, params *CashBalanceUpdateParams) (*CashBalance, error) {
	if params == nil || params.ID == nil {
		return nil, fmt.Errorf("params cannot be nil, and params.ID must be set")
	}
	if params == nil {
		params = &CashBalanceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/cash_balance", StringValue(params.ID))
	cashbalance := &CashBalance{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cashbalance)
	return cashbalance, err
}
