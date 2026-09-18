//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Catalog partition the project belongs to.
type V2ProvisioningProjectCatalog string

// List of values that V2ProvisioningProjectCatalog can take
const (
	V2ProvisioningProjectCatalogDev     V2ProvisioningProjectCatalog = "dev"
	V2ProvisioningProjectCatalogProd    V2ProvisioningProjectCatalog = "prod"
	V2ProvisioningProjectCatalogTesting V2ProvisioningProjectCatalog = "testing"
)

// Fields of the developer profile that have been verified.
type V2ProvisioningProjectProfileVerifiedField string

// List of values that V2ProvisioningProjectProfileVerifiedField can take
const (
	V2ProvisioningProjectProfileVerifiedFieldCountry V2ProvisioningProjectProfileVerifiedField = "country"
	V2ProvisioningProjectProfileVerifiedFieldEmail   V2ProvisioningProjectProfileVerifiedField = "email"
	V2ProvisioningProjectProfileVerifiedFieldName    V2ProvisioningProjectProfileVerifiedField = "name"
	V2ProvisioningProjectProfileVerifiedFieldPhone   V2ProvisioningProjectProfileVerifiedField = "phone"
)

// Use the /v2/provisioning/identity endpoint instead for IAM information.
type V2ProvisioningProjectProfile struct {
	// Email address associated with the developer profile.
	Email string `json:"email,omitempty"`
	// Fields of the developer profile that have been verified.
	VerifiedFields []V2ProvisioningProjectProfileVerifiedField `json:"verified_fields"`
}

// The `Project` resource represents a container for provisioned resources and their
// associated configuration.
type V2ProvisioningProject struct {
	APIResource
	// Catalog partition the project belongs to.
	Catalog V2ProvisioningProjectCatalog `json:"catalog"`
	// Time at which the project was created.
	Created time.Time `json:"created"`
	// Unique identifier for the project.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Human-readable name of the project.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Use the /v2/provisioning/identity endpoint instead for IAM information.
	Profile *V2ProvisioningProjectProfile `json:"profile,omitempty"`
	// Identifier of the developer profile associated with the project.
	ProjectProfile string `json:"project_profile,omitempty"`
}
