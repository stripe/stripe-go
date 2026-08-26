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

// Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.
type FinancialConnectionsSessionFiltersAccountSubcategory string

// List of values that FinancialConnectionsSessionFiltersAccountSubcategory can take
const (
	FinancialConnectionsSessionFiltersAccountSubcategoryChecking     FinancialConnectionsSessionFiltersAccountSubcategory = "checking"
	FinancialConnectionsSessionFiltersAccountSubcategoryCreditCard   FinancialConnectionsSessionFiltersAccountSubcategory = "credit_card"
	FinancialConnectionsSessionFiltersAccountSubcategoryLineOfCredit FinancialConnectionsSessionFiltersAccountSubcategory = "line_of_credit"
	FinancialConnectionsSessionFiltersAccountSubcategoryMortgage     FinancialConnectionsSessionFiltersAccountSubcategory = "mortgage"
	FinancialConnectionsSessionFiltersAccountSubcategorySavings      FinancialConnectionsSessionFiltersAccountSubcategory = "savings"
)

// Whether the Session should require that linked accounts support payments and retrieve account numbers before completion.
type FinancialConnectionsSessionFiltersRequirePaymentMethodSupport string

// List of values that FinancialConnectionsSessionFiltersRequirePaymentMethodSupport can take
const (
	FinancialConnectionsSessionFiltersRequirePaymentMethodSupportAll        FinancialConnectionsSessionFiltersRequirePaymentMethodSupport = "all"
	FinancialConnectionsSessionFiltersRequirePaymentMethodSupportAtLeastOne FinancialConnectionsSessionFiltersRequirePaymentMethodSupport = "at_least_one"
	FinancialConnectionsSessionFiltersRequirePaymentMethodSupportNone       FinancialConnectionsSessionFiltersRequirePaymentMethodSupport = "none"
)

// Controls how manual entry of bank account details is presented to the user.
type FinancialConnectionsSessionManualEntryMode string

