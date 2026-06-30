package stripe

// most of this file is codegen, but not quite all of it.
// V2Events: The beginning of the section generated from our OpenAPI spec

import (
	"encoding/json"
	"fmt"
	"strings"
)

// V2CoreEvent is the interface implemented by V2 Events. To get the underlying Event,
// use a type switch or type assertion to one of the concrete event types.
type V2CoreEvent interface {
	getBaseEvent() *V2BaseEvent
}

// V2CoreRawEvent is the raw event type for V2 events. It is used to unmarshal the
// event data into a generic structure, and can also be used a default event
// type when the event type is not known.
type V2CoreRawEvent struct {
	V2BaseEvent
	Data          *json.RawMessage          `json:"data"`
	RelatedObject *V2CoreEventRelatedObject `json:"related_object"`
}

// Used for everything internal to the EventNotifications
type eventNotificationParams struct {
	Params `form:"*"`
}
// V2Events: The end of the section generated from our OpenAPI spec

// EventNotificationFromJSON is a helper for constructing an Event Notification. Doesn't perform signature validation,
// so you should use [Client.ParseEventNotification] instead for initial handling.
// This is useful in unit tests and working with EventNotifications that you've already validated the authenticity of.
func EventNotificationFromJSON(payload []byte, client Client) (EventNotificationContainer, error) {
	var result = &struct {
		Type   string `json:"type"`
		Object string `json:"object"`
	}{}
	if err := json.Unmarshal(payload, result); err != nil {
		return nil, err
	}

	// TODO(DEVSDK-2968): Remove once Custom Events are sent with the correct type.
	if strings.HasPrefix(result.Type, "v2.extend.objects.object_record") {
		// Parse RelatedObject to extract the type
		var temp struct {
			RelatedObject *struct {
				Type string `json:"type"`
			} `json:"related_object"`
		}
		if err := json.Unmarshal(payload, &temp); err == nil {
			if temp.RelatedObject != nil && temp.RelatedObject.Type != "" {
				result.Type = strings.Replace(result.Type, "v2.extend.objects.object_record", temp.RelatedObject.Type, 1)
			}
		}
	}

	if result.Object == "event" {
		return nil, fmt.Errorf("did you use EventNotificationFromJSON to parse a webhook payload? If so, use ConstructEvent instead")
	}

	// V2EventNotificationTypes: The beginning of the section generated from our OpenAPI spec
	evt := UnknownEventNotification{}
	if err := json.Unmarshal(payload, &evt); err != nil {
		return nil, err
	}
	evt.client = client
	return &evt, nil
	// V2EventNotificationTypes: The end of the section generated from our OpenAPI spec
}
