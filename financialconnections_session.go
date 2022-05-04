//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of account holder that this account belongs to.
type FinancialConnectionsSessionAccountHolderType string

// List of values that FinancialConnectionsSessionAccountHolderType can take
const (
	FinancialConnectionsSessionAccountHolderTypeAccount  FinancialConnectionsSessionAccountHolderType = "account"
	FinancialConnectionsSessionAccountHolderTypeCustomer FinancialConnectionsSessionAccountHolderType = "customer"
)

// Permissions requested for accounts collected during this session.
type FinancialConnectionsSessionPermission string

// List of values that FinancialConnectionsSessionPermission can take
const (
	FinancialConnectionsSessionPermissionBalances      FinancialConnectionsSessionPermission = "balances"
	FinancialConnectionsSessionPermissionOwnership     FinancialConnectionsSessionPermission = "ownership"
	FinancialConnectionsSessionPermissionPaymentMethod FinancialConnectionsSessionPermission = "payment_method"
	FinancialConnectionsSessionPermissionTransactions  FinancialConnectionsSessionPermission = "transactions"
)

// The account holder to link accounts for.
type FinancialConnectionsSessionAccountHolderParams struct {
	// The ID of the Stripe account whose accounts will be retrieved. Should only be present if `type` is `account`.
	Account *string `form:"account"`
	// The ID of the Stripe customer whose accounts will be retrieved. Should only be present if `type` is `customer`.
	Customer *string `form:"customer"`
	// Type of account holder to collect accounts for.
	Type *string `form:"type"`
}

// Filters to restrict the kinds of accounts to collect.
type FinancialConnectionsSessionFiltersParams struct {
	// List of countries from which to collect accounts.
	Countries []*string `form:"countries"`
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
type FinancialConnectionsSessionParams struct {
	Params `form:"*"`
	// The account holder to link accounts for.
	AccountHolder *FinancialConnectionsSessionAccountHolderParams `form:"account_holder"`
	// Filters to restrict the kinds of accounts to collect.
	Filters *FinancialConnectionsSessionFiltersParams `form:"filters"`
	// List of data features that you would like to request access to.
	//
	// Possible values are `balances`, `transactions`, `ownership`, and `payment_method`.
	Permissions []*string `form:"permissions"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url"`
}

// The account holder for whom accounts are collected in this session.
type FinancialConnectionsSessionAccountHolder struct {
	// The ID of the Stripe account this account belongs to. Should only be present if `account_holder.type` is `account`.
	Account *Account `json:"account"`
	// ID of the Stripe customer this account belongs to. Present if and only if `account_holder.type` is `customer`.
	Customer *Customer `json:"customer"`
	// Type of account holder that this account belongs to.
	Type FinancialConnectionsSessionAccountHolderType `json:"type"`
}
type FinancialConnectionsSessionFilters struct {
	// List of countries from which to filter accounts.
	Countries []string `json:"countries"`
}

// A Financial Connections Session is the secure way to programmatically launch the client-side Stripe.js modal that lets your users link their accounts.
type FinancialConnectionsSession struct {
	APIResource
	// The account holder for whom accounts are collected in this session.
	AccountHolder *FinancialConnectionsSessionAccountHolder `json:"account_holder"`
	// The accounts that were collected as part of this Session.
	Accounts *FinancialConnectionsAccountList `json:"accounts"`
	// A value that will be passed to the client to launch the authentication flow.
	ClientSecret string                              `json:"client_secret"`
	Filters      *FinancialConnectionsSessionFilters `json:"filters"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Permissions requested for accounts collected during this session.
	Permissions []FinancialConnectionsSessionPermission `json:"permissions"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url"`
}
