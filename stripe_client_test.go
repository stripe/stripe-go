package stripe_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v82"
)

func TestNewClient(t *testing.T) {
	client := stripe.NewClient("foo")
	assert.NotNil(t, client)
}

func TestNewClientWithNilBackend(t *testing.T) {
	client := stripe.NewClient("foo", nil)
	assert.NotNil(t, client)
}

func TestParseThinEvent(t *testing.T) {
	thinEvent := &stripe.EventNotification{
		ID:       "evt_123",
		Object:   "event",
		Type:     "charge.succeeded",
		Livemode: true,
		Created:  time.Now(),
		Context:  nil,
		RelatedObject: &stripe.RelatedObject{
			ID:   "ch_123",
			Type: "charge",
			URL:  "/v1/charges/ch_123",
		},
	}
	thinEventJSON, err := json.Marshal(thinEvent)
	assert.NoError(t, err)
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   thinEventJSON,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)
	sc := stripe.NewClient("sk_test_secret")
	thinEvent, err = sc.ParseThinEvent(p.Payload, p.Header, p.Secret)
	assert.NoError(t, err)
	assert.Equal(t, "evt_123", thinEvent.ID)
	assert.Equal(t, "event", thinEvent.Object)
	assert.Equal(t, "charge.succeeded", thinEvent.Type)
	assert.Equal(t, true, thinEvent.Livemode)
	assert.Equal(t, "ch_123", thinEvent.RelatedObject.ID)
	assert.Equal(t, "charge", thinEvent.RelatedObject.Type)
	assert.Equal(t, "/v1/charges/ch_123", thinEvent.RelatedObject.URL)
	assert.NotNil(t, thinEvent)
}

func TestParseThinEventNoTolerance(t *testing.T) {
	thinEvent := &stripe.EventNotification{
		ID:       "evt_123",
		Object:   "event",
		Type:     "charge.succeeded",
		Livemode: true,
		Created:  time.Now(),
		Context:  nil,
		RelatedObject: &stripe.RelatedObject{
			ID:   "ch_123",
			Type: "charge",
			URL:  "/v1/charges/ch_123",
		},
	}
	thinEventJSON, err := json.Marshal(thinEvent)
	assert.NoError(t, err)
	p := stripe.SignedPayload{
		UnsignedPayload: stripe.UnsignedPayload{
			Payload:   thinEventJSON,
			Timestamp: time.Now(),
			Secret:    "whsec_test_secret",
			Scheme:    "v1",
		},
	}
	p.Signature = stripe.ComputeSignature(p.Timestamp, p.Payload, p.Secret)
	p.Header = generateHeader(p)
	sc := stripe.NewClient("sk_test_secret")
	_, err = sc.ParseThinEvent(p.Payload, p.Header, p.Secret, stripe.WithTolerance(0))
	assert.Equal(t, stripe.ErrWebhookTooOld, err)
}

func generateHeader(p stripe.SignedPayload) string {
	return fmt.Sprintf("t=%d,%s=%s", p.Timestamp.Unix(), p.Scheme, hex.EncodeToString(p.Signature))
}
