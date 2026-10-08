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

// v2MoneyManagementFundingSessionService is used to invoke fundingsession related APIs.
type v2MoneyManagementFundingSessionService struct {
	B   Backend
	Key string
}

// Create a FundingSession: a hosted funding surface for a customer to fund a FinancialAccount.
func (c v2MoneyManagementFundingSessionService) Create(ctx context.Context, params *V2MoneyManagementFundingSessionCreateParams) (*V2MoneyManagementFundingSession, error) {
	if params == nil {
		params = &V2MoneyManagementFundingSessionCreateParams{}
	}
	params.Context = ctx
	fundingsession := &V2MoneyManagementFundingSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/funding_sessions", c.Key, params, fundingsession)
	return fundingsession, err
}
