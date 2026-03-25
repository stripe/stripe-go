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

// How the user enters the hosted flow. You can only use the values `email` and `url` if you provide `relink_options`.
type FinancialConnectionsSessionHostedDeliveryMethod string

// List of values that FinancialConnectionsSessionHostedDeliveryMethod can take
const (
	FinancialConnectionsSessionHostedDeliveryMethodEmail FinancialConnectionsSessionHostedDeliveryMethod = "email"
	FinancialConnectionsSessionHostedDeliveryMethodURL   FinancialConnectionsSessionHostedDeliveryMethod = "url"
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
	FinancialConnectionsSessionPrefetchBalances         FinancialConnectionsSessionPrefetch = "balances"
	FinancialConnectionsSessionPrefetchInferredBalances FinancialConnectionsSessionPrefetch = "inferred_balances"
	FinancialConnectionsSessionPrefetchOwnership        FinancialConnectionsSessionPrefetch = "ownership"
	FinancialConnectionsSessionPrefetchTransactions     FinancialConnectionsSessionPrefetch = "transactions"
)

// Reason for why relink failed. One of `no_authorization`, `no_account`, or `other`.
type FinancialConnectionsSessionRelinkResultFailureReason string

// List of values that FinancialConnectionsSessionRelinkResultFailureReason can take
const (
	FinancialConnectionsSessionRelinkResultFailureReasonNoAccount       FinancialConnectionsSessionRelinkResultFailureReason = "no_account"
	FinancialConnectionsSessionRelinkResultFailureReasonNoAuthorization FinancialConnectionsSessionRelinkResultFailureReason = "no_authorization"
	FinancialConnectionsSessionRelinkResultFailureReasonOther           FinancialConnectionsSessionRelinkResultFailureReason = "other"
)

// The current state of the session.
type FinancialConnectionsSessionStatus string

// List of values that FinancialConnectionsSessionStatus can take
const (
	FinancialConnectionsSessionStatusCancelled FinancialConnectionsSessionStatus = "cancelled"
	FinancialConnectionsSessionStatusFailed    FinancialConnectionsSessionStatus = "failed"
	FinancialConnectionsSessionStatusPending   FinancialConnectionsSessionStatus = "pending"
	FinancialConnectionsSessionStatusSucceeded FinancialConnectionsSessionStatus = "succeeded"
)

// The reason for the Session being cancelled.
type FinancialConnectionsSessionStatusDetailsCancelledReason string

// List of values that FinancialConnectionsSessionStatusDetailsCancelledReason can take
const (
	FinancialConnectionsSessionStatusDetailsCancelledReasonCustomManualEntry FinancialConnectionsSessionStatusDetailsCancelledReason = "custom_manual_entry"
	FinancialConnectionsSessionStatusDetailsCancelledReasonOther             FinancialConnectionsSessionStatusDetailsCancelledReason = "other"
)

// The UI mode for this session.
type FinancialConnectionsSessionUIMode string

