package stripe

import (
	"encoding/json"
	"fmt"
)

// extractFromCloudProviderEnvelope detects and unwraps an AWS EventBridge
// (https://docs.stripe.com/event-destinations/eventbridge) or Azure Event Grid
// (https://docs.stripe.com/event-destinations/eventgrid) payload, returning the
// raw inner bytes of the Stripe event or event notification.
func extractFromCloudProviderEnvelope(payload []byte) (json.RawMessage, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(payload, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse cloud event payload: %w", err)
	}

	var innerBytes json.RawMessage

	// Could add as many checks as we want here, but we'll start simple
	if detail, ok := raw["detail"]; ok {
		// AWS
		// https://docs.stripe.com/event-destinations/eventbridge#event-structure
		innerBytes = detail
	} else if _, ok := raw["specversion"]; ok {
		// Azure
		// https://docs.stripe.com/event-destinations/eventgrid#event-structure
		data, dataOk := raw["data"]
		if !dataOk {
			return nil, fmt.Errorf(
				"unrecognized cloud event format; the payload must be an " +
					"AWS EventBridge or Azure Event Grid event envelope")
		}
		innerBytes = data
	} else if idRaw, ok := raw["id"]; ok {
		var id string
		if json.Unmarshal(idRaw, &id) == nil && len(id) > 4 && id[:4] == "evt_" {
			return nil, fmt.Errorf(
				"it looks like you passed a Stripe Event directly; " +
					"use ConstructEvent instead to parse a webhook payload " +
					"with signature verification")
		}
		return nil, fmt.Errorf(
			"unrecognized cloud event format; the payload must be an " +
				"AWS EventBridge or Azure Event Grid event envelope")
	} else {
		return nil, fmt.Errorf(
			"unrecognized cloud event format; the payload must be an " +
				"AWS EventBridge or Azure Event Grid event envelope")
	}

	return innerBytes, nil
}
