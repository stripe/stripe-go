package stripe

import (
	"encoding/json"
	"strings"
)

// Event is the resource representing a Stripe event.
// For more details see https://stripe.com/docs/api#events.
type Event struct {
	ID       string     `json:"id"`
	Live     bool       `json:"livemode"`
	Created  int64      `json:"created"`
	Data     *EventData `json:"data"`
	Webhooks uint64     `json:"pending_webhooks"`
	Type     string     `json:"type"`
	Req      string     `json:"request"`
}

// EventData is the unmarshalled object as a map.
type EventData struct {
	Raw     json.RawMessage        `json:"object"`
	PrevRaw json.RawMessage        `json:"previous_attributes"`
	Prev    map[string]interface{} `json:"-"`
	Obj     map[string]interface{} `json:"-"`
}

// EventListParams is the set of parameters that can be used when listing events.
// For more details see https://stripe.com/docs/api#list_events.
type EventListParams struct {
	ListParams
	Created int64
	// Type is one of the values documented at https://stripe.com/docs/api#event_types.
	Type string
}

// GetObjValue returns the value from the e.Data.Obj bag based on the keys hierarchy.
func (e *Event) GetObjValue(keys ...string) string {
	return getValue(e.Data.Obj, keys)
}

// GetPrevValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPrevValue(keys ...string) string {
	return getValue(e.Data.Prev, keys)
}

// EventData returns an interface that represents the correct struct for this event type. If no struct is matched, a map[string]interface{} is returned.
func (e *Event) EventData() interface{} {
	eventData := parseEventType(e.Type)
	if err := json.Unmarshal(e.Data.Raw, &eventData); err == nil {
		return eventData
	} else {
		return nil
	}
}

// PrevEventData returns an interface that represents the correct struct for this event type. If no struct is matched, a map[string]interface{} is returned.
func (e *Event) PrevEventData() interface{} {
	eventData := parseEventType(e.Type)
	if err := json.Unmarshal(e.Data.PrevRaw, &eventData); err == nil {
		return eventData
	} else {
		return nil
	}
}

// UnmarshalJSON handles deserialization of the EventData.
// This custom unmarshaling exists so that we can keep both the map and raw data.
func (e *EventData) UnmarshalJSON(data []byte) error {
	type eventdata EventData
	var ee eventdata
	if err := json.Unmarshal(data, &ee); err != nil {
		return err
	}

	*e = EventData(ee)

	if e.PrevRaw != nil {
		if err := json.Unmarshal(e.PrevRaw, &e.Prev); err != nil {
			return err
		}
	}
	return json.Unmarshal(e.Raw, &e.Obj)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		node = node.(map[string]interface{})[keys[i]]
	}

	if node == nil {
		return ""
	}

	return node.(string)
}

func parseEventType(eventTypeName string) interface{} {
	switch eventTypeName {
	case "bitcoin.receiver.created":
		return &BitcoinReceiver{}
	case "bitcoin.receiver.transaction.created":
		return &BitcoinReceiver{}
	case "bitcoin.receiver.filled":
		return &BitcoinReceiver{}
	case "ping":
		return make(map[string]interface{})
	default:
		if parts := strings.Split(eventTypeName, "."); len(parts) > 1 {
			resourceType := parts[len(parts)-2]
			switch resourceType {
			case "account":
				return &Account{}
			case "application_fee":
				return &Fee{}
			case "balance":
				return &Balance{}
			case "charge":
				return &Charge{}
			case "dispute":
				return &Dispute{}
			case "customer":
				return &Customer{}
			case "card":
				return &Card{}
			case "subscription":
				return &Sub{}
			case "discount":
				return &Discount{}
			case "invoice":
				return &Invoice{}
			case "invoiceitem":
				return &InvoiceItem{}
			case "plan":
				return &Plan{}
			case "coupon":
				return &Coupon{}
			case "recipient":
				return &Recipient{}
			case "transfer":
				return &Transfer{}
			}
		}
	}
	return make(map[string]interface{})
}
