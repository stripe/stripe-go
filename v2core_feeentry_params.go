//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List FeeEntries optionally filtered by incurred_by, fee_batch, or collection_record (at most one filter at a time).
type V2CoreFeeEntryListParams struct {
	Params `form:"*"`
	// Filter by money movement id (e.g. txn_xxx, bt_xxx).
	CollectionRecord *string `form:"collection_record" json:"collection_record,omitempty"`
	// Filter for FeeEntries created at the exact specified timestamp.
	Created *time.Time `form:"created" json:"created,omitempty"`
	// Filter for FeeEntries created after the specified timestamp (exclusive).
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for FeeEntries created at or after the specified timestamp (inclusive).
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for FeeEntries created before the specified timestamp (exclusive).
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for FeeEntries created at or before the specified timestamp (inclusive).
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// Filter by fee batch id (fb_xxx).
	FeeBatch *string `form:"fee_batch" json:"fee_batch,omitempty"`
	// Filter by usage object id (e.g. ch_xxx, py_xxx).
	IncurredBy *string `form:"incurred_by" json:"incurred_by,omitempty"`
	// Maximum number of results to return per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a FeeEntry by id.
type V2CoreFeeEntryParams struct {
	Params `form:"*"`
}

// Retrieve a FeeEntry by id.
type V2CoreFeeEntryRetrieveParams struct {
	Params `form:"*"`
}
