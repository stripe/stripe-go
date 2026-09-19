//
//
// File generated from our OpenAPI spec
//
//

// Package walletexport provides the walletexport related APIs
package walletexport

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke walletexport related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the wallet export metadata for a closed FinancialAccount. Credentials are returned only by the export_credentials action.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementFinancialAccountsWalletExportParams) (*stripe.V2MoneyManagementFinancialAccountWalletExport, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/financial_accounts/%s/wallet_export", id)
	financialaccountwalletexport := &stripe.V2MoneyManagementFinancialAccountWalletExport{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, financialaccountwalletexport)
	return financialaccountwalletexport, err
}

// Exports wallet credentials encrypted to the supplied recipient key. The first successful request starts one fixed one-hour retrieval window; later requests may use a different recipient key without extending it.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ExportCredentials(id string, params *stripe.V2MoneyManagementFinancialAccountsWalletExportExportCredentialsParams) (*stripe.V2MoneyManagementFinancialAccountWalletExportCredentials, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/financial_accounts/%s/wallet_export/export_credentials", id)
	financialaccountwalletexportcredentials := &stripe.V2MoneyManagementFinancialAccountWalletExportCredentials{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaccountwalletexportcredentials)
	return financialaccountwalletexportcredentials, err
}
