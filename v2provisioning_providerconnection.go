//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Current status of the service.
type V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus string

// List of values that V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus can take
const (
	V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatusActive   V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus = "active"
	V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatusInactive V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus = "inactive"
	V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatusPending  V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus = "pending"
)

// Action taken when the account was linked.
type V2ProvisioningProviderConnectionProviderAccountDetailsLinkAction string

// List of values that V2ProvisioningProviderConnectionProviderAccountDetailsLinkAction can take
const (
	V2ProvisioningProviderConnectionProviderAccountDetailsLinkActionCreated        V2ProvisioningProviderConnectionProviderAccountDetailsLinkAction = "created"
	V2ProvisioningProviderConnectionProviderAccountDetailsLinkActionLinkedExisting V2ProvisioningProviderConnectionProviderAccountDetailsLinkAction = "linked_existing"
)

// Current status of the provider connection.
type V2ProvisioningProviderConnectionStatus string

// List of values that V2ProvisioningProviderConnectionStatus can take
const (
	V2ProvisioningProviderConnectionStatusActive  V2ProvisioningProviderConnectionStatus = "active"
	V2ProvisioningProviderConnectionStatusExpired V2ProvisioningProviderConnectionStatus = "expired"
	V2ProvisioningProviderConnectionStatusUnknown V2ProvisioningProviderConnectionStatus = "unknown"
)

// Services active for the connected account.
type V2ProvisioningProviderConnectionProviderAccountDetailsActiveService struct {
	// Display name of the service.
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the resource at the provider that backs this service, if any.
	ProviderResourceID string `json:"provider_resource_id,omitempty"`
	// Identifier of the service at the provider.
	ServiceID string `json:"service_id"`
	// Current status of the service.
	Status V2ProvisioningProviderConnectionProviderAccountDetailsActiveServiceStatus `json:"status"`
}

// Details about the connected provider account.
type V2ProvisioningProviderConnectionProviderAccountDetails struct {
	// Services active for the connected account.
	ActiveServices []*V2ProvisioningProviderConnectionProviderAccountDetailsActiveService `json:"active_services"`
	// True when the provider explicitly supplied active_services, including an empty array.
	ActiveServicesProvided bool `json:"active_services_provided"`
	// Display name of the connected account.
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the connected account at the provider.
	ID string `json:"id"`
	// Action taken when the account was linked.
	LinkAction V2ProvisioningProviderConnectionProviderAccountDetailsLinkAction `json:"link_action,omitempty"`
	// Primary email address of the connected account.
	PrimaryEmail string `json:"primary_email,omitempty"`
}

// A ProviderConnection represents a link between a project and a provider account that
// resources can be created against; unlinking it prevents further resource creation.
type V2ProvisioningProviderConnection struct {
	APIResource
	// Time at which the provider connection was created.
	Created time.Time `json:"created,omitempty"`
	// Unique identifier for the provider connection.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Identifier of the provider this connection is linked to.
	Provider string `json:"provider"`
	// Identifier of the connected account at the provider, if one has been established.
	ProviderAccount string `json:"provider_account,omitempty"`
	// Details about the connected provider account.
	ProviderAccountDetails *V2ProvisioningProviderConnectionProviderAccountDetails `json:"provider_account_details,omitempty"`
	// Current status of the provider connection.
	Status V2ProvisioningProviderConnectionStatus `json:"status"`
}
