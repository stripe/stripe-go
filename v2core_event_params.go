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
	CreatedGt *time.Time `form:"created_gt" json:"created_gt,omitempty"`
	// Filter for events created at or after the specified timestamp.
	CreatedGTE *time.Time `form:"created_gte" json:"created_gte,omitempty"`
	// Filter for events created before the specified timestamp.
	CreatedLT *time.Time `form:"created_lt" json:"created_lt,omitempty"`
	// Filter for events created at or before the specified timestamp.
	CreatedLte *time.Time `form:"created_lte" json:"created_lte,omitempty"`
	// Filter events based on whether they were successfully delivered to all subscribed event destinations. If false, events which are still pending or have failed all delivery attempts to a event destination will be returned.
	DeliverySuccess *bool `form:"delivery_success" json:"delivery_success,omitempty"`
	// The page size.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Primary object ID used to retrieve related events.
	ObjectID *string `form:"object_id" json:"object_id,omitempty"`
}

// Retrieves the details of an event.
type V2CoreEventParams struct {
	Params `form:"*"`
}

// Retrieves the details of an event.
type V2CoreEventRetrieveParams struct {
	Params `form:"*"`
}
