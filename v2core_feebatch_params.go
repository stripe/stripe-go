//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List FeeBatches optionally filtered by collection_record.
type V2CoreFeeBatchListParams struct {
	Params `form:"*"`
	// Filter to return only FeeBatches associated with this collection record ID.
	CollectionRecord *string `form:"collection_record" json:"collection_record,omitempty"`
	// Filter for FeeBatches created at the exact specified timestamp.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for FeeBatches created after the specified timestamp (exclusive).
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for FeeBatches created at or after the specified timestamp (inclusive).
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for FeeBatches created before the specified timestamp (exclusive).
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for FeeBatches created at or before the specified timestamp (inclusive).
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// Maximum number of results to return per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a FeeBatch by id.
type V2CoreFeeBatchParams struct {
	Params `form:"*"`
}

// Retrieve a FeeBatch by id.
type V2CoreFeeBatchRetrieveParams struct {
	Params `form:"*"`
}
