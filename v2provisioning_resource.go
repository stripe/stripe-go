//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

type V2ProvisioningResourceCatalog string

// List of values that V2ProvisioningResourceCatalog can take
const (
	V2ProvisioningResourceCatalogDev     V2ProvisioningResourceCatalog = "dev"
	V2ProvisioningResourceCatalogProd    V2ProvisioningResourceCatalog = "prod"
	V2ProvisioningResourceCatalogTesting V2ProvisioningResourceCatalog = "testing"
)

type V2ProvisioningResourceEnvironment string

// List of values that V2ProvisioningResourceEnvironment can take
const (
	V2ProvisioningResourceEnvironmentDev  V2ProvisioningResourceEnvironment = "dev"
	V2ProvisioningResourceEnvironmentProd V2ProvisioningResourceEnvironment = "prod"
)

type V2ProvisioningResourceStatus string

// List of values that V2ProvisioningResourceStatus can take
const (
	V2ProvisioningResourceStatusComplete         V2ProvisioningResourceStatus = "complete"
	V2ProvisioningResourceStatusErrored          V2ProvisioningResourceStatus = "errored"
	V2ProvisioningResourceStatusNeedsInformation V2ProvisioningResourceStatus = "needs_information"
	V2ProvisioningResourceStatusPending          V2ProvisioningResourceStatus = "pending"
	V2ProvisioningResourceStatusRemoved          V2ProvisioningResourceStatus = "removed"
)

type V2ProvisioningResourceUserMessage struct {
	Message    string    `json:"message"`
	ReceivedAt time.Time `json:"received_at"`
}

// The `Resource` resource represents a provider-managed resource provisioned on behalf of
// a `Project`.
type V2ProvisioningResource struct {
	APIResource
	Catalog      V2ProvisioningResourceCatalog     `json:"catalog,omitempty"`
	Created      time.Time                         `json:"created"`
	Environment  V2ProvisioningResourceEnvironment `json:"environment"`
	ErrorMessage string                            `json:"error_message,omitempty"`
	ID           string                            `json:"id"`
	// Whether this resource uses Stripe live-mode objects. This is independent of the provider
	// catalog and is immutable for the lifetime of the resource.
	Livemode               bool           `json:"livemode"`
	Name                   string         `json:"name,omitempty"`
	NeedsInformationSchema map[string]any `json:"needs_information_schema,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object      string                             `json:"object"`
	Provider    string                             `json:"provider"`
	ServiceRef  string                             `json:"service_ref"`
	Status      V2ProvisioningResourceStatus       `json:"status"`
	UserMessage *V2ProvisioningResourceUserMessage `json:"user_message,omitempty"`
}
