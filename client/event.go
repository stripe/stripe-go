package client

import (
	"encoding/json"
	"time"

	"github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/webhook"
)

// ParseThinEvent parses a Stripe event from the payload and verifies its signature.
// It returns a ThinEvent object and an error if the parsing or verification fails.
func (a *API) ParseThinEvent(payload []byte, header string, secret string) (*stripe.ThinEvent, error) {
	return a.ParseThinEventWithTolerance(payload, header, secret, webhook.DefaultTolerance)
}

// ParseThinEventWithTolerance parses a Stripe event from the payload and verifies its signature with a custom tolerance.
// Generally, you should use ParseThinEvent, which uses the default tolerance.
func (a *API) ParseThinEventWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) (*stripe.ThinEvent, error) {
	if err := webhook.ValidatePayloadWithTolerance(payload, header, secret, tolerance); err != nil {
		return nil, err
	}
	var event stripe.ThinEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