// List of values that FinancialConnectionsSessionManualEntryMode can take
const (
	FinancialConnectionsSessionManualEntryModeAutomatic FinancialConnectionsSessionManualEntryMode = "automatic"
	FinancialConnectionsSessionManualEntryModeDisabled  FinancialConnectionsSessionManualEntryMode = "disabled"
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

// Data features requested to be retrieved upon account creation.
type FinancialConnectionsSessionPrefetch string

// List of values that FinancialConnectionsSessionPrefetch can take
const (
	FinancialConnectionsSessionPrefetchBalances     FinancialConnectionsSessionPrefetch = "balances"
	FinancialConnectionsSessionPrefetchOwnership    FinancialConnectionsSessionPrefetch = "ownership"
	FinancialConnectionsSessionPrefetchTransactions FinancialConnectionsSessionPrefetch = "transactions"
)

// Retrieves the details of a Financial Connections Session
type FinancialConnectionsSessionParams struct {
	Params `form:"*"`
	// The account holder to link accounts for.
	AccountHolder *FinancialConnectionsSessionAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filters to restrict the kinds of accounts to collect.
	Filters *FinancialConnectionsSessionFiltersParams `form:"filters" json:"filters,omitempty"`
	// Settings for configuring Session-specific limits.
	Limits *FinancialConnectionsSessionLimitsParams `form:"limits" json:"limits,omitempty"`
	// Customize manual entry behavior
	ManualEntry *FinancialConnectionsSessionManualEntryParams `form:"manual_entry" json:"manual_entry,omitempty"`
	// List of data features that you would like to request access to.
	//
	// Possible values are `balances`, `transactions`, `ownership`, and `payment_method`.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder to link accounts for.
type FinancialConnectionsSessionAccountHolderParams struct {
	// The ID of the Stripe account whose accounts you will retrieve. Only available when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// The ID of the Stripe customer whose accounts you will retrieve. Only available when `type` is `customer`.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of Account representing a customer whose accounts you will retrieve. Only available when `type` is `customer`.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Type of account holder to collect accounts for.
	Type *string `form:"type" json:"type"`
}

// Filters to restrict the kinds of accounts to collect.
type FinancialConnectionsSessionFiltersParams struct {
	// Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
	// List of countries from which to collect accounts.
	Countries []*string `form:"countries" json:"countries,omitempty"`
	// Whether the session should require payment method support and successful account number retrieval before completion.
	RequirePaymentMethodSupport *string `form:"require_payment_method_support" json:"require_payment_method_support,omitempty"`
}

// Settings for configuring Session-specific limits.
type FinancialConnectionsSessionLimitsParams struct {
	// The number of accounts that can be linked in this Session. Pass an empty value to allow any number of accounts.
	Accounts    *int64                                              `form:"accounts" json:"accounts"`
	UnsetFields []FinancialConnectionsSessionLimitsParamsUnsetField `form:"-" json:"-"`
}

// FinancialConnectionsSessionLimitsParamsUnsetField is the list of fields that can be cleared/unset on FinancialConnectionsSessionLimitsParams.
type FinancialConnectionsSessionLimitsParamsUnsetField string

const (
	FinancialConnectionsSessionLimitsParamsUnsetFieldAccounts FinancialConnectionsSessionLimitsParamsUnsetField = "accounts"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *FinancialConnectionsSessionLimitsParams) AddUnsetField(field FinancialConnectionsSessionLimitsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Customize manual entry behavior
type FinancialConnectionsSessionManualEntryParams struct {
	// How manual entry should be handled.
	Mode *string `form:"mode" json:"mode,omitempty"`
}

// Retrieves the details of a Financial Connections Session
type FinancialConnectionsSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder to link accounts for.
type FinancialConnectionsSessionCreateAccountHolderParams struct {
	// The ID of the Stripe account whose accounts you will retrieve. Only available when `type` is `account`.
	Account *string `form:"account" json:"account,omitempty"`
	// The ID of the Stripe customer whose accounts you will retrieve. Only available when `type` is `customer`.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of Account representing a customer whose accounts you will retrieve. Only available when `type` is `customer`.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Type of account holder to collect accounts for.
	Type *string `form:"type" json:"type"`
}

// Filters to restrict the kinds of accounts to collect.
type FinancialConnectionsSessionCreateFiltersParams struct {
	// Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
	// List of countries from which to collect accounts.
	Countries []*string `form:"countries" json:"countries,omitempty"`
	// Whether the session should require payment method support and successful account number retrieval before completion.
	RequirePaymentMethodSupport *string `form:"require_payment_method_support" json:"require_payment_method_support,omitempty"`
}

// Settings for configuring Session-specific limits.
type FinancialConnectionsSessionCreateLimitsParams struct {
	// The number of accounts that can be linked in this Session. Pass an empty value to allow any number of accounts.
	Accounts    *int64                                                    `form:"accounts" json:"accounts"`
	UnsetFields []FinancialConnectionsSessionCreateLimitsParamsUnsetField `form:"-" json:"-"`
}

// FinancialConnectionsSessionCreateLimitsParamsUnsetField is the list of fields that can be cleared/unset on FinancialConnectionsSessionCreateLimitsParams.
type FinancialConnectionsSessionCreateLimitsParamsUnsetField string

const (
	FinancialConnectionsSessionCreateLimitsParamsUnsetFieldAccounts FinancialConnectionsSessionCreateLimitsParamsUnsetField = "accounts"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *FinancialConnectionsSessionCreateLimitsParams) AddUnsetField(field FinancialConnectionsSessionCreateLimitsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Customize manual entry behavior
type FinancialConnectionsSessionCreateManualEntryParams struct {
	// How manual entry should be handled.
	Mode *string `form:"mode" json:"mode,omitempty"`
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
type FinancialConnectionsSessionCreateParams struct {
	Params `form:"*"`
	// The account holder to link accounts for.
	AccountHolder *FinancialConnectionsSessionCreateAccountHolderParams `form:"account_holder" json:"account_holder"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filters to restrict the kinds of accounts to collect.
	Filters *FinancialConnectionsSessionCreateFiltersParams `form:"filters" json:"filters,omitempty"`
	// Settings for configuring Session-specific limits.
	Limits *FinancialConnectionsSessionCreateLimitsParams `form:"limits" json:"limits,omitempty"`
	// Customize manual entry behavior
	ManualEntry *FinancialConnectionsSessionCreateManualEntryParams `form:"manual_entry" json:"manual_entry,omitempty"`
	// List of data features that you would like to request access to.
	//
	// Possible values are `balances`, `transactions`, `ownership`, and `payment_method`.
	Permissions []*string `form:"permissions" json:"permissions"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder for whom accounts are collected in this session.
type FinancialConnectionsSessionAccountHolder struct {
	// The ID of the Stripe account that this account belongs to. Only available when `account_holder.type` is `account`.
	Account *Account `json:"account,omitempty"`
	// The ID for an Account representing a customer that this account belongs to. Only available when `account_holder.type` is `customer`.
	Customer        *Customer `json:"customer,omitempty"`
	CustomerAccount string    `json:"customer_account,omitempty"`
	// Type of account holder that this account belongs to.
	Type FinancialConnectionsSessionAccountHolderType `json:"type"`
}
type FinancialConnectionsSessionFilters struct {
	// Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.
	AccountSubcategories []FinancialConnectionsSessionFiltersAccountSubcategory `json:"account_subcategories"`
	// List of countries from which to filter accounts.
	Countries []string `json:"countries"`
	// Country from which to filter accounts.
	Country string `json:"country"`
	// Whether the Session should require that linked accounts support payments and retrieve account numbers before completion.
	RequirePaymentMethodSupport FinancialConnectionsSessionFiltersRequirePaymentMethodSupport `json:"require_payment_method_support,omitempty"`
}
type FinancialConnectionsSessionLimits struct {
	// The number of accounts that can be linked in this Session.
	Accounts int64 `json:"accounts"`
}
type FinancialConnectionsSessionManualEntry struct {
	// Controls how manual entry of bank account details is presented to the user.
	Mode FinancialConnectionsSessionManualEntryMode `json:"mode,omitempty"`
}

// A Financial Connections Session is the secure way to programmatically launch the client-side Stripe.js modal that lets your users link their accounts.
type FinancialConnectionsSession struct {
	APIResource
	// The account holder for whom accounts are collected in this session.
	AccountHolder *FinancialConnectionsSessionAccountHolder `json:"account_holder"`
	// The accounts that were collected as part of this Session.
	Accounts *FinancialConnectionsAccountList `json:"accounts"`
	// Tokenization is the process Stripe uses to collect sensitive card or bank
	// account details, or personally identifiable information (PII), directly from
	// your customers in a secure manner. A token representing this information is
	// returned to your server to use. Use our
	// [recommended payments integrations](https://docs.stripe.com/payments) to perform this process
	// on the client-side. This guarantees that no sensitive card data touches your server,
	// and allows your integration to operate in a PCI-compliant way.
	//
	// If you can't use client-side tokenization, you can also create tokens using
	// the API with either your publishable or secret API key. If
	// your integration uses this method, you're responsible for any PCI compliance
	// that it might require, and you must keep your secret API key safe. Unlike with
	// client-side tokenization, your customer's information isn't sent directly to
	// Stripe, so we can't determine how it's handled or stored.
	//
	// You can't store or use tokens more than once. To store card or bank account
	// information for later use, create [Customer](https://docs.stripe.com/api#customers)
	// objects or [External accounts](https://docs.stripe.com/api#external_accounts).
	// [Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,
	// performs best with integrations that use client-side tokenization.
	BankAccountToken *Token `json:"bank_account_token,omitempty"`
	// A value that will be passed to the client to launch the authentication flow.
	ClientSecret string                              `json:"client_secret"`
	Filters      *FinancialConnectionsSessionFilters `json:"filters,omitempty"`
	// Unique identifier for the object.
	ID     string                             `json:"id"`
	Limits *FinancialConnectionsSessionLimits `json:"limits,omitempty"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode    bool                                    `json:"livemode"`
	ManualEntry *FinancialConnectionsSessionManualEntry `json:"manual_entry,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Permissions requested for accounts collected during this session.
	Permissions []FinancialConnectionsSessionPermission `json:"permissions"`
	// Data features requested to be retrieved upon account creation.
	Prefetch []FinancialConnectionsSessionPrefetch `json:"prefetch"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url,omitempty"`
}
