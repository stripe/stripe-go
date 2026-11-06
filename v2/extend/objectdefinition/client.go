//
//
// File generated from our OpenAPI spec
//
//

// Package objectdefinition provides the objectdefinition related APIs
package objectdefinition

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke objectdefinition related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an object definition.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2ExtendObjectDefinitionParams) (*stripe.V2ExtendObjectDefinition, error) {
	path := stripe.FormatURLPath("/v2/extend/object_definitions/%s", id)
	objectdefinition := &stripe.V2ExtendObjectDefinition{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, objectdefinition)
	return objectdefinition, err
}

// Lists a page of object definitions.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2ExtendObjectDefinitionListParams) stripe.Seq2[*stripe.V2ExtendObjectDefinition, error] {
	if listParams == nil {
		listParams = &stripe.V2ExtendObjectDefinitionListParams{}
	}
	return stripe.NewV2List("/v2/extend/object_definitions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2ExtendObjectDefinition], error) {
		page := &stripe.V2Page[*stripe.V2ExtendObjectDefinition]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
