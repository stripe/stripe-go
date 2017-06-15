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

	if signedPayload.signature == nil {
		signedPayload.signature = computeSignature(signedPayload.timestamp, signedPayload.payload, signedPayload.secret)
	}
	signedPayload.header = generateHeader(*signedPayload)
	return signedPayload
}

func (p *SignedPayload) hexSignature() string {
	return hex.EncodeToString(p.signature)
}

func generateHeader(p SignedPayload) string {
	return fmt.Sprintf("t=%d,%s=%s", p.timestamp.Unix(), p.scheme, hex.EncodeToString(p.signature))
}

func TestTokenNew(t *testing.T) {
	p := newSignedPayload()

	evt, err := ConstructEvent(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Error validating signature: %v", err)
	} else if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test payload, got %v", evt)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.payload = append(p.payload, byte('['))
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret)
	if err == nil {
		t.Errorf("Invalid JSON did not cause a parse error")
	}

	p = newSignedPayload()
	evt, err = ConstructEvent(p.payload, "", p.secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}

	evt, err = ConstructEvent(p.payload, "v1,t=1", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	evt, err = ConstructEvent(p.payload, "t=", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	evt, err = ConstructEvent(p.payload, p.header+",v1=bad_signature", p.secret)
	if err != nil {
		t.Errorf("Received unexpected %v error with an unreadable signature in the header (should be ignored)", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.scheme = "v0"
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.signature = []byte("deadbeef")
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}

	p = newSignedPayload()
	p2 := newSignedPayload(func(p *SignedPayload) {
		p.secret = testSecret + "_rolled_key"
	})
	headerWithRolledKey := p.header + ",v1=" + p2.hexSignature()
	if p.hexSignature() == p2.hexSignature() {
		t.Errorf("Got the same signature with two different secret keys")
	}

	evt, err = ConstructEvent(p.payload, headerWithRolledKey, p.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with old key after rolling key, but got %v", err)
	}
	evt, err = ConstructEvent(p.payload, headerWithRolledKey, p2.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with new key after rolling key, but got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Now().Add(-15 * time.Second)
	})
	evt, err = ConstructEventWithTolerance(p.payload, p.header, p.secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}

	evt, err = ConstructEventWithTolerance(p.payload, p.header, p.secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Unix(12345, 0)
	})
	evt, err = ConstructEventIgnoringTolerance(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
}
