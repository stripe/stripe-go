//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a new provider resource.
type V2ProvisioningResourceParams struct {
	Params `form:"*"`
	// Catalog partition to create the resource in.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// New provider-specific configuration payload for the resource.
	Configuration map[string]any `form:"configuration" json:"configuration,omitempty"`
	// Environment the resource should be created in.
	Environment *string `form:"environment" json:"environment,omitempty"`
	// Whether the resource should use Stripe live-mode objects. When omitted, this resolves to true.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Human-readable name for the resource.
	Name *string `form:"name" json:"name,omitempty"`
	// Identifier of the project to create the resource in.
	Project *string `form:"project" json:"project,omitempty"`
	// Identifier of the provider to create the resource with.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// Provider's service id to switch the resource to. If omitted, the resource's existing service
	// is retained and this is treated as a config-only update.
	ServiceRef *string `form:"service_ref" json:"service_ref,omitempty"`
}

// Links an existing provider resource to a project or account.
type V2ProvisioningResourceLinkParams struct {
	Params `form:"*"`
	// Catalog partition of the existing resource.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// Environment the existing resource runs in.
	Environment *string `form:"environment" json:"environment,omitempty"`
	// Whether the resource should use Stripe live-mode objects. When omitted, this resolves to true.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Identifier of the project to link the resource to.
	Project *string `form:"project" json:"project,omitempty"`
	// Identifier of the provider that hosts the existing resource.
	Provider *string `form:"provider" json:"provider"`
	// Identifier of the provider service the existing resource belongs to.
	ServiceRef *string `form:"service_ref" json:"service_ref"`
}

// Removes a resource.
type V2ProvisioningResourceRemoveParams struct {
	Params `form:"*"`
}

// Rotates a resource's credentials.
type V2ProvisioningResourceRotateCredentialsParams struct {
	Params `form:"*"`
}

// Submits additional information requested by the provider for a resource.
type V2ProvisioningResourceSubmitInformationParams struct {
	Params `form:"*"`
	// Additional information being submitted for the resource.
	SubmittedInformation map[string]any `form:"submitted_information" json:"submitted_information"`
}

// Unlinks a resource without removing it from the provider.
type V2ProvisioningResourceUnlinkParams struct {
	Params `form:"*"`
}

// Creates a new provider resource.
type V2ProvisioningResourceCreateParams struct {
	Params `form:"*"`
	// Catalog partition to create the resource in.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// Provider-specific configuration payload for the resource.
	Configuration map[string]any `form:"configuration" json:"configuration"`
	// Environment the resource should be created in.
	Environment *string `form:"environment" json:"environment,omitempty"`
	// Whether the resource should use Stripe live-mode objects. When omitted, this resolves to true.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Human-readable name for the resource.
	Name *string `form:"name" json:"name,omitempty"`
	// Identifier of the project to create the resource in.
	Project *string `form:"project" json:"project,omitempty"`
	// Identifier of the provider to create the resource with.
	Provider *string `form:"provider" json:"provider"`
	// Identifier of the provider service to create the resource from.
	ServiceRef *string `form:"service_ref" json:"service_ref"`
}

// Retrieves a provider resource.
type V2ProvisioningResourceRetrieveParams struct {
	Params `form:"*"`
}

// Updates a resource's configuration or service.
type V2ProvisioningResourceUpdateParams struct {
	Params `form:"*"`
	// Catalog partition of the resource.
	Catalog *string `form:"catalog" json:"catalog,omitempty"`
	// New provider-specific configuration payload for the resource.
	Configuration map[string]any `form:"configuration" json:"configuration,omitempty"`
	// Provider's service id to switch the resource to. If omitted, the resource's existing service
	// is retained and this is treated as a config-only update.
	ServiceRef *string `form:"service_ref" json:"service_ref,omitempty"`
}
