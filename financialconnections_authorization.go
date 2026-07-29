//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the connection to the Authorization.
type FinancialConnectionsAuthorizationStatus string

// List of values that FinancialConnectionsAuthorizationStatus can take
const (
	FinancialConnectionsAuthorizationStatusActive   FinancialConnectionsAuthorizationStatus = "active"
	FinancialConnectionsAuthorizationStatusInactive FinancialConnectionsAuthorizationStatus = "inactive"
)

// The action (if any) to proactively relink the Authorization.
type FinancialConnectionsAuthorizationStatusDetailsActiveAction string

// List of values that FinancialConnectionsAuthorizationStatusDetailsActiveAction can take
const (
	FinancialConnectionsAuthorizationStatusDetailsActiveActionNone           FinancialConnectionsAuthorizationStatusDetailsActiveAction = "none"
	FinancialConnectionsAuthorizationStatusDetailsActiveActionRelinkRequired FinancialConnectionsAuthorizationStatusDetailsActiveAction = "relink_required"
)

// The action (if any) to relink the inactive Authorization.
type FinancialConnectionsAuthorizationStatusDetailsInactiveAction string

// List of values that FinancialConnectionsAuthorizationStatusDetailsInactiveAction can take
const (
	FinancialConnectionsAuthorizationStatusDetailsInactiveActionNone           FinancialConnectionsAuthorizationStatusDetailsInactiveAction = "none"
	FinancialConnectionsAuthorizationStatusDetailsInactiveActionRelinkRequired FinancialConnectionsAuthorizationStatusDetailsInactiveAction = "relink_required"
)

type FinancialConnectionsAuthorizationStatusDetailsActive struct {
	// The action (if any) to proactively relink the Authorization.
	Action FinancialConnectionsAuthorizationStatusDetailsActiveAction `json:"action"`
	// When the Authorization is expected to become inactive, if applicable.
	ExpectedDeactivationDate int64 `json:"expected_deactivation_date"`
}
type FinancialConnectionsAuthorizationStatusDetailsInactive struct {
	// The action (if any) to relink the inactive Authorization.
	Action FinancialConnectionsAuthorizationStatusDetailsInactiveAction `json:"action"`
}
type FinancialConnectionsAuthorizationStatusDetails struct {
	Active   *FinancialConnectionsAuthorizationStatusDetailsActive   `json:"active,omitempty"`
	Inactive *FinancialConnectionsAuthorizationStatusDetailsInactive `json:"inactive,omitempty"`
}

// An Authorization represents the set of credentials used to connect a group of Financial Connections Accounts.
type FinancialConnectionsAuthorization struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// The name of the institution that this authorization belongs to.
	InstitutionName string `json:"institution_name"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the connection to the Authorization.
	Status        FinancialConnectionsAuthorizationStatus         `json:"status"`
	StatusDetails *FinancialConnectionsAuthorizationStatusDetails `json:"status_details"`
}
