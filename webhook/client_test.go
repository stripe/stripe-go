package webhook

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

var testPayload = []byte(`{
  "id": "evt_test_webhook",
  "object": "event"
}`)
var testSecret = "whsec_test_secret"

type SignedPayload struct {
	timestamp time.Time
	payload   []byte
	secret    string
	scheme    string
	signature []byte
	header    string
}

func newSignedPayload(options ...func(*SignedPayload)) *SignedPayload {
	signedPayload := &SignedPayload{}
	signedPayload.timestamp = time.Now()
	signedPayload.payload = testPayload
	signedPayload.secret = testSecret
	signedPayload.scheme = "v1"

	for _, opt := range options {
		opt(signedPayload)
	}

	signedPayload.header = generateHeader(*signedPayload)
	return signedPayload
}

func generateHeader(p SignedPayload) string {
	signature := p.signature
	if signature == nil {
		signature = computeSignature(p.timestamp, p.payload, p.secret)
	}

	return fmt.Sprintf("t=%d,%s=%s", p.timestamp.Unix(), p.scheme, hex.EncodeToString(signature))
}

func TestTokenNew(t *testing.T) {
	p := newSignedPayload()

	evt, err := ValidateEvent(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Error validating signature: %v", err)
	} else if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test payload, got %v", evt)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.payload = append(p.payload, byte('['))
	})
	evt, err = ValidateEvent(p.payload, p.header, p.secret)
	if err == nil {
		t.Errorf("Invalid JSON did not cause a parse error")
	}

	p = newSignedPayload()
	evt, err = ValidateEvent(p.payload, "", p.secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}

	evt, err = ValidateEvent(p.payload, "v1,t=1", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	evt, err = ValidateEvent(p.payload, "t=", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	evt, err = ValidateEvent(p.payload, p.header+",v1=bad_signature", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Received no error with an unreadable signature in the header", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.scheme = "v0"
	})
	evt, err = ValidateEvent(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.signature = []byte("deadbeef")
	})
	evt, err = ValidateEvent(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Now().Add(-15 * time.Second)
	})
	evt, err = ValidateEventWithTolerance(p.payload, p.header, p.secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}

	evt, err = ValidateEventWithTolerance(p.payload, p.header, p.secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Unix(12345, 0)
	})
	evt, err = ValidateEventIgnoringTolerance(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
}
