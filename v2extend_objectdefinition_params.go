//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists a page of object definitions.
type V2ExtendObjectDefinitionListParams struct {
	Params `form:"*"`
	// Maximum page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves an object definition.
type V2ExtendObjectDefinitionParams struct {
	Params `form:"*"`
}

// Retrieves an object definition.
type V2ExtendObjectDefinitionRetrieveParams struct {
	Params `form:"*"`
}
