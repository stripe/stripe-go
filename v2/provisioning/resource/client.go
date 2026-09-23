//
//
// File generated from our OpenAPI spec
//
//

// Package resource provides the resource related APIs
package resource

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke resource related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new provider resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2ProvisioningResourceParams) (*stripe.V2ProvisioningResource, error) {
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/resources", c.Key, params, resource)
	return resource, err
}

// Retrieves a provider resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ProvisioningResourceParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath("/v2/provisioning/resources/%s", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, resource)
	return resource, err
}

// Updates a resource's configuration or service.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2ProvisioningResourceParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath("/v2/provisioning/resources/%s", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Links an existing provider resource to a project or account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Link(params *stripe.V2ProvisioningResourceLinkParams) (*stripe.V2ProvisioningResource, error) {
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/resources/link", c.Key, params, resource)
	return resource, err
}

// Removes a resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Remove(id string, params *stripe.V2ProvisioningResourceRemoveParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath("/v2/provisioning/resources/%s/remove", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Rotates a resource's credentials.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RotateCredentials(id string, params *stripe.V2ProvisioningResourceRotateCredentialsParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath(
		"/v2/provisioning/resources/%s/rotate_credentials", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Submits additional information requested by the provider for a resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SubmitInformation(id string, params *stripe.V2ProvisioningResourceSubmitInformationParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath(
		"/v2/provisioning/resources/%s/submit_information", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}

// Unlinks a resource without removing it from the provider.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Unlink(id string, params *stripe.V2ProvisioningResourceUnlinkParams) (*stripe.V2ProvisioningResource, error) {
	path := stripe.FormatURLPath("/v2/provisioning/resources/%s/unlink", id)
	resource := &stripe.V2ProvisioningResource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, resource)
	return resource, err
}
