package stripe

import (
	"encoding/json"
	"fmt"
)

// maybeExtractFromCloudProviderEnvelope detects and unwraps an AWS EventBridge
// (https://docs.stripe.com/event-destinations/eventbridge) or Azure Event Grid
// (https://docs.stripe.com/event-destinations/eventgrid) payload, returning the
// raw inner bytes of the Stripe event or event notification. If the payload is
// already a raw Stripe event (object is "event" or "v2.core.event"), it is
// returned as-is.
func maybeExtractFromCloudProviderEnvelope(payload []byte) (json.RawMessage, error) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(payload, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse cloud event payload: %w", err)
	}

	// Could add as many checks as we want here, but we'll start simple
	if detail, ok := raw["detail"]; ok {
		// AWS
		// https://docs.stripe.com/event-destinations/eventbridge#event-structure
		return detail, nil
	}
	if _, ok := raw["specversion"]; ok {
		// Azure
		// https://docs.stripe.com/event-destinations/eventgrid#event-structure
		if data, dataOk := raw["data"]; dataOk {
			return data, nil
		}
	}
	if objRaw, ok := raw["object"]; ok {
		var obj string
		if json.Unmarshal(objRaw, &obj) == nil && (obj == "event" || obj == "v2.core.event") {
			// Raw Stripe event passed directly: use as-is
			return json.RawMessage(payload), nil
		}
	}

	return nil, fmt.Errorf(
		"unrecognized event format. The payload must be an " +
			"AWS EventBridge/Azure Event Grid event envelope or a Stripe webhook (thin event notification or snapshot)")
}
