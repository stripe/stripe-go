//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Status of the underlying account-linking workflow. Unset once the workflow completes; see
// provider_connection for the resulting connection's status.
type V2ProvisioningProviderConnectionRequestRequestStatus string

// List of values that V2ProvisioningProviderConnectionRequestRequestStatus can take
const (
	V2ProvisioningProviderConnectionRequestRequestStatusComplete         V2ProvisioningProviderConnectionRequestRequestStatus = "complete"
	V2ProvisioningProviderConnectionRequestRequestStatusError            V2ProvisioningProviderConnectionRequestRequestStatus = "error"
	V2ProvisioningProviderConnectionRequestRequestStatusNeedsInformation V2ProvisioningProviderConnectionRequestRequestStatus = "needs_information"
	V2ProvisioningProviderConnectionRequestRequestStatusPendingAuth      V2ProvisioningProviderConnectionRequestRequestStatus = "pending_auth"
	V2ProvisioningProviderConnectionRequestRequestStatusRequested        V2ProvisioningProviderConnectionRequestRequestStatus = "requested"
)

// Error from the account-linking workflow, set when request_status is ERROR.
type V2ProvisioningProviderConnectionRequestError struct {
	// Machine-readable error code.
	Code string `json:"code"`
	// Human-readable error message.
	Message string `json:"message"`
}

// A ProviderConnectionRequest represents an in-progress account-linking workflow. Once the
// workflow completes, `provider_connection` is populated with the resulting ProviderConnection.
type V2ProvisioningProviderConnectionRequest struct {
	APIResource
	// Time at which the provider connection request was created.
	Created time.Time `json:"created,omitempty"`
	// Error from the account-linking workflow, set when request_status is ERROR.
	Error *V2ProvisioningProviderConnectionRequestError `json:"error,omitempty"`
	// Unique identifier for the provider connection request.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Schema describing the information the provider still needs, set when request_status is
	// NEEDS_INFORMATION.
	NeedsInformationSchema map[string]any `json:"needs_information_schema,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Identifier of the provider this connection request is linked to.
	Provider string `json:"provider"`
	// A ProviderConnection represents a link between a project and a provider account that
	// resources can be created against; unlinking it prevents further resource creation.
	ProviderConnection *V2ProvisioningProviderConnection `json:"provider_connection,omitempty"`
	// URL the caller should redirect to in order to continue the account-linking workflow.
	RedirectURL string `json:"redirect_url,omitempty"`
	// Status of the underlying account-linking workflow. Unset once the workflow completes; see
	// provider_connection for the resulting connection's status.
	RequestStatus V2ProvisioningProviderConnectionRequestRequestStatus `json:"request_status"`
	// Scopes requested for the account-linking workflow.
	Scopes []string `json:"scopes"`
}
