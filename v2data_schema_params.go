//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of schemas describing the tables available to query.
type V2DataSchemaListParams struct {
	Params `form:"*"`
	// If supplied, only return schemas belonging to this dataset.
	Dataset *string `form:"dataset" json:"dataset,omitempty"`
	// Any optional includes (see https://docs.stripe.com/api-includable-response-values).
	Include []*string `form:"include" json:"include,omitempty"`
	// The maximum number of results per page. Defaults to 10. Maximum is 100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// If supplied, only return schemas with this name.
	Name *string `form:"name" json:"name,omitempty"`
}

// Retrieves the schema for a particular table.
type V2DataSchemaParams struct {
	Params `form:"*"`
}

// Retrieves the schema for a particular table.
type V2DataSchemaRetrieveParams struct {
	Params `form:"*"`
}
