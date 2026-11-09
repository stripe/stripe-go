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

// v2TestHelpersMoneyManagementService is used to invoke moneymanagement related APIs.
type v2TestHelpersMoneyManagementService struct {
	B   Backend
	Key string
}

// Creates a RecipientVerification with a specified match result for testing purposes in a Sandbox environment.
func (c v2TestHelpersMoneyManagementService) RecipientVerifications(ctx context.Context, params *V2TestHelpersMoneyManagementRecipientVerificationsParams) (*V2MoneyManagementRecipientVerification, error) {
	if params == nil {
		params = &V2TestHelpersMoneyManagementRecipientVerificationsParams{}
	}
	params.Context = ctx
	recipientverification := &V2MoneyManagementRecipientVerification{}
	err := c.B.Call(
		http.MethodPost, "/v2/test_helpers/money_management/recipient_verifications", c.Key, params, recipientverification)
	return recipientverification, err
}
