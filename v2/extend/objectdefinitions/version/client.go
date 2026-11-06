//
//
// File generated from our OpenAPI spec
//
//

// Package version provides the version related APIs
package version

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke version related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an object definition version.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ExtendObjectDefinitionsVersionParams) (*stripe.V2ExtendObjectDefinitionVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/extend/object_definitions/%s/versions/%s", stripe.StringValue(
			params.ObjectDefinitionID), id)
	objectdefinitionversion := &stripe.V2ExtendObjectDefinitionVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, objectdefinitionversion)
	return objectdefinitionversion, err
}

// Lists a page of object definition versions in descending semantic version order.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ExtendObjectDefinitionsVersionListParams) stripe.Seq2[*stripe.V2ExtendObjectDefinitionVersion, error] {
	if listParams == nil {
		listParams = &stripe.V2ExtendObjectDefinitionsVersionListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/extend/object_definitions/%s/versions", stripe.StringValue(
			listParams.ObjectDefinitionID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ExtendObjectDefinitionVersion], error) {
		page := &stripe.V2Page[*stripe.V2ExtendObjectDefinitionVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
