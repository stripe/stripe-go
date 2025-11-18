//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List events, going back up to 30 days.
type V2CoreEventListParams struct {
	Params `form:"*"`
	// Set of filters to query events within a range of `created` timestamps.
	Created *RangeQueryParams `form:"created" json:"created,omitempty"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Primary object ID used to retrieve related events.
	ObjectID *string `form:"object_id" json:"object_id,omitempty"`
	// An array of up to 20 strings containing specific event names.
	Types []*string `form:"types" json:"types,omitempty"`
}

// Retrieves the details of an event.
type V2CoreEventParams struct {
	Params `form:"*"`
}

// Retrieves the details of an event.
type V2CoreEventRetrieveParams struct {
	Params `form:"*"`
}
