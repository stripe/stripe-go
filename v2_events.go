package stripe

import "strings"

// most of this file is codegen, but not quite all of it.
// V2Events: The beginning of the section generated from our OpenAPI spec
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
	// V2EventNotificationTypes: The end of the section generated from our OpenAPI spec
}
