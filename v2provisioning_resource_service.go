//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2ProvisioningResourceService is used to invoke resource related APIs.
type v2ProvisioningResourceService struct {
	B   Backend
	Key string
}

// Creates a new provider resource.
func (c v2ProvisioningResourceService) Create(ctx context.Context, params *V2ProvisioningResourceCreateParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceCreateParams{}
	}
	params.Context = ctx
	resource := &V2ProvisioningResource{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/resources", c.Key, params, resource)
	return resource, err
}

// Retrieves a provider resource.
func (c v2ProvisioningResourceService) Retrieve(ctx context.Context, id string, params *V2ProvisioningResourceRetrieveParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, resource)
	return resource, err
}

// Updates a resource's configuration or service.
func (c v2ProvisioningResourceService) Update(ctx context.Context, id string, params *V2ProvisioningResourceUpdateParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Links an existing provider resource to a project or account.
func (c v2ProvisioningResourceService) Link(ctx context.Context, params *V2ProvisioningResourceLinkParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceLinkParams{}
	}
	params.Context = ctx
	resource := &V2ProvisioningResource{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/resources/link", c.Key, params, resource)
	return resource, err
}

// Removes a resource.
func (c v2ProvisioningResourceService) Remove(ctx context.Context, id string, params *V2ProvisioningResourceRemoveParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceRemoveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s/remove", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Reveals the current provider-issued access configuration for a completed Resource.
// This is a read-only disclosure: it does not create, refresh, mint, or rotate credentials.
// Repeated calls are safe and do not require an idempotency key, but can return a newer
// configuration after a separate Rotate operation completes.
func (c v2ProvisioningResourceService) RevealAccessConfiguration(ctx context.Context, id string, params *V2ProvisioningResourceRevealAccessConfigurationParams) (*V2ProvisioningResourceAccessConfiguration, error) {
	if params == nil {
		params = &V2ProvisioningResourceRevealAccessConfigurationParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/provisioning/resources/%s/reveal_access_configuration", id)
	resourceaccessconfiguration := &V2ProvisioningResourceAccessConfiguration{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, resourceaccessconfiguration)
	return resourceaccessconfiguration, err
}

// Rotates a resource's credentials.
func (c v2ProvisioningResourceService) RotateCredentials(ctx context.Context, id string, params *V2ProvisioningResourceRotateCredentialsParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceRotateCredentialsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s/rotate_credentials", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Submits additional information requested by the provider for a resource.
func (c v2ProvisioningResourceService) SubmitInformation(ctx context.Context, id string, params *V2ProvisioningResourceSubmitInformationParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceSubmitInformationParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s/submit_information", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Unlinks a resource without removing it from the provider.
func (c v2ProvisioningResourceService) Unlink(ctx context.Context, id string, params *V2ProvisioningResourceUnlinkParams) (*V2ProvisioningResource, error) {
	if params == nil {
		params = &V2ProvisioningResourceUnlinkParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/resources/%s/unlink", id)
	resource := &V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}
