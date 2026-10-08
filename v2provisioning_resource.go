//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Catalog partition containing the resource's provider service.
type V2ProvisioningResourceCatalog string

// List of values that V2ProvisioningResourceCatalog can take
const (
	V2ProvisioningResourceCatalogDev     V2ProvisioningResourceCatalog = "dev"
	V2ProvisioningResourceCatalogProd    V2ProvisioningResourceCatalog = "prod"
	V2ProvisioningResourceCatalogTesting V2ProvisioningResourceCatalog = "testing"
)

// Provider environment in which the resource runs.
type V2ProvisioningResourceEnvironment string

// List of values that V2ProvisioningResourceEnvironment can take
const (
	V2ProvisioningResourceEnvironmentDev  V2ProvisioningResourceEnvironment = "dev"
	V2ProvisioningResourceEnvironmentProd V2ProvisioningResourceEnvironment = "prod"
)

// Current provisioning status of the resource.
type V2ProvisioningResourceStatus string

// List of values that V2ProvisioningResourceStatus can take
const (
	V2ProvisioningResourceStatusComplete         V2ProvisioningResourceStatus = "complete"
	V2ProvisioningResourceStatusErrored          V2ProvisioningResourceStatus = "errored"
	V2ProvisioningResourceStatusNeedsInformation V2ProvisioningResourceStatus = "needs_information"
	V2ProvisioningResourceStatusPending          V2ProvisioningResourceStatus = "pending"
	V2ProvisioningResourceStatusRemoved          V2ProvisioningResourceStatus = "removed"
)

// Message supplied by the provider when the resource becomes ready.
type V2ProvisioningResourceUserMessage struct {
	// Message from the provider to display to the user.
	Message string `json:"message"`
	// Time at which Stripe received the message from the provider.
	ReceivedAt time.Time `json:"received_at"`
}

// The `Resource` resource represents a provider-managed resource provisioned on behalf of
// a `Project`.
type V2ProvisioningResource struct {
	APIResource
	// Catalog partition containing the resource's provider service.
	Catalog V2ProvisioningResourceCatalog `json:"catalog,omitempty"`
	// Time at which the resource was created.
	Created time.Time `json:"created"`
	// Provider environment in which the resource runs.
	Environment V2ProvisioningResourceEnvironment `json:"environment"`
	// Error reported when provisioning the resource fails.
	ErrorMessage string `json:"error_message,omitempty"`
	// Unique identifier for the resource.
	ID string `json:"id"`
	// Whether this resource uses Stripe live-mode objects. This is independent of the provider
	// catalog and is immutable for the lifetime of the resource.
	Livemode bool `json:"livemode"`
	// Human-readable name of the resource.
	Name string `json:"name,omitempty"`
	// Schema describing additional information the provider requires to finish provisioning.
	NeedsInformationSchema map[string]any `json:"needs_information_schema,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Identifier of the provider that manages the resource.
	Provider string `json:"provider"`
	// Identifier of the provider service used to provision the resource.
	ServiceRef string `json:"service_ref"`
	// Current provisioning status of the resource.
	Status V2ProvisioningResourceStatus `json:"status"`
	// Message supplied by the provider when the resource becomes ready.
	UserMessage *V2ProvisioningResourceUserMessage `json:"user_message,omitempty"`
}
