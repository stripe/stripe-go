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

// v2MoneyManagementFinancialAccountsWalletExportService is used to invoke walletexport related APIs.
type v2MoneyManagementFinancialAccountsWalletExportService struct {
	B   Backend
	Key string
}

// Retrieves the wallet export metadata for a closed FinancialAccount. Credentials are returned only by the export_credentials action.
func (c v2MoneyManagementFinancialAccountsWalletExportService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountsWalletExportRetrieveParams) (*V2MoneyManagementFinancialAccountWalletExport, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountsWalletExportRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/financial_accounts/%s/wallet_export", id)
	financialaccountwalletexport := &V2MoneyManagementFinancialAccountWalletExport{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, financialaccountwalletexport)
	return financialaccountwalletexport, err
}

// Exports wallet credentials encrypted to the supplied recipient key. The first successful request starts one fixed one-hour retrieval window; later requests may use a different recipient key without extending it.
func (c v2MoneyManagementFinancialAccountsWalletExportService) ExportCredentials(ctx context.Context, id string, params *V2MoneyManagementFinancialAccountsWalletExportExportCredentialsParams) (*V2MoneyManagementFinancialAccountWalletExportCredentials, error) {
	if params == nil {
		params = &V2MoneyManagementFinancialAccountsWalletExportExportCredentialsParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/financial_accounts/%s/wallet_export/export_credentials", id)
	financialaccountwalletexportcredentials := &V2MoneyManagementFinancialAccountWalletExportCredentials{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaccountwalletexportcredentials)
	return financialaccountwalletexportcredentials, err
}
