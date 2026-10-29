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

// v2DataSchemaService is used to invoke schema related APIs.
type v2DataSchemaService struct {
	B   Backend
	Key string
}

// Retrieves the schema for a particular table.
func (c v2DataSchemaService) Retrieve(ctx context.Context, id string, params *V2DataSchemaRetrieveParams) (*V2DataSchema, error) {
	if params == nil {
		params = &V2DataSchemaRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/data/schemas/%s", id)
	schema := &V2DataSchema{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, schema)
	return schema, err
}

// Returns a list of schemas describing the tables available to query.
func (c v2DataSchemaService) List(ctx context.Context, listParams *V2DataSchemaListParams) *V2List[*V2DataSchema] {
	if listParams == nil {
		listParams = &V2DataSchemaListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/data/schemas", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2DataSchema], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2DataSchema]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
