package webhook

import (
	"encoding/hex"
	"fmt"
	"testing"
  "time"

	stripe "github.com/stripe/stripe-go"
)

var testPayload = []byte(`{
  "id": "evt_test_webhook",
  "object": "event"
}`)
var testSecret = "whsec_test_secret"


type SignedPayload struct {
	timestamp time.Time
	payload []byte
	secret string
	scheme string
	signature []byte
	header string
}

func newSignedPayload(options ...func (*SignedPayload)) (*SignedPayload) {
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
		signature = computeSignature(fmt.Sprintf("%d", p.timestamp.Unix()), p.payload, p.secret)
	}

	return fmt.Sprintf("t=%d,%s=%s", p.timestamp.Unix(), p.scheme, hex.EncodeToString(signature))
}

func testEvent(p SignedPayload) (stripe.Event, error) {
	return ConstructEvent(p.payload, p.header, p.secret, ToleranceDefault)
}

func TestTokenNew(t *testing.T) {
	p := newSignedPayload()

	evt, err := ConstructEvent(p.payload, p.header, p.secret, ToleranceDefault)
	if err != nil {
		t.Errorf("Webhook did not match generated signature")
	} else if evt.ID != "evt_test_webhook" {
		t.Errorf("Expected a parsed event matching the test payload, got %t", evt)
	}

	p = newSignedPayload(func (p *SignedPayload) {
		p.payload = append(p.payload, byte('['))
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret, ToleranceDefault)
	if err == nil {
		t.Errorf("Invalid JSON did not cause a parse error")
	}

	p = newSignedPayload()
	evt, err = ConstructEvent(p.payload, "", p.secret, ToleranceDefault)
	if err != ErrNotSigned {
		t.Errorf("Expected ErrNotSigned from missing signature, got %t", err)
	}

	evt, err = ConstructEvent(p.payload, "v1,t=1", p.secret, ToleranceDefault)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %t", err)
	}

	evt, err = ConstructEvent(p.payload, "t=", p.secret, ToleranceDefault)
	if err != ErrInvalidHeader {
		t.Errorf("Expected ErrInvalidHeader from bad header format, got %t", err)
	}

	p = newSignedPayload(func (p *SignedPayload) {
		p.scheme = "v0"
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret, ToleranceDefault)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %t", err)
	}

	p = newSignedPayload(func (p *SignedPayload) {
		p.signature = []byte("deadbeef")
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret, ToleranceDefault)
	if err != ErrNoValidSignature {
		t.Errorf("Expected error from fake signature, got %t", err)
	}


	p = newSignedPayload(func (p *SignedPayload) {
		p.timestamp = time.Now().Add(-15 * time.Second)
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret, 10)
	if err != ErrTooOld {
		t.Errorf("Received %t error when validating timestamp outside of allowed timing window", err)
	}

	evt, err = ConstructEvent(p.payload, p.header, p.secret, 15)
	if err != nil {
		t.Errorf("Received %t error when validating timestamp inside allowed timing window", err)
	}

	evt, err = ConstructEvent(p.payload, p.header + ",v1=bad_signature", p.secret, 15)
	if err != nil {
		t.Errorf("Received %t error when at least one valid signature in the header", err)
	}

	p = newSignedPayload(func (p *SignedPayload) {
		p.timestamp = time.Unix(12345, 0)
	})
	evt, err = ConstructEvent(p.payload, p.header, p.secret, ToleranceIgnoreTimestamp)
	if err != nil {
		t.Errorf("Received %t error when timestamp outside window but no tolerance specified", err)
	}
}