// List of values that FinancialConnectionsSessionUIMode can take
const (
	FinancialConnectionsSessionUIModeHosted FinancialConnectionsSessionUIMode = "hosted"
	FinancialConnectionsSessionUIModeModal  FinancialConnectionsSessionUIMode = "modal"
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
	// Settings for hosted Sessions. Required if `ui_mode` is `hosted`.
	Hosted *FinancialConnectionsSessionHostedParams `form:"hosted" json:"hosted,omitempty"`
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
	// Options for specifying a Session targeted to relinking an authorization.
	RelinkOptions *FinancialConnectionsSessionRelinkOptionsParams `form:"relink_options" json:"relink_options,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
	// The UI mode of the Session. Defaults to `modal`.
	UIMode *string `form:"ui_mode" json:"ui_mode,omitempty"`
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
	// Stripe ID of the institution with which the customer should be directed to log in.
	Institution *string `form:"institution" json:"institution,omitempty"`
}

// Settings for hosted Sessions. Required if `ui_mode` is `hosted`.
type FinancialConnectionsSessionHostedParams struct {
	// How the user should enter the hosted flow. The values `email` and `url` can only be used if `relink_options` is provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
}

// Settings for configuring Session-specific limits.
type FinancialConnectionsSessionLimitsParams struct {
	// The number of accounts that can be linked in this Session.
	Accounts *int64 `form:"accounts" json:"accounts"`
}

// Customize manual entry behavior
type FinancialConnectionsSessionManualEntryParams struct {
	// How manual entry should be handled.
	Mode *string `form:"mode" json:"mode,omitempty"`
}

// Options for specifying a Session targeted to relinking an authorization.
type FinancialConnectionsSessionRelinkOptionsParams struct {
	// The account to relink. Must belong to the authorization specified in `authorization`.
	Account *string `form:"account" json:"account,omitempty"`
	// The authorization to relink.
	Authorization *string `form:"authorization" json:"authorization"`
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
	// Stripe ID of the institution with which the customer should be directed to log in.
	Institution *string `form:"institution" json:"institution,omitempty"`
}

// Settings for hosted Sessions. Required if `ui_mode` is `hosted`.
type FinancialConnectionsSessionCreateHostedParams struct {
	// How the user should enter the hosted flow. The values `email` and `url` can only be used if `relink_options` is provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
}

// Settings for configuring Session-specific limits.
type FinancialConnectionsSessionCreateLimitsParams struct {
	// The number of accounts that can be linked in this Session.
	Accounts *int64 `form:"accounts" json:"accounts"`
}

// Customize manual entry behavior
type FinancialConnectionsSessionCreateManualEntryParams struct {
	// How manual entry should be handled.
	Mode *string `form:"mode" json:"mode,omitempty"`
}

// Options for specifying a Session targeted to relinking an authorization.
type FinancialConnectionsSessionCreateRelinkOptionsParams struct {
	// The account to relink. Must belong to the authorization specified in `authorization`.
	Account *string `form:"account" json:"account,omitempty"`
	// The authorization to relink.
	Authorization *string `form:"authorization" json:"authorization"`
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
type FinancialConnectionsSessionCreateParams struct {
	Params `form:"*"`
	// The account holder to link accounts for.
	AccountHolder *FinancialConnectionsSessionCreateAccountHolderParams `form:"account_holder" json:"account_holder,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filters to restrict the kinds of accounts to collect.
	Filters *FinancialConnectionsSessionCreateFiltersParams `form:"filters" json:"filters,omitempty"`
	// Settings for hosted Sessions. Required if `ui_mode` is `hosted`.
	Hosted *FinancialConnectionsSessionCreateHostedParams `form:"hosted" json:"hosted,omitempty"`
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
	// Options for specifying a Session targeted to relinking an authorization.
	RelinkOptions *FinancialConnectionsSessionCreateRelinkOptionsParams `form:"relink_options" json:"relink_options,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
	// The UI mode of the Session. Defaults to `modal`.
	UIMode *string `form:"ui_mode" json:"ui_mode,omitempty"`
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
	// Stripe ID of the institution with which the customer should be directed to log in.
	Institution string `json:"institution,omitempty"`
}

// Settings for the Hosted UI mode.
type FinancialConnectionsSessionHosted struct {
	// How the user enters the hosted flow. You can only use the values `email` and `url` if you provide `relink_options`.
	DeliveryMethod FinancialConnectionsSessionHostedDeliveryMethod `json:"delivery_method,omitempty"`
	// The URL to redirect your customer back to after they link their accounts or cancel this Session. This parameter is required if `ui_mode` is `hosted`.
	ReturnURL string `json:"return_url"`
}
type FinancialConnectionsSessionLimits struct {
	// The number of accounts that can be linked in this Session.
	Accounts int64 `json:"accounts"`
}
type FinancialConnectionsSessionManualEntry struct{}
type FinancialConnectionsSessionRelinkOptions struct {
	// Requires the end user to repair this specific account during the authentication flow instead of connecting a different one.
	Account string `json:"account,omitempty"`
	// The authorization to relink in the Session.
	Authorization string `json:"authorization"`
}
type FinancialConnectionsSessionRelinkResult struct {
	// The account relinked in the Session. Only present if `relink_options[account]` is set and relink is successful.
	Account string `json:"account"`
	// The authorization relinked in the Session. Only present if relink is successful.
	Authorization string `json:"authorization"`
	// Reason for why relink failed. One of `no_authorization`, `no_account`, or `other`.
	FailureReason FinancialConnectionsSessionRelinkResultFailureReason `json:"failure_reason"`
}
type FinancialConnectionsSessionStatusDetailsCancelled struct {
	// The reason for the Session being cancelled.
	Reason FinancialConnectionsSessionStatusDetailsCancelledReason `json:"reason"`
}
type FinancialConnectionsSessionStatusDetails struct {
	Cancelled *FinancialConnectionsSessionStatusDetailsCancelled `json:"cancelled,omitempty"`
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
	Filters      *FinancialConnectionsSessionFilters `json:"filters,omitempty"`
	// Settings for the Hosted UI mode.
	Hosted *FinancialConnectionsSessionHosted `json:"hosted,omitempty"`
	// Unique identifier for the object.
	ID     string                             `json:"id"`
	Limits *FinancialConnectionsSessionLimits `json:"limits,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode    bool                                    `json:"livemode"`
	ManualEntry *FinancialConnectionsSessionManualEntry `json:"manual_entry,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Permissions requested for accounts collected during this session.
	Permissions []FinancialConnectionsSessionPermission `json:"permissions"`
	// Data features requested to be retrieved upon account creation.
	Prefetch      []FinancialConnectionsSessionPrefetch     `json:"prefetch"`
	RelinkOptions *FinancialConnectionsSessionRelinkOptions `json:"relink_options,omitempty"`
	RelinkResult  *FinancialConnectionsSessionRelinkResult  `json:"relink_result,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnURL string `json:"return_url,omitempty"`
	// The current state of the session.
	Status        FinancialConnectionsSessionStatus         `json:"status,omitempty"`
	StatusDetails *FinancialConnectionsSessionStatusDetails `json:"status_details,omitempty"`
	// The UI mode for this session.
	UIMode FinancialConnectionsSessionUIMode `json:"ui_mode,omitempty"`
	// The hosted URL for this Session. Redirect customers to this URL to take them to the hosted authentication flow. This value is only present when the Session is active and the `ui_mode` is `hosted`.
	URL string `json:"url,omitempty"`
}
