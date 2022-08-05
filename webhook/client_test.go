package webhook

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stripe/stripe-go/v73"
)

var testPayload = []byte(fmt.Sprintf(`{
  "id": "evt_test_webhook",
  "object": "event",
  "api_version": "%s"
}`, stripe.APIVersion))
var testPayloadWithAPIVersionMismatch = []byte(`{
	"id": "evt_test_webhook",
	"object": "event",
	"api_version": "2020-01-01"
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
		signedPayload.signature = ComputeSignature(signedPayload.timestamp, signedPayload.payload, signedPayload.secret)
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
	err = ValidatePayload(p.payload, "", p.secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}
	evt, err = ConstructEvent(p.payload, "", p.secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}

	evt, err = ConstructEvent(p.payload, "v1,t=1", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	err = ValidatePayload(p.payload, "t=", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}
	evt, err = ConstructEvent(p.payload, "t=", p.secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	err = ValidatePayload(p.payload, p.header+",v1=bad_signature", p.secret)
	if err != nil {
		t.Errorf("Received unexpected %v error with an unreadable signature in the header (should be ignored)", err)
	}
	evt, err = ConstructEvent(p.payload, p.header+",v1=bad_signature", p.secret)
	if err != nil {
		t.Errorf("Received unexpected %v error with an unreadable signature in the header (should be ignored)", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.scheme = "v0"
	})
	err = ValidatePayload(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}
	evt, err = ConstructEvent(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.signature = []byte("deadbeef")
	})
	err = ValidatePayload(p.payload, p.header, p.secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}
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

	err = ValidatePayload(p.payload, headerWithRolledKey, p.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with old key after rolling key, but got %v", err)
	}
	evt, err = ConstructEvent(p.payload, headerWithRolledKey, p.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with old key after rolling key, but got %v", err)
	}
	err = ValidatePayload(p.payload, headerWithRolledKey, p2.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with new key after rolling key, but got %v", err)
	}
	evt, err = ConstructEvent(p.payload, headerWithRolledKey, p2.secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with new key after rolling key, but got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Now().Add(-15 * time.Second)
	})
	err = ValidatePayloadWithTolerance(p.payload, p.header, p.secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}
	evt, err = ConstructEventWithTolerance(p.payload, p.header, p.secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}

	err = ValidatePayloadWithTolerance(p.payload, p.header, p.secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}
	evt, err = ConstructEventWithTolerance(p.payload, p.header, p.secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Unix(12345, 0)
	})
	err = ValidatePayloadIgnoringTolerance(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
	evt, err = ConstructEventIgnoringTolerance(p.payload, p.header, p.secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
}

func TestConstructEvent_ErrorOnAPIVersionMismatch(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.payload = testPayloadWithAPIVersionMismatch
	})

	_, err := ConstructEvent(p.payload, p.header, p.secret)

	if err == nil {
		t.Errorf("Expected error due to API version mismatch.")
	}

	if !strings.Contains(err.Error(), "Received event with API version") {
		t.Errorf("Expected API version mismatch error but received %v", err)
	}
}

func TestConstructEventWithOptions_IgnoreAPIVersionMismatch(t *testing.T) {

	p := newSignedPayload(func(p *SignedPayload) {
		p.payload = testPayloadWithAPIVersionMismatch
	})

	evt, err := ConstructEventWithOptions(p.payload, p.header, p.secret, ConstructEventOptions{IgnoreAPIVersionMismatch: true})

	if err != nil {
		t.Errorf("Expected no error due ignoreAPIVersionMismatch.")
	}

	if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test payload, got %v", evt)
	}
}

func TestConstructEventWithOptions_UsesDefaultToleranceWhenNoneProvided(t *testing.T) {

	p := newSignedPayload(func(p *SignedPayload) {
		// Get close to the default tolerance, but give wiggle room to avoid
		// a flaky test.
		p.timestamp = time.Now().Add(-DefaultTolerance).Add(1 * time.Second)
	})

	_, err := ConstructEventWithOptions(p.payload, p.header, p.secret, ConstructEventOptions{})

	if err != nil {
		t.Errorf("Expected no error due tolerance, but got %v.", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.timestamp = time.Now().Add(-DefaultTolerance).Add(-1 * time.Millisecond)
	})

	_, err = ConstructEventWithOptions(p.payload, p.header, p.secret, ConstructEventOptions{})

	if err != ErrTooOld {
		t.Errorf("Expected error due to being too old, but got %v.", err)
	}
}
