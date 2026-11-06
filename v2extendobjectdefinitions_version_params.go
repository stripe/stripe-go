//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Lists a page of object definition versions in descending semantic version order.
type V2ExtendObjectDefinitionsVersionListParams struct {
	Params `form:"*"`
	// The ID of the object definition whose versions to list.
	ObjectDefinitionID *string `form:"-" json:"-"` // Included in URL
	// Maximum page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves an object definition version.
type V2ExtendObjectDefinitionsVersionParams struct {
	Params `form:"*"`
	// The ID of the parent object definition.
	ObjectDefinitionID *string `form:"-" json:"-"` // Included in URL
}

// Retrieves an object definition version.
type V2ExtendObjectDefinitionsVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the parent object definition.
	ObjectDefinitionID *string `form:"-" json:"-"` // Included in URL
}
