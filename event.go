package stripe

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	Account  string        `json:"account"`
	Created  int64         `json:"created"`
	Data     *EventData    `json:"data"`
	ID       string        `json:"id"`
	Live     bool          `json:"livemode"`
	Request  *EventRequest `json:"request"`
	Type     string        `json:"type"`
	Webhooks uint64        `json:"pending_webhooks"`
}

// EventRequest contains information on a request that created an event.
type EventRequest struct {
	// ID is the request ID of the request that created an event, if the event
	// was created by a request.
	ID string `json:"id"`

	// IdempotencyKey is the idempotency key of the request that created an
	// event, if the event was created by a request and if an idempotency key
	// was specified for that request.
	IdempotencyKey string `json:"idempotency_key"`
}

// EventData is the unmarshalled object as a map.
type EventData struct {
	Obj  map[string]interface{}
	Prev map[string]interface{} `json:"previous_attributes"`
	Raw  json.RawMessage        `json:"object"`
}

// EventList is a list of events as retrieved from a list endpoint.
type EventList struct {
	ListMeta
	Values []*Event `json:"data"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Type         string            `form:"type"`
	Types        []string          `form:"types"`
}

// GetObjValue returns the value from the e.Data.Obj bag based on the keys hierarchy.
func (e *Event) GetObjValue(keys ...string) string {
	return getValue(e.Data.Obj, keys)
}

// GetPrevValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPrevValue(keys ...string) string {
	return getValue(e.Data.Prev, keys)
}

// UnmarshalJSON handles deserialization of the EventData.
// This custom unmarshaling exists so that we can keep both the map and raw data.
func (e *EventData) UnmarshalJSON(data []byte) error {
	type eventdata EventData
	var ee eventdata
	err := json.Unmarshal(data, &ee)
	if err != nil {
		return err
	}

	*e = EventData(ee)
	return json.Unmarshal(e.Raw, &e.Obj)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		key := keys[i]

		sliceNode, ok := node.([]interface{})
		if ok {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				panic(fmt.Sprintf(
					"Cannot access nested slice element with non-integer key: %s",
					key))
			}
			node = sliceNode[intKey]
			continue
		}

		mapNode, ok := node.(map[string]interface{})
		if ok {
			node = mapNode[key]
			continue
		}

		panic(fmt.Sprintf(
			"Cannot descend into non-map non-slice object with key: %s", key))
	}

	if node == nil {
		return ""
	}

	return fmt.Sprintf("%v", node)
}
