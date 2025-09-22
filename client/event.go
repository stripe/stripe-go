package client

import (
	"encoding/json"
	"time"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
)

// ParseThinEvent parses a Stripe event from the payload and verifies its signature.
// It returns a union of all possible event notification types that implement EventNotificationContainer.
func (a *API) ParseThinEvent(payload []byte, header string, secret string) (stripe.EventNotificationContainer, error) {
	return a.ParseThinEventWithTolerance(payload, header, secret, webhook.DefaultTolerance)
}

// ParseThinEventWithTolerance parses a Stripe event from the payload and verifies its signature with a custom tolerance.
// Generally, you should use ParseThinEvent, which uses the default tolerance.
func (a *API) ParseThinEventWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) (stripe.EventNotificationContainer, error) {
	if err := webhook.ValidatePayloadWithTolerance(payload, header, secret, tolerance); err != nil {
		return nil, err
	}
	var event stripe.UnknownEventNotification
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}

	switch event.Type {
	case "v1.billing.meter.error_report_triggered":
		result := &stripe.V1BillingMeterErrorReportTriggeredEventNotification{}
		result.EventNotification = event.EventNotification
		if err := json.Unmarshal(payload, result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return &event, nil
	}
}
