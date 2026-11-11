package stripe

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

var testPayload = []byte(fmt.Sprintf(`{
  "id": "evt_test_webhook",
  "object": "event",
  "api_version": "%s"
}`, APIVersion))

var testPayloadWithNewVersionInReleaseTrain = []byte(fmt.Sprintf(`{
	"id": "evt_test_webhook",
	"object": "event",
	"api_version": "%s"
  }`, "2099-10-10."+strings.Split(APIVersion, ".")[1]))

var testPayloadWithAPIVersionMismatch = []byte(`{
	"id": "evt_test_webhook",
	"object": "event",
	"api_version": "2020-01-01"
  }`)

var testPayloadWithReleaseTrainVersionMismatch = []byte(`{
	"id": "evt_test_webhook",
	"object": "event",
	"api_version": "2099-10-10.the_larch"
  }`)

var testPayloadWithISO8601Timestamp = []byte(`{
	"id": "evt_test_webhook",
	"object": "v2.core.event",
	"created": "2025-10-04T15:13:32.232083866Z"
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

func newSignedPayloadWithDebugMetadata(metadata map[string]string, options ...func(*SignedPayload)) *SignedPayload {
	payload := newSignedPayload(options...)
	for key, value := range metadata {
		payload.Header += "," + key + "=" + value
	}
	return payload
}

func TestWebhookDiagnoseSecretMismatch(t *testing.T) {
	secret := "whsec_abcdefghijklmnopqrstuvwxyz"
	masked := "whsec_****************************1q84"
	cases := []struct {
		secret        string
		masked        string
		destination   string
		suffix        string
		dashboardPath string
	}{
		// A mismatched secret links to the supplied destination.
		{secret, masked, "we_example", "wxyz", "/we_example"},
		// A missing destination uses the bare Dashboard URL.
		{secret, masked, "", "wxyz", ""},
		// Suffix matches but length differs compared to masked secret
		{"whsec_short1q84", masked, "", "1q84", ""},
		// A mostly unmasked value still reveals only the last four secret characters.
		{secret, "whsec_*bcdefghijklmnopqrstuvwxyZ", "", "wxyz", ""},
		// A secret shorter than four characters is shown in full.
		{"abc", "*", "", "abc", ""},
	}
	for _, tc := range cases {
		p := newSignedPayloadWithDebugMetadata(map[string]string{
			"secret": tc.masked, "destination": tc.destination, "payload": "00000000",
		}, func(p *SignedPayload) { p.Signature = []byte("bad_signature") })
		err := ValidatePayload(p.Payload, p.Header, tc.secret)
		assert.Error(t, err)
		assert.Equal(t, fmt.Sprintf(
			"Webhook signature verification failed. This request may not have originated from Stripe.\n\n"+
				"The signing secret you provided is %d characters long and ends in `%s`. However, the signing secret for this destination is %d characters long and looks like `%s`. Ensure you're using the signing secret for this destination, as shown on the Dashboard: %s\n\n"+
				"Learn more: https://docs.stripe.com/webhooks/signature",
			len(tc.secret), tc.suffix, len(tc.masked), tc.masked, "https://dashboard.stripe.com/webhooks"+tc.dashboardPath,
		), err.Error())
		assert.True(t, errors.Is(err, ErrWebhookNoValidSignature))
	}
}

func TestWebhookDiagnosePayloadMismatch(t *testing.T) {
	// User supplies the original body with an errant trailing space.
	userPayload := append(append([]byte{}, testPayload...), ' ')

	originalHash := sha256.Sum256(testPayload)
	userHash := sha256.Sum256(userPayload)
	expectedHash := hex.EncodeToString(originalHash[len(originalHash)-4:])
	actualHash := hex.EncodeToString(userHash[len(userHash)-4:])
	p := newSignedPayloadWithDebugMetadata(map[string]string{
		"secret": "whsec_*******cret", "payload": strings.ToUpper(expectedHash),
	})

	err := ValidatePayload(userPayload, p.Header, testSecret)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf(
		"Webhook signature verification failed.\n\n"+
			"The request body you passed to this function differs from the request body sent by Stripe to this webhook destination.\n"+
			"The SHA-256 hash of the body should end in `%s`. However, the SHA-256 hash of the body you provided ends in `%s`.\n\n"+
			"Make sure your code does not reformat or otherwise alter the raw HTTP request body before passing it to this function.\n\n"+
			"Learn more: https://docs.stripe.com/webhooks/signature", expectedHash, actualHash,
	), err.Error())
	assert.True(t, errors.Is(err, ErrWebhookNoValidSignature))
}

func TestWebhookInvalidSecretMetadataFallsBackToGenericError(t *testing.T) {
	hash := sha256.Sum256(testPayload)
	payloadHash := hex.EncodeToString(hash[len(hash)-4:])
	for _, mask := range []string{"unmasked", strings.Repeat("*", 101)} {
		p := newSignedPayloadWithDebugMetadata(map[string]string{"secret": mask, "payload": payloadHash}, func(p *SignedPayload) { p.Signature = []byte("bad_signature") })
		assert.Equal(t, ErrWebhookNoValidSignature, ValidatePayload(p.Payload, p.Header, testSecret))
	}
}

func TestWebhookInvalidPayloadMetadataFallsBackToGenericError(t *testing.T) {
	for _, payloadHash := range []string{"1234567", "not-hex!", strings.Repeat("a", 101)} {
		p := newSignedPayloadWithDebugMetadata(map[string]string{"secret": "whsec_*******cret", "payload": payloadHash}, func(p *SignedPayload) { p.Signature = []byte("bad_signature") })
		assert.Equal(t, ErrWebhookNoValidSignature, ValidatePayload(p.Payload, p.Header, testSecret))
	}
}

func TestWebhookMalformedMetadataReturnsInvalidHeader(t *testing.T) {
	p := newSignedPayload()
	for _, key := range []string{"secret", "destination", "payload"} {
		// Metadata entries must have an equals sign, like any other header entry.
		assert.Equal(t, ErrWebhookInvalidHeader, ValidatePayload(p.Payload, p.Header+","+key, p.Secret))
	}
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
	if err != ErrWebhookNotSigned {
		t.Errorf("Expected ErrWebhookNotSigned from missing signature, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, "", p.Secret)
	if err != ErrWebhookNotSigned {
		t.Errorf("Expected ErrWebhookNotSigned from missing signature, got %v", err)
	}

	evt, err = ConstructEvent(p.Payload, "v1,t=1", p.Secret)
	if err != ErrWebhookInvalidHeader {
		t.Errorf("Expected ErrWebhookInvalidHeader from bad header format, got %v", err)
	}

	err = ValidatePayload(p.Payload, "t=", p.Secret)
	if err != ErrWebhookInvalidHeader {
		t.Errorf("Expected ErrWebhookInvalidHeader from bad header format, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, "t=", p.Secret)
	if err != ErrWebhookInvalidHeader {
		t.Errorf("Expected ErrWebhookInvalidHeader from bad header format, got %v", err)
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
	if err != ErrWebhookNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret)
	if err != ErrWebhookNoValidSignature {
		t.Errorf("Expected error from mismatched schema, got %v", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Signature = []byte("deadbeef")
	})
	err = ValidatePayload(p.Payload, p.Header, p.Secret)
	if err != ErrWebhookNoValidSignature {
		t.Errorf("Expected error from fake signature, got %v", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret)
	if err != ErrWebhookNoValidSignature {
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
	err = ValidatePayload(p.Payload, p.Header, p.Secret, WithTolerance(10*time.Second))
	if err != ErrWebhookTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret, WithTolerance(10*time.Second))
	if err != ErrWebhookTooOld {
		t.Errorf("Received %v error when validating timestamp outside of allowed timing window", err)
	}

	err = ValidatePayload(p.Payload, p.Header, p.Secret, WithTolerance(20*time.Second))
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret, WithTolerance(20*time.Second))
	if err != nil {
		t.Errorf("Received %v error when validating timestamp inside allowed timing window", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Timestamp = time.Unix(12345, 0)
	})
	err = ValidatePayload(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
	evt, err = ConstructEvent(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
	if err != nil {
		t.Errorf("Received %v error when timestamp outside window but no tolerance specified", err)
	}
}

func TestConstructEvent_SuccessOnExpectedAPIVersion(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayload
	})

	evt, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	if err != nil {
		t.Errorf("Unexpected error in ConstructEvent: %v", err)
	}

	if evt.APIVersion != APIVersion {
		t.Errorf("Expected API versions to match")
	}
}

func TestConstructEvent_SuccessOnNewAPIVersionInExpectedReleaseTrain(t *testing.T) {
	// this test only makes sense on GA versions- the exact version for preview versions doesn't
	// work this way and we can't mock private methods from this test class.
	if strings.HasSuffix(APIVersion, ".preview") {
		t.Skip("Skipping test for new API version in expected release train")
	}

	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithNewVersionInReleaseTrain
	})

	evt, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	if err != nil {
		t.Errorf("Unexpected error in ConstructEvent: %v", err)
	}

	expectedSuffix := "." + strings.Split(APIVersion, ".")[1]
	if !strings.HasSuffix(evt.APIVersion, expectedSuffix) {
		t.Errorf("Expected API release trains to match")
	}
}
func TestConstructEvent_ErrorOnLegacyAPIVersionMismatch(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithAPIVersionMismatch
	})

	_, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	if err == nil {
		t.Errorf("Expected error due to API version mismatch.")
	}

	if !strings.Contains(err.Error(), "received event with API version") {
		t.Errorf("Expected API version mismatch error but received %v", err)
	}
}

func TestConstructEvent_ErrorOnReleaseTrainMismatch(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithReleaseTrainVersionMismatch
	})

	_, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	if err == nil {
		t.Errorf("Expected error due to API version mismatch.")
	}

	if !strings.Contains(err.Error(), "received event with API version") {
		t.Errorf("Expected API version mismatch error but received %v", err)
	}
}
func TestConstructEventWithOptions_IgnoreAPIVersionMismatch(t *testing.T) {

	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithAPIVersionMismatch
	})

	evt, err := ConstructEvent(p.Payload, p.Header, p.Secret, WithIgnoreAPIVersionMismatch())

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
		p.Timestamp = time.Now().Add(-WebhookDefaultTolerance).Add(1 * time.Second)
	})

	_, err := ConstructEvent(p.Payload, p.Header, p.Secret, nil)

	if err != nil {
		t.Errorf("Expected no error due tolerance, but got %v.", err)
	}

	p = newSignedPayload(func(p *SignedPayload) {
		p.Timestamp = time.Now().Add(-WebhookDefaultTolerance).Add(-1 * time.Millisecond)
	})

	_, err = ConstructEvent(p.Payload, p.Header, p.Secret, nil)

	if err != ErrWebhookTooOld {
		t.Errorf("Expected error due to being too old, but got %v.", err)
	}
}

func TestApiVersionCompatibility(t *testing.T) {
	tests := []struct {
		sdkAPIVersion   string
		eventAPIVersion string
		expected        bool
	}{
		{"2024-02-31.acacia", "1999-03-31", false},
		{"2024-02-31.acacia", "2025-03-31.basil", false},
		{"2024-04-31.basil", "2025-03-31.basil", true},
		{"2024-01-01.preview", "2025-03-31.basil", false},
		{"2024-01-01.preview", "2025-03-31.preview", false},
		{"2024-01-01.preview", "2024-01-01.preview", true},
	}

	for _, test := range tests {
		result := isCompatibleAPIVersion(test.sdkAPIVersion, test.eventAPIVersion)
		if result != test.expected {
			t.Errorf("Expected %v for API version %s <> %s, got %v", test.expected, test.sdkAPIVersion, test.eventAPIVersion, result)
		}
	}
}

func TestGenerateTestSignedPayload(t *testing.T) {
	payload := []byte(`{"id":"evt_test"}`)
	signed := GenerateTestSignedPayload(&UnsignedPayload{
		Payload: payload,
		Secret:  "whsec_test",
	})
	if signed.Header == "" {
		t.Fatal("expected non-empty header")
	}
	if !strings.HasPrefix(signed.Header, "t=") {
		t.Errorf("expected header to start with 't=', got %q", signed.Header)
	}
	if !strings.Contains(signed.Header, ",v1=") {
		t.Errorf("expected header to contain ',v1=', got %q", signed.Header)
	}
	err := ValidatePayload(payload, signed.Header, "whsec_test", WithIgnoreTolerance())
	if err != nil {
		t.Fatalf("expected valid payload, got error: %v", err)
	}
}

func TestValidatePayload_ValidPayload(t *testing.T) {
	p := newSignedPayload()
	err := ValidatePayload(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
	if err != nil {
		t.Errorf("expected nil error for valid payload, got: %v", err)
	}
}

func TestValidatePayload_BadSignature(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Signature = []byte("deadbeef")
	})
	err := ValidatePayload(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
	if err != ErrWebhookNoValidSignature {
		t.Errorf("expected ErrWebhookNoValidSignature for bad signature, got: %v", err)
	}
}

// An empty secret must be rejected outright. HMAC accepts a zero-length key, so
// without this an attacker could forge a payload signed with an empty secret and
// have it verify successfully.
func TestValidatePayload_EmptySecret(t *testing.T) {
	tests := map[string]string{
		"empty":           "",
		"space":           " ",
		"tab":             "\t",
		"carriage return": "\r",
		"line feed":       "\n",
		"form feed":       "\f",
		"vertical tab":    "\v",
		"mixed":           " \t\r\n\f\v",
	}
	for name, secret := range tests {
		t.Run(name, func(t *testing.T) {
			p := newSignedPayload(func(p *SignedPayload) {
				p.Secret = secret
			})
			err := ValidatePayload(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
			if err != ErrWebhookEmptySecret {
				t.Errorf("expected ErrWebhookEmptySecret for blank secret, got: %v", err)
			}
		})
	}
}

func TestConstructEvent_EmptySecret(t *testing.T) {
	for name, secret := range map[string]string{
		"empty":           "",
		"space":           " ",
		"tab":             "\t",
		"carriage return": "\r",
		"line feed":       "\n",
		"form feed":       "\f",
		"vertical tab":    "\v",
		"mixed":           " \t\r\n\f\v",
	} {
		t.Run(name, func(t *testing.T) {
			p := newSignedPayload(func(p *SignedPayload) {
				p.Secret = secret
			})
			_, err := ConstructEvent(p.Payload, p.Header, p.Secret, WithIgnoreTolerance())
			if err != ErrWebhookEmptySecret {
				t.Errorf("expected ErrWebhookEmptySecret for blank secret, got: %v", err)
			}
		})
	}
}

func TestValidatePayload_NonBlankSecret(t *testing.T) {
	for name, secret := range map[string]string{
		"surrounded by ASCII whitespace": " \tsecret\r\n",
		"non-breaking space only":        "\u00a0",
	} {
		t.Run(name, func(t *testing.T) {
			p := newSignedPayload(func(p *SignedPayload) {
				p.Secret = secret
			})
			if err := ValidatePayload(p.Payload, p.Header, p.Secret, WithIgnoreTolerance()); err != nil {
				t.Errorf("expected exact non-blank secret to verify, got: %v", err)
			}
		})
	}
}

func TestConstructEvent_ErrorOnEventNotification(t *testing.T) {
	p := newSignedPayload(func(p *SignedPayload) {
		p.Payload = testPayloadWithISO8601Timestamp
	})

	_, err := ConstructEvent(p.Payload, p.Header, p.Secret)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "EventNotification")
}
