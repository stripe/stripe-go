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

// v2ExtendObjectDefinitionService is used to invoke objectdefinition related APIs.
type v2ExtendObjectDefinitionService struct {
	B   Backend
	Key string
}

// Retrieves an object definition.
func (c v2ExtendObjectDefinitionService) Retrieve(ctx context.Context, id string, params *V2ExtendObjectDefinitionRetrieveParams) (*V2ExtendObjectDefinition, error) {
	if params == nil {
		params = &V2ExtendObjectDefinitionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/extend/object_definitions/%s", id)
	objectdefinition := &V2ExtendObjectDefinition{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, objectdefinition)
	return objectdefinition, err
}

// Lists a page of object definitions.
func (c v2ExtendObjectDefinitionService) List(ctx context.Context, listParams *V2ExtendObjectDefinitionListParams) *V2List[*V2ExtendObjectDefinition] {
	if listParams == nil {
		listParams = &V2ExtendObjectDefinitionListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/extend/object_definitions", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ExtendObjectDefinition], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ExtendObjectDefinition]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
