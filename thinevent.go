package stripe

import "time"

// ThinEvent represents the json that's delivered from an Event Destination.
// Use it to check basic information about a delivered event.
// If you want more details, use `sc.V2Events.Get(thinEvent.ID)`
// to fetch the full event object.
type ThinEvent struct {
	// Unique identifier for the event
	ID string `json:"id"`
	// The string "event"
	Object string `json:"object"`
	// The type of the event
	Type string `json:"type"`
	// Livemode indicates if the event is from a production(true) or test(false) account
	Livemode bool `json:"livemode"`
	// Time at which the event was created
	Created time.Time `json:"created"`
	// [Optional] Object containing the reference to API resource relevant to the event
	RelatedObject *RelatedObject `json:"related_object"`
	// [Optional] Authentication context needed to fetch the event or related object
	Context *string `json:"context"`
	// [Optional] Reason for the event
	Reason *V2EventReason `json:"reason"`
}

type RelatedObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
