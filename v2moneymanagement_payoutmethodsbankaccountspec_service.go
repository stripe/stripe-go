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

// v2MoneyManagementPayoutMethodsBankAccountSpecService is used to invoke payoutmethodsbankaccountspec related APIs.
type v2MoneyManagementPayoutMethodsBankAccountSpecService struct {
	B   Backend
	Key string
}

// Fetch the specifications for a set of countries to know which
// credential fields are required, the validations for each fields, and how to translate these
// country-specific fields to the generic fields in the PayoutMethodBankAccount type.
func (c v2MoneyManagementPayoutMethodsBankAccountSpecService) Retrieve(ctx context.Context, params *V2MoneyManagementPayoutMethodsBankAccountSpecRetrieveParams) (*V2MoneyManagementPayoutMethodsBankAccountSpec, error) {
	if params == nil {
		params = &V2MoneyManagementPayoutMethodsBankAccountSpecRetrieveParams{}
	}
	params.Context = ctx
	payoutmethodsbankaccountspec := &V2MoneyManagementPayoutMethodsBankAccountSpec{}
	err := c.B.Call(
		http.MethodGet, "/v2/money_management/payout_methods_bank_account_spec", c.Key, params, payoutmethodsbankaccountspec)
	return payoutmethodsbankaccountspec, err
}
