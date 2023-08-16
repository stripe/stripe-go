package webhook

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stripe/stripe-go/v75"
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

func newSignedPayload(options ...func(*SignedPayload)) *SignedPayload {
	signedPayload := &SignedPayload{}
	signedPayload.Timestamp = time.Now()
	signedPayload.Payload = testPayload
	signedPayload.Secret = testSecret
	signedPayload.Scheme = "v1"

	for _, opt := range options {
		opt(signedPayload)
	}

	if signedPayload.Signature == nil {
		signedPayload.Signature = ComputeSignature(signedPayload.Timestamp, signedPayload.Payload, signedPayload.Secret)
	}
	signedPayload.Header = generateHeader(*signedPayload)
	return signedPayload
}

func (p *SignedPayload) hexSignature() string {
	return hex.EncodeToString(p.Signature)
}

func TestTokenNew(t *testing.T) {
	p := newSignedPayload()

	evt, err := ConstructEvent(p.Payload, p.Header, p.Secret)
	if err != nil {
		t.Errorf("Error validating signature: %v", err)
	} else if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test Payload, got %v", evt)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Payload = append(p.Payload, byte('['))
	})
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret)
	if err == nil {
		t.Errorf("Invalid JSON did not cause a parse error")
	}

	p = newSignedPayload()
	err = ValidatePayload(p.Payload, "", p.Secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, "", p.Secret)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %v", err)
	}

	evt, err = ConstructEvent(p.Payload, "v1,t=1", p.Secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	err = ValidatePayload(p.Payload, "t=", p.Secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, "t=", p.Secret)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %v", err)
	}

	err = ValidatePayload(p.Payload, p.Header+",v1=bad_signature", p.Secret)
	if err != nil {
		t.Errorf("Received unexpected %v error with an unreadable signature in the header (should be ignored)", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header+",v1=bad_signature", p.Secret)
	if err != nil {
		t.Errorf("Received unexpected %v error with an unreadable signature in the header (should be ignored)", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Scheme = "v0"
	})
	err = ValidatePayload(p.Payload, p.Header, p.Secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Signature = []byte("deadbeef")
	})
	err = ValidatePayload(p.Payload, p.Header, p.Secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}

	p = newSignedPayload()
	p2 := newSignedPayload(func(p *SignedPayload) {
		p.Secret = testSecret + "_rolled_key"
	})
	headerWithRolledKey := p.Header + ",v1=" + p2.hexSignature()
	if p.hexSignature() == p2.hexSignature() {
		t.Errorf("Got the same signature with two different secret keys")
	}

	err = ValidatePayload(p.Payload, headerWithRolledKey, p.Secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with old key after rolling key, but got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, headerWithRolledKey, p.Secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with old key after rolling key, but got %v", err)
	}
	err = ValidatePayload(p.Payload, headerWithRolledKey, p2.Secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with new key after rolling key, but got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, headerWithRolledKey, p2.Secret)
	if err != nil {
		t.Errorf("Expected to be able to decode webhook with new key after rolling key, but got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Timestamp = time.Now().Add(-15 * time.Second)
	})
	err = ValidatePayloadWithTolerance(p.Payload, p.Header, p.Secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}
	evt, err = ConstructEventWithTolerance(p.Payload, p.Header, p.Secret, 10*time.Second)
	if err != ErrTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}

	err = ValidatePayloadWithTolerance(p.Payload, p.Header, p.Secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}
	evt, err = ConstructEventWithTolerance(p.Payload, p.Header, p.Secret, 20*time.Second)
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Timestamp = time.Unix(12345, 0)
	})
	err = ValidatePayloadIgnoringTolerance(p.Payload, p.Header, p.Secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
	evt, err = ConstructEventIgnoringTolerance(p.Payload, p.Header, p.Secret)
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
}

func TestConstructEvent_ErrorOnAPIVersionMismatch(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithAPIVersionMismatch
	})

	_, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	if err == nil {
		t.Errorf("Expected error due to API version mismatch.")
	}

	if !strings.Contains(err.Error(), "Received event with API version") {
		t.Errorf("Expected API version mismatch error but received %v", err)
	}
}

func TestConstructEventWithOptions_IgnoreAPIVersionMismatch(t *testing.T) {

	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithAPIVersionMismatch
	})

	evt, err := ConstructEventWithOptions(p.Payload, p.Header, p.Secret, ConstructEventOptions{IgnoreAPIVersionMismatch: true})

	if err != nil {
		t.Errorf("Expected no error due ignoreAPIVersionMismatch.")
	}

	if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test Payload, got %v", evt)
	}
}

func TestConstructEventWithOptions_UsesDefaultToleranceWhenNoneProvided(t *testing.T) {

	p := newSignedPayload(func(p *SignedPayload) {
		// Get close to the default tolerance, but give wiggle room to avoid
		// a flaky test.
		p.Timestamp = time.Now().Add(-DefaultTolerance).Add(1 * time.Second)
	})

	_, err := ConstructEventWithOptions(p.Payload, p.Header, p.Secret, ConstructEventOptions{})

	if err != nil {
		t.Errorf("Expected no error due tolerance, but got %v.", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Timestamp = time.Now().Add(-DefaultTolerance).Add(-1 * time.Millisecond)
	})

	_, err = ConstructEventWithOptions(p.Payload, p.Header, p.Secret, ConstructEventOptions{})

	if err != ErrTooOld {
		t.Errorf("Expected error due to being too old, but got %v.", err)
	}
}
