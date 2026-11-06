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

// v2ExtendObjectDefinitionsVersionService is used to invoke version related APIs.
type v2ExtendObjectDefinitionsVersionService struct {
	B   Backend
	Key string
}

// Retrieves an object definition version.
func (c v2ExtendObjectDefinitionsVersionService) Retrieve(ctx context.Context, id string, params *V2ExtendObjectDefinitionsVersionRetrieveParams) (*V2ExtendObjectDefinitionVersion, error) {
	if params == nil {
		params = &V2ExtendObjectDefinitionsVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/extend/object_definitions/%s/versions/%s", StringValue(
			params.ObjectDefinitionID), id)
	objectdefinitionversion := &V2ExtendObjectDefinitionVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, objectdefinitionversion)
	return objectdefinitionversion, err
}

// Lists a page of object definition versions in descending semantic version order.
func (c v2ExtendObjectDefinitionsVersionService) List(ctx context.Context, listParams *V2ExtendObjectDefinitionsVersionListParams) *V2List[*V2ExtendObjectDefinitionVersion] {
	if listParams == nil {
		listParams = &V2ExtendObjectDefinitionsVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/extend/object_definitions/%s/versions", StringValue(
			listParams.ObjectDefinitionID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ExtendObjectDefinitionVersion], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ExtendObjectDefinitionVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
