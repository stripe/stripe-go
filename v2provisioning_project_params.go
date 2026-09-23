//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a new project.
type V2ProvisioningProjectParams struct {
	Params `form:"*"`
	// Catalog partition to create the project in.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// Human-readable name for the new project.
	Name *string `form:"name" json:"name"`
	// Identifier of the developer profile to associate with the new project.
	ProjectProfile *string `form:"project_profile" json:"project_profile,omitempty"`
}

// Creates a new project.
type V2ProvisioningProjectCreateParams struct {
	Params `form:"*"`
	// Catalog partition to create the project in.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// Human-readable name for the new project.
	Name *string `form:"name" json:"name"`
	// Identifier of the developer profile to associate with the new project.
	ProjectProfile *string `form:"project_profile" json:"project_profile,omitempty"`
}
