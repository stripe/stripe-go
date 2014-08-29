package stripe

import (
	"net/url"
	"strconv"
)

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	Id       string     `json:"id"`
	Live     bool       `json:"livemode"`
	Created  int64      `json:"created"`
	Data     *EventData `json:"data"`
	Webhooks uint64     `json:"pending_webhooks"`
	Type     string     `json:"type"`
	Req      string     `json:"request"`
}

// EventData is the unmarshalled object as a map.
type EventData struct {
	Obj  map[string]interface{} `json:"object"`
	Prev map[string]interface{} `json:"previous_attributes"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams
	Created int64
	// Type is one of the values documented at https://stripe.com/docs/api#event_types.
	Type string
}

// EventList is a list object for events.
type EventList struct {
	ListResponse
	Values []*Event `json:"data"`
}

// EventClient is the client used to invoke /events APIs.
type EventClient struct {
	api   Api
	token string
}

// Get returns the details of an event
// For more details see https://stripe.com/docs/api#retrieve_event.
func (c *EventClient) Get(id string) (*Event, error) {
	event := &Event{}
	err := c.api.Call("GET", "/events/"+id, c.token, nil, event)

	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func (c *EventClient) List(params *EventListParams) (*EventList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Type) > 0 {
			body.Add("type", params.Type)
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Start) > 0 {
			body.Add("starting_after", params.Start)
		}

		if len(params.End) > 0 {
			body.Add("ending_before", params.End)
		}

		if params.Limit > 0 {
			if params.Limit > 100 {
				params.Limit = 100
			}

			body.Add("limit", strconv.FormatUint(params.Limit, 10))
		}
	}

	list := &EventList{}
	err := c.api.Call("GET", "/events", c.token, body, list)

	return list, err
}

// GetObjValue returns the value from the e.Data.Obj bag based on the keys hierarchy.
func (e *Event) GetObjValue(keys ...string) string {
	return getValue(e.Data.Obj, keys)
}

// GetPrevValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPrevValue(keys ...string) string {
	return getValue(e.Data.Prev, keys)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		node = node.(map[string]interface{})[keys[i]]
	}

	return node.(string)
}
