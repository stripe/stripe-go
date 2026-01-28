//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of account holder that this account belongs to.
type FinancialConnectionsAuthorizationAccountHolderType string

// List of values that FinancialConnectionsAuthorizationAccountHolderType can take
const (
	FinancialConnectionsAuthorizationAccountHolderTypeAccount  FinancialConnectionsAuthorizationAccountHolderType = "account"
	FinancialConnectionsAuthorizationAccountHolderTypeCustomer FinancialConnectionsAuthorizationAccountHolderType = "customer"
)

// The status of the connection to the Authorization.
type FinancialConnectionsAuthorizationStatus string

// List of values that FinancialConnectionsAuthorizationStatus can take
const (
	FinancialConnectionsAuthorizationStatusActive       FinancialConnectionsAuthorizationStatus = "active"
	FinancialConnectionsAuthorizationStatusDisconnected FinancialConnectionsAuthorizationStatus = "disconnected"
	FinancialConnectionsAuthorizationStatusInactive     FinancialConnectionsAuthorizationStatus = "inactive"
)

// The action (if any) to relink the inactive Authorization.
type FinancialConnectionsAuthorizationStatusDetailsInactiveAction string

// List of values that FinancialConnectionsAuthorizationStatusDetailsInactiveAction can take
const (
	FinancialConnectionsAuthorizationStatusDetailsInactiveActionNone           FinancialConnectionsAuthorizationStatusDetailsInactiveAction = "none"
	FinancialConnectionsAuthorizationStatusDetailsInactiveActionRelinkRequired FinancialConnectionsAuthorizationStatusDetailsInactiveAction = "relink_required"
)

// Retrieves the details of an Financial Connections Authorization.
type FinancialConnectionsAuthorizationParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsAuthorizationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of an Financial Connections Authorization.
type FinancialConnectionsAuthorizationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsAuthorizationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The account holder that this authorization belongs to.
type FinancialConnectionsAuthorizationAccountHolder struct {
	// The ID of the Stripe account that this account belongs to. Only available when `account_holder.type` is `account`.
	Account *Account `json:"account"`
	// The ID for an Account representing a customer that this account belongs to. Only available when `account_holder.type` is `customer`.
	Customer        *Customer `json:"customer"`
	CustomerAccount string    `json:"customer_account"`
	// Type of account holder that this account belongs to.
	Type FinancialConnectionsAuthorizationAccountHolderType `json:"type"`
}
type FinancialConnectionsAuthorizationStatusDetailsInactive struct {
	// The action (if any) to relink the inactive Authorization.
	Action FinancialConnectionsAuthorizationStatusDetailsInactiveAction `json:"action"`
}
type FinancialConnectionsAuthorizationStatusDetails struct {
	Inactive *FinancialConnectionsAuthorizationStatusDetailsInactive `json:"inactive"`
}

// An Authorization represents the set of credentials used to connect a group of Financial Connections Accounts.
type FinancialConnectionsAuthorization struct {
	APIResource
	// The account holder that this authorization belongs to.
	AccountHolder *FinancialConnectionsAuthorizationAccountHolder `json:"account_holder"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of the Financial Connections Institution this account belongs to. Note that this relationship may sometimes change in rare circumstances (e.g. institution mergers).
	Institution *FinancialConnectionsInstitution `json:"institution"`
	// The name of the institution that this authorization belongs to.
	InstitutionName string `json:"institution_name"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the connection to the Authorization.
	Status        FinancialConnectionsAuthorizationStatus         `json:"status"`
	StatusDetails *FinancialConnectionsAuthorizationStatusDetails `json:"status_details"`
}
