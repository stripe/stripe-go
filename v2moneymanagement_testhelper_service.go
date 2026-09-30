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

// v2MoneyManagementTestHelperService is used to invoke testhelper related APIs.
type v2MoneyManagementTestHelperService struct {
	B   Backend
	Key string
}

// Creates an EarnedCredit in a Sandbox environment for testing purposes.
func (c v2MoneyManagementTestHelperService) EarnedCredits(ctx context.Context, params *V2MoneyManagementTestHelperEarnedCreditsParams) (*V2MoneyManagementEarnedCreditSimulation, error) {
	if params == nil {
		params = &V2MoneyManagementTestHelperEarnedCreditsParams{}
	}
	params.Context = ctx
	earnedcreditsimulation := &V2MoneyManagementEarnedCreditSimulation{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/test_helpers/earned_credits", c.Key, params, earnedcreditsimulation)
	return earnedcreditsimulation, err
}
