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

// v2MoneyManagementFinancialAddressService is used to invoke financialaddress related APIs.
type v2MoneyManagementFinancialAddressService struct {
	B   Backend
	Key string
}

// Create a new FinancialAddress for a FinancialAccount.
func (c v2MoneyManagementFinancialAddressService) Create(ctx context.Context, params *V2MoneyManagementFinancialAddressCreateParams) (*V2MoneyManagementFinancialAddress, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAddressCreateParams{}
	}
	params.Context = ctx
	financialaddress := &V2MoneyManagementFinancialAddress{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/financial_addresses", c.Key, params, financialaddress)
	return financialaddress, err
}

// Retrieve a FinancialAddress. By default, the FinancialAddress will be returned in its unexpanded state, revealing only the last 4 digits of the account number.
func (c v2MoneyManagementFinancialAddressService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementFinancialAddressRetrieveParams) (*V2MoneyManagementFinancialAddress, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAddressRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/financial_addresses/%s", id)
	financialaddress := &V2MoneyManagementFinancialAddress{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaddress)
	return financialaddress, err
}

// List all FinancialAddresses for a FinancialAccount.
func (c v2MoneyManagementFinancialAddressService) List(ctx context.Context, listParams *V2MoneyManagementFinancialAddressListParams) Seq2[*V2MoneyManagementFinancialAddress, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementFinancialAddressListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/financial_addresses", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementFinancialAddress], error) {
		page := &V2Page[*V2MoneyManagementFinancialAddress]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
