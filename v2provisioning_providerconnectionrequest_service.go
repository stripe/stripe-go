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

// v2ProvisioningProviderConnectionRequestService is used to invoke providerconnectionrequest related APIs.
type v2ProvisioningProviderConnectionRequestService struct {
	B   Backend
	Key string
}

// Creates a new provider connection.
func (c v2ProvisioningProviderConnectionRequestService) Create(ctx context.Context, params *V2ProvisioningProviderConnectionRequestCreateParams) (*V2ProvisioningProviderConnectionRequest, error) {
	if params == nil {
		params = &V2ProvisioningProviderConnectionRequestCreateParams{}
	}
	params.Context = ctx
	providerconnectionrequest := &V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/provider_connection_requests", c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}

// Retrieves a provider connection.
func (c v2ProvisioningProviderConnectionRequestService) Retrieve(ctx context.Context, id string, params *V2ProvisioningProviderConnectionRequestRetrieveParams) (*V2ProvisioningProviderConnectionRequest, error) {
	if params == nil {
		params = &V2ProvisioningProviderConnectionRequestRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/provider_connection_requests/%s", id)
	providerconnectionrequest := &V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}

// Submits additional information requested by the provider for a provider connection.
func (c v2ProvisioningProviderConnectionRequestService) SubmitInformation(ctx context.Context, id string, params *V2ProvisioningProviderConnectionRequestSubmitInformationParams) (*V2ProvisioningProviderConnectionRequest, error) {
	if params == nil {
		params = &V2ProvisioningProviderConnectionRequestSubmitInformationParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/provisioning/provider_connection_requests/%s/submit_information", id)
	providerconnectionrequest := &V2ProvisioningProviderConnectionRequest{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, providerconnectionrequest)
	return providerconnectionrequest, err
}
