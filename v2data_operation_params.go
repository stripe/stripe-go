//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Resolve a natural-language query into ranked public analytical resources
// (schemas, reports, and metrics) with execution context.
type V2DataOperationSearchParams struct {
	Params `form:"*"`
	// Restrict results to public resource catalogs. Defaults to all catalogs
	// available to the caller.
	Catalogs []*string `form:"catalogs" json:"catalogs,omitempty"`
	// Maximum results. Default: 5, Max: 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The page token for pagination.
	Page *string `form:"page" json:"page,omitempty"`
	// Natural-language question. Required.
	Query *string `form:"query" json:"query"`
}
