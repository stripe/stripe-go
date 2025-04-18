//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List events, going back up to 30 days.
type V2CoreEventListParams struct {
	Params `form:"*"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Primary object ID used to retrieve related events.
	ObjectID *string `form:"object_id" json:"object_id"`
}

// Retrieves the details of an event.
type V2CoreEventParams struct {
	Params `form:"*"`
}

// Retrieves the details of an event.
type V2CoreEventRetrieveParams struct {
	Params `form:"*"`
}
