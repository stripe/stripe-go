//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// List events, going back up to 30 days. Each event data is rendered according to Stripe API version at its creation time, specified in [event object](https://stripe.com/docs/api/events/object) api_version attribute (not according to your current Stripe API version or Stripe-Version header).
type EventListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Filter events by whether all webhooks were successfully delivered. If false, events which are still pending or have failed all delivery attempts to a webhook endpoint will be returned.
	DeliverySuccess *bool `form:"delivery_success"`
	// A string containing a specific event name, or group of events using * as a wildcard. The list will be filtered to include only events with a matching event property.
	Type *string `form:"type"`
	// An array of up to 20 strings containing specific event names. The list will be filtered to include only events with a matching event property. You may pass either `type` or `types`, but not both.
	Types []*string `form:"types"`
}

// Retrieves the details of an event. Supply the unique identifier of the event, which you might have received in a webhook.
type EventParams struct {
	Params `form:"*"`
}
type EventData struct {
	// Object is a raw mapping of the API resource contained in the event.
	// Although marked with json:"-", it's still populated independently by
	// a custom UnmarshalJSON implementation.
	// Object containing the API resource relevant to the event. For example, an `invoice.created` event will have a full [invoice object](https://stripe.com/docs/api#invoice_object) as the value of the object key.
	Object map[string]interface{} `json:"-"`
	// Object containing the names of the attributes that have changed, and their previous values (sent along only with *.updated events).
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	Raw                json.RawMessage        `json:"object"`
}

// Information on the API request that instigated the event.
type EventRequest struct {
	// ID is the request ID of the request that created an event, if the event
	// was created by a request.
	// ID of the API request that caused the event. If null, the event was automatic (e.g., Stripe's automatic subscription handling). Request logs are available in the [dashboard](https://dashboard.stripe.com/logs), but currently not in the API.
	ID string `json:"id"`

	// IdempotencyKey is the idempotency key of the request that created an
	// event, if the event was created by a request and if an idempotency key
	// was specified for that request.
	// The idempotency key transmitted during the request, if any. *Note: This property is populated only for events on or after May 23, 2017*.
	IdempotencyKey string `json:"idempotency_key"`
}

// Events are our way of letting you know when something interesting happens in
// your account. When an interesting event occurs, we create a new `Event`
// object. For example, when a charge succeeds, we create a `charge.succeeded`
// event; and when an invoice payment attempt fails, we create an
// `invoice.payment_failed` event. Note that many API requests may cause multiple
// events to be created. For example, if you create a new subscription for a
// customer, you will receive both a `customer.subscription.created` event and a
// `charge.succeeded` event.
//
// Events occur when the state of another API resource changes. The state of that
// resource at the time of the change is embedded in the event's data field. For
// example, a `charge.succeeded` event will contain a charge, and an
// `invoice.payment_failed` event will contain an invoice.
//
// As with other API resources, you can use endpoints to retrieve an
// [individual event](https://stripe.com/docs/api#retrieve_event) or a [list of events](https://stripe.com/docs/api#list_events)
// from the API. We also have a separate
// [webhooks](http://en.wikipedia.org/wiki/Webhook) system for sending the
// `Event` objects directly to an endpoint on your server. Webhooks are managed
// in your
// [account settings](https://dashboard.stripe.com/account/webhooks),
// and our [Using Webhooks](https://stripe.com/docs/webhooks) guide will help you get set up.
//
// When using [Connect](https://stripe.com/docs/connect), you can also receive notifications of
// events that occur in connected accounts. For these events, there will be an
// additional `account` attribute in the received `Event` object.
//
// **NOTE:** Right now, access to events through the [Retrieve Event API](https://stripe.com/docs/api#retrieve_event) is
// guaranteed only for 30 days.
type Event struct {
	APIResource
	// The connected account that originated the event.
	Account string `json:"account"`
	// The Stripe API version used to render `data`. *Note: This property is populated only for events on or after October 31, 2014*.
	APIVersion string `json:"api_version"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64      `json:"created"`
	Data    *EventData `json:"data"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Number of webhooks that have yet to be successfully delivered (i.e., to return a 20x response) to the URLs you've specified.
	PendingWebhooks int64 `json:"pending_webhooks"`
	// Information on the API request that instigated the event.
	Request *EventRequest `json:"request"`
	// Description of the event (e.g., `invoice.created` or `charge.refunded`).
	Type string `json:"type"`
}

// EventList is a list of Events as retrieved from a list endpoint.
type EventList struct {
	APIResource
	ListMeta
	Data []*Event `json:"data"`
}

// GetObjectValue returns the value from the e.Data.Object bag based on the keys hierarchy.
func (e *Event) GetObjectValue(keys ...string) string {
	return getValue(e.Data.Object, keys)
}

// GetPreviousValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPreviousValue(keys ...string) string {
	return getValue(e.Data.PreviousAttributes, keys)
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
	return json.Unmarshal(e.Raw, &e.Object)
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
