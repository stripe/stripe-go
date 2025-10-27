//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List events, going back up to 30 days.
type V2CoreEventListParams struct {
	Params `form:"*"`
	// Filter for events created after the specified timestamp.
	Gt *time.Time `form:"gt" json:"gt,omitempty"`
	// Filter for events created at or after the specified timestamp.
	GTE *time.Time `form:"gte" json:"gte,omitempty"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter for events created before the specified timestamp.
	LT *time.Time `form:"lt" json:"lt,omitempty"`
	// Filter for events created at or before the specified timestamp.
	Lte *time.Time `form:"lte" json:"lte,omitempty"`
	// Primary object ID used to retrieve related events.
	ObjectID *string `form:"object_id" json:"object_id,omitempty"`
	// An array of up to 20 strings containing specific event names.
	Types []*string `form:"types,flat_array" json:"types,omitempty"`
}

// Retrieves the details of an event.
type V2CoreEventParams struct {
	Params `form:"*"`
}

// Retrieves the details of an event.
type V2CoreEventRetrieveParams struct {
	Params `form:"*"`
}
